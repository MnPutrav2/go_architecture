package query

import (
	"fmt"
	"reflect"
	"strings"
)

func (q *InitQuery[T]) Insert(D any) *InitQuery[T] {

	val2 := reflect.TypeOf(D)
	val3 := reflect.ValueOf(D)

	t := reflect.TypeOf(q.model)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	var args []string
	var pre []string
	var v []any

	for i := range val2.NumField() {
		args = append(args, val2.Field(i).Tag.Get("db"))
		v = append(v, val3.Field(i).Interface())
		pre = append(pre, fmt.Sprintf("$%d", i+1))
	}

	q.value = v
	q.query = fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)",
		strings.ToLower(t.Name()),
		strings.Join(args, ", "),
		strings.Join(pre, ","),
	)

	return q
}
