package staysChecker

import (
	"database/sql"
	"golang.org/x/net/context"
)

func CheckStayExists(ctx context.Context, db *sql.DB, stayID string) (bool, error) {
	var exists bool
	err := db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM stays WHERE id = $1)", stayID).Scan(&exists)
	return exists, err
}
