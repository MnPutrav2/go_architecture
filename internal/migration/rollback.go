package migration

import (
	"database/sql"

	"github.com/MnPutrav2/go_architecture/internal/model"
	"github.com/MnPutrav2/go_architecture/pkg/query"
)

func Rollback(db *sql.DB) {
	query.InitDB(db).Rollback(
		model.Users{},
	)
}
