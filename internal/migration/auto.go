package migration

import (
	"database/sql"

	"github.com/MnPutrav2/go_architecture/internal/model"
	"github.com/MnPutrav2/go_architecture/pkg/query"
)

func Auto(db *sql.DB) {
	query.InitDB(db).Migrate(
		model.Users{},
	)
}
