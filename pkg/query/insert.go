package query

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (q *InitQuery[T]) Insert(D any) *InitQuery[T] {

	val := reflect.TypeOf(q.model)
	field := val.NumField()
	val2 := reflect.ValueOf(D)

	t := reflect.TypeOf(q.model)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	var args string
	var v string

	for i := range field {
		if i == val.NumField()-1 {
			args += val.Field(i).Tag.Get("db")
			x := val2.Field(i).Interface()
			switch x.(type) {
			case int, bool:
				v += fmt.Sprintf("%d", val2.Field(i).Int())
			case string, time.Time, uuid.UUID:
				v += fmt.Sprintf("'%s'", val2.Field(i).String())
			}
		} else {
			args += val.Field(i).Tag.Get("db") + ", "
			x := val2.Field(i).Interface()
			switch x.(type) {
			case int, bool:
				v += fmt.Sprintf("%d, ", val2.Field(i).Int())
			case string, time.Time, uuid.UUID:
				v += fmt.Sprintf("'%s', ", val2.Field(i).String())
			}
		}
	}

	q.query = fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", strings.ToLower(t.Name()), args, v)
	return q
}
