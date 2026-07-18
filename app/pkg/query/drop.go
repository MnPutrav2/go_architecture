package query

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	textformater "github.com/MnPutrav2/go_architecture/app/pkg/text_formater"
)

func (q *Initdb) Rollback(table ...any) {
	fmt.Println("Running rollback... ==================")
	for _, tbl := range table {
		t := reflect.TypeOf(tbl)

		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		if t.Kind() != reflect.Struct {
			continue
		}

		queries := []string{}

		queries = append(
			queries,
			fmt.Sprintf(
				"DROP TABLE IF EXISTS %s CASCADE",
				strings.ToLower(textformater.ToSnakeCase(t.Name())),
			),
		)

		for i := 0; i < t.NumField(); i++ {
			ts := t.Field(i).Tag.Get("structure")

			if strings.Contains(ts, "enum") {
				queries = append(
					queries,
					fmt.Sprintf(
						"DROP TYPE IF EXISTS %s_ty",
						strings.ToLower(textformater.ToSnakeCase(t.Name())),
					),
				)
			}
		}

		query := strings.Join(queries, "; ")
		if _, err := q.db.Exec(query); err != nil {
			if strings.Contains(err.Error(), "does not exist") ||
				strings.Contains(err.Error(), "not found") {
				fmt.Printf("⚠️  Skipping %s: %v\n", strings.ToLower(textformater.ToSnakeCase(t.Name())), err)
				continue
			} else {
				fmt.Println(query)
				log.Fatalf("exec %s: %v", strings.ToLower(textformater.ToSnakeCase(t.Name())), err)
			}
		}
	}

	fmt.Println("Rollback completed.")
}
