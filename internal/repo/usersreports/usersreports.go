package usersreports

import (
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/usersreports"
	"github.com/imperatorofdwelling/Full-backend/pkg/staysChecker"
	"golang.org/x/net/context"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) CreateUsersReports(ctx context.Context, userId, toBlameId, title, description string) error {
	const op = "repo.UsersReports.CreateUsersReports"

	// Generate a new UUID
	id, _ := uuid.NewV4()

	// Checking user for existence
	exists, err := staysChecker.CheckUserExists(ctx, r.Db, toBlameId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: user does not exist: %s", op, toBlameId)
	}

	stmt, err := r.Db.PrepareContext(ctx, "INSERT INTO users_reports (id, user_id, owner_id, title, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id, userId, toBlameId, title, description)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) GetAllUsersReports(ctx context.Context, userId string) ([]usersreports.UsersReportEntity, error) {
	const op = "repo.UsersReports.GetAllUsersReports"

	// Prepare the SQL statement to fetch user reports, including report ID
	stmt, err := r.Db.PrepareContext(ctx, "SELECT ur.id AS report_id, u1.name AS user_name, u2.name AS owner_name, ur.title AS title, ur.description AS description FROM users_reports ur INNER JOIN users u1 ON ur.user_id = u1.id INNER JOIN users u2 ON ur.owner_id = u2.id WHERE ur.user_id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	// Execute the query
	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var reports []usersreports.UsersReportEntity

	for rows.Next() {
		var report usersreports.UsersReportEntity
		if err := rows.Scan(&report.ReportID, &report.UserName, &report.OwnerName, &report.Title, &report.Description); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		reports = append(reports, report)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return reports, nil
}

func (r *Repo) UpdateUsersReports(ctx context.Context, userId, reportId, title, description string) error {
	const op = "repo.UsersReports.UpdateUsersReports"

	// Checking User existence
	exists, err := staysChecker.CheckUsersReportExists(ctx, r.Db, reportId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: report does not exist: %s", op, reportId)
	}

	// Preparing query
	stmt, err := r.Db.PrepareContext(ctx, "UPDATE users_reports SET title = $1, description = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3 AND user_id = $4")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, title, description, reportId, userId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Check if any rows were updated
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("%s: no report found for user: %s and owner: %s", op, userId, reportId)
	}

	return nil
}

func (r *Repo) DeleteUsersReports(ctx context.Context, userId, reportId string) error {
	const op = "repo.UsersReports.DeleteUsersReports"

	// Checking User existence
	exists, err := staysChecker.CheckUsersReportExists(ctx, r.Db, reportId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: report does not exist: %s", op, reportId)
	}

	// Check if the report exists
	stmt, err := r.Db.PrepareContext(ctx, "DELETE FROM users_reports WHERE user_id = $1 AND id = $2")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userId, reportId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
