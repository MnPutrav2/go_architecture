package query

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type InitQuery struct {
	db    *sql.DB
	query string
}

func Query(db *sql.DB) *InitQuery {
	return &InitQuery{db: db}
}

func (q *InitQuery) Insert(T any, table string, D any) *InitQuery {

	val := reflect.TypeOf(T)
	field := val.NumField()
	val2 := reflect.ValueOf(D)

	var args string
	var v string

	for i := range field {
		if i == val.NumField()-1 {
			args += val.Field(i).Tag.Get("json")
			x := val2.Field(i).Interface()
			switch x.(type) {
			case int, bool:
				v += fmt.Sprintf("%d", val2.Field(i).Int())
			case string, time.Time, uuid.UUID:
				v += fmt.Sprintf("'%s'", val2.Field(i).String())
			}
		} else {
			args += val.Field(i).Tag.Get("json") + ", "
			x := val2.Field(i).Interface()
			switch x.(type) {
			case int, bool:
				v += fmt.Sprintf("%d, ", val2.Field(i).Int())
			case string, time.Time, uuid.UUID:
				v += fmt.Sprintf("'%s', ", val2.Field(i).String())
			}
		}
	}

	q.query = fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", table, args, v)
	return q
}

func (q *InitQuery) Delete(table string) *InitQuery {
	q.query = fmt.Sprintf("DELETE FROM %s", table)
	return q
}

func (q *InitQuery) Where(que string, arg ...string) *InitQuery {

	cut := strings.Split(que, ",")
	if len(cut) == 1 {
		q.query += fmt.Sprintf(" WHERE %s = %s", que, arg[0])
	} else {
		q.query += fmt.Sprintf(" WHERE %s = %s", cut[0], arg[0])
		for i := range cut {
			if i == 0 {
				continue
			}
			q.query += fmt.Sprintf(" AND %s = %s", cut[i], arg[i])
		}
	}

	return q
}

func (q *InitQuery) Like(que string, arg string) *InitQuery {
	q.query += fmt.Sprintf("%s LIKE '%s'", que, "%"+arg+"%")
	return q
}

func (q *InitQuery) And() *InitQuery {
	q.query += " AND "
	return q
}

func (q *InitQuery) Exec(ctx context.Context) error {

	fmt.Println(q.query)
	if _, err := q.db.ExecContext(ctx, q.query); err != nil {
		return err
	}

	return nil

}

func (q *InitQuery) Read() {
	fmt.Println(q.query)
}
