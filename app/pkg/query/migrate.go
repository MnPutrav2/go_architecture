package query

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	textformater "github.com/MnPutrav2/go_architecture/app/pkg/text_formater"
)

func (q *Initdb) Migrate(table ...any) {
	fmt.Println("Running migrations... ==================")

	for _, tbl := range table {

		t := reflect.TypeOf(tbl)

		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		if t.Kind() != reflect.Struct {
			continue
		}

		var (
			createTypeSQL string
			columns       []string
			constraints   []string
		)

		for i := 0; i < t.NumField(); i++ {

			field := t.Field(i)

			ts := field.Tag.Get("structure")
			td := strings.ToLower(field.Tag.Get("db"))

			var columnDef string

			// ==========================
			// ENUM
			// ==========================
			if strings.Contains(ts, "enum(") {

				start := strings.Index(ts, "(")
				end := strings.LastIndex(ts, ")")

				values := strings.Split(ts[start+1:end], ",")

				var enumValues []string
				for _, v := range values {
					v = strings.TrimSpace(v)
					enumValues = append(enumValues, fmt.Sprintf("'%s'", v))
				}

				enumType := strings.ToLower(textformater.ToSnakeCase(t.Name())) + "_ty"
				createTypeSQL = fmt.Sprintf(`
					DO $$
					BEGIN
						IF NOT EXISTS (
							SELECT 1
							FROM pg_type
							WHERE typname = '%s'
						) THEN
							CREATE TYPE %s AS ENUM (%s);
						END IF;
					END
					$$;`,
					enumType,
					enumType,
					strings.Join(enumValues, ","),
				)
				columnDef = fmt.Sprintf("%s %s NOT NULL DEFAULT %s", td, enumType, enumValues[0])

			} else {

				var parts []string

				for _, item := range strings.Split(ts, ";") {

					for _, p := range strings.Split(item, "-") {
						parts = append(parts, p)
					}

				}

				columnDef = fmt.Sprintf("%s %s", td, strings.Join(parts, " "))
			}

			columns = append(columns, columnDef)

			// ==========================
			// RELATION
			// ==========================
			rel := field.Tag.Get("relation")
			if rel != "" {

				r := strings.Split(rel, ";")

				if len(r) != 2 {
					log.Fatalf("invalid relation tag: %s", rel)
				}

				ref := strings.Split(r[0], ".")

				if len(ref) != 2 {
					log.Fatalf("invalid relation reference: %s", rel)
				}

				constraints = append(constraints,
					fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s(%s) ON DELETE %s", td, ref[0], ref[1], strings.ToUpper(r[1])),
				)
			}
		}

		args := strings.Join(append(columns, constraints...), ", ")
		query := fmt.Sprintf("%s CREATE TABLE IF NOT EXISTS %s (%s)", createTypeSQL, strings.ToLower(textformater.ToSnakeCase(t.Name())), args)

		if _, err := q.db.Exec(query); err != nil {

			if strings.Contains(err.Error(), "does not exist") ||
				strings.Contains(err.Error(), "not found") {

				fmt.Printf("⚠️  Skipping %s: %v\n", strings.ToLower(t.Name()), err)
				continue
			}

			fmt.Println(query)
			log.Fatalf("exec %s: %v", strings.ToLower(t.Name()), err)
		}
	}

	fmt.Println("Migration completed.")
}
