package query

import (
	"fmt"
	"reflect"
	"strings"

	textformater "github.com/MnPutrav2/go_architecture/app/pkg/text_formater"
)

func (q *InitQuery[T]) Select(arg string) *InitQuery[T] {
	t := reflect.TypeOf(q.model)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	var query string
	cut := strings.Split(arg, ",")
	for i, v := range cut {
		if i == t.NumField()-1 {
			query += fmt.Sprintf("%s.%s", strings.ToLower(t.Name()), strings.TrimSpace(v))
		} else {
			query += fmt.Sprintf("%s.%s, ", strings.ToLower(t.Name()), strings.TrimSpace(v))
		}
	}

	q.query += fmt.Sprintf("SELECT %s FROM %s", query, strings.ToLower(textformater.ToSnakeCase(t.Name())))
	return q
}
