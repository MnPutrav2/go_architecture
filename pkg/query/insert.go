package query

import (
	"fmt"
	"reflect"
	"strings"
)

func (q *InitQuery[T]) Insert(D any) *InitQuery[T] {
	val := reflect.ValueOf(D)
	typ := reflect.TypeOf(D)

	t := reflect.TypeOf(q.model)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	var columns []string
	var placeholders []string
	var values []any

	// 🔹 Jika slice / array
	if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
		if val.Len() == 0 {
			return q
		}

		elemType := typ.Elem()
		if elemType.Kind() == reflect.Pointer {
			elemType = elemType.Elem()
		}

		// ambil kolom dari struct pertama
		for i := 0; i < elemType.NumField(); i++ {
			columns = append(columns, elemType.Field(i).Tag.Get("db"))
		}

		counter := 1

		for i := 0; i < val.Len(); i++ {
			row := val.Index(i)
			if row.Kind() == reflect.Pointer {
				row = row.Elem()
			}

			var rowPlaceholders []string

			for j := 0; j < row.NumField(); j++ {
				values = append(values, row.Field(j).Interface())
				rowPlaceholders = append(rowPlaceholders, fmt.Sprintf("$%d", counter))
				counter++
			}

			placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ",")))
		}

		q.query = fmt.Sprintf(
			"INSERT INTO %s(%s) VALUES %s",
			strings.ToLower(t.Name()),
			strings.Join(columns, ", "),
			strings.Join(placeholders, ", "),
		)

		q.value = values
		return q
	}

	// 🔹 Jika single struct
	if typ.Kind() == reflect.Pointer {
		val = val.Elem()
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		columns = append(columns, typ.Field(i).Tag.Get("db"))
		values = append(values, val.Field(i).Interface())
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	q.query = fmt.Sprintf(
		"INSERT INTO %s(%s) VALUES(%s)",
		strings.ToLower(t.Name()),
		strings.Join(columns, ", "),
		strings.Join(placeholders, ","),
	)

	q.value = values
	return q
}
