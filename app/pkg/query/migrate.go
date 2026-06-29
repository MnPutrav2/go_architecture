package query

import (
	"fmt"
	"log"
	"reflect"
	"strings"
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

		var ty string
		var args string
		for i := 0; i < t.NumField(); i++ {

			ts := t.Field(i).Tag.Get("structure")
			td := t.Field(i).Tag.Get("db")
			var x string

			if strings.Contains(ts, "enum") {
				var n []string
				s := strings.Split(ts, "(")
				c := strings.SplitSeq(strings.TrimSpace(strings.ReplaceAll(s[1], ")", "")), ",")

				for m := range c {
					n = append(n, fmt.Sprintf(`'%s'`, m))
				}

				x += fmt.Sprintf("%s_ty NOT NULL DEFAULT %s", strings.ToLower(t.Name()), n[0])
				ty += fmt.Sprintf(`CREATE TYPE %s_ty AS ENUM %s; `, strings.ToLower(t.Name()), "("+strings.Join(n, ",")+")")
			}

			cut := strings.Split(ts, ";")

			if !strings.Contains(ts, "enum") {
				for l := range len(cut) {
					c := strings.Split(cut[l], "-")

					for n := range len(c) {
						x += fmt.Sprintf("%s ", c[n])
					}
				}
			}

			if i == t.NumField()-1 {
				args += fmt.Sprintf("%s %s", strings.ToLower(td), x)
			} else {
				args += fmt.Sprintf("%s %s,", strings.ToLower(td), x)
			}

			st := t.Field(i).Tag.Get("relation")
			if st != "" {
				ct := strings.Split(st, ";")
				args += fmt.Sprintf(", FOREIGN KEY (%s) REFERENCES %s ON DELETE %s", td, ct[0], strings.ToUpper(ct[1]))
			}
		}

		query := fmt.Sprintf(`%s CREATE TABLE IF NOT EXISTS %s (%s)`, ty, strings.ToLower(t.Name()), args)
		if _, err := q.db.Exec(query); err != nil {
			if strings.Contains(err.Error(), "does not exist") ||
				strings.Contains(err.Error(), "not found") {
				fmt.Printf("⚠️  Skipping %s: %v\n", strings.ToLower(t.Name()), err)
				continue
			} else {
				fmt.Println(query)
				log.Fatalf("exec %s: %v", strings.ToLower(t.Name()), err)
			}
		}
	}

	fmt.Println("Migration completed.")
}
