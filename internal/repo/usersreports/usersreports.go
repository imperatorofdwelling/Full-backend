package usersreports

import (
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/usersreports"
	"github.com/imperatorofdwelling/Full-backend/pkg/checkers"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) CreateUsersReports(ctx context.Context, userId, toBlameId, title, description, imagePath string) error {
	const op = "repo.UsersReports.CreateUsersReports"

	// Generate a new UUID
	id, _ := uuid.NewV4()

	// Checking user for existence
	exists, err := checkers.CheckUserExists(ctx, r.Db, toBlameId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: user does not exist: %s", op, toBlameId)
	}

	stmt, err := r.Db.PrepareContext(ctx, "INSERT INTO users_reports (id, user_id, owner_id, title, description, report_attach, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id, userId, toBlameId, title, description, imagePath)
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

func (r *Repo) GetUsersReportById(ctx context.Context, userId, id string) (*usersreports.UsersReport, error) {
	const op = "repo.UsersReports.GetUsersReportById"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT id, user_id, owner_id, title, description, report_attach, created_at, updated_at FROM users_reports WHERE user_id = $1 AND id = $2")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, userId, id)

	var report usersreports.UsersReport

	err = row.Scan(
		&report.ID,
		&report.UserID,
		&report.OwnerID,
		&report.Title,
		&report.Description,
		&report.ReportAttach,
		&report.CreatedAt,
		&report.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: no report found for user_id %s: %w", op, id, err)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &report, nil
}

func (r *Repo) UpdateUsersReports(ctx context.Context, userId, reportId, title, description, updatedImagePath string) (*usersreports.UsersReportEntity, error) {
	const op = "repo.UsersReports.UpdateUsersReports"
	fmt.Printf("userId: %s, reportId: %s\n", userId, reportId)

	// Checking User existence
	exists, err := checkers.CheckUsersReportExists(ctx, r.Db, reportId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return nil, fmt.Errorf("%s: report does not exist: %s", op, reportId)
	}

	// Preparing query
	stmt, err := r.Db.PrepareContext(ctx, "UPDATE users_reports SET title = $1, description = $2, report_attach = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4 AND user_id = $5")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, title, description, updatedImagePath, reportId, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Check if any rows were updated
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("%s: no report found for user: %s and owner: %s", op, userId, reportId)
	}

	// Prepare the SQL statement to fetch user report
	selectStmt, err := r.Db.PrepareContext(ctx, "SELECT ur.id AS report_id, u1.name AS user_name, u2.name AS owner_name, ur.title AS title, ur.description AS description FROM users_reports ur INNER JOIN users u1 ON ur.user_id = u1.id INNER JOIN users u2 ON ur.owner_id = u2.id WHERE ur.user_id = $1 AND ur.id = $2")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer selectStmt.Close()

	row := selectStmt.QueryRowContext(ctx, userId, reportId)

	// Scanning results into UsersReportEntity
	var report usersreports.UsersReportEntity
	err = row.Scan(&report.ReportID, &report.UserName, &report.OwnerName, &report.Title, &report.Description)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &report, nil
}

func (r *Repo) DeleteUsersReports(ctx context.Context, userId, reportId string) error {
	const op = "repo.UsersReports.DeleteUsersReports"

	// Checking User existence
	exists, err := checkers.CheckUsersReportExists(ctx, r.Db, reportId)
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
