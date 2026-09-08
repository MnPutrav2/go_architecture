package migration

import (
	"database/sql"
	"fmt"
)

func Fresh(db *sql.DB) {
	fmt.Println("Running fresh... ==================")
	Rollback(db)
	Migrate(db)
	fmt.Println("Fresh completed.")
}
