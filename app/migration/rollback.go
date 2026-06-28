package migration

import (
	"database/sql"

	"github.com/MnPutrav2/go_architecture/app/model"
	"github.com/MnPutrav2/go_architecture/app/pkg/query"
)

func Rollback(db *sql.DB) {
	query.InitDB(db).Rollback(
		model.Users{},
	)
}
