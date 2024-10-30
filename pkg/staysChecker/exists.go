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

func CheckUserExists(ctx context.Context, db *sql.DB, userID string) (bool, error) {
	var exists bool
	err := db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", userID).Scan(&exists)
	return exists, err
}

func CheckStaysReportExists(ctx context.Context, db *sql.DB, reportID string) (bool, error) {
	var exists bool
	err := db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM stays_reports WHERE id = $1)", reportID).Scan(&exists)
	return exists, err
}

func CheckUsersReportExists(ctx context.Context, db *sql.DB, reportID string) (bool, error) {
	var exists bool
	err := db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users_reports WHERE id = $1)", reportID).Scan(&exists)
	return exists, err
}
