package staysreports

import (
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreports"
	"github.com/imperatorofdwelling/Full-backend/pkg/checkers"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) CreateStaysReports(ctx context.Context, userId, stayId, title, description, imagePath string) error {
	const op = "repo.StaysReports.CreateStaysReports"

	id, _ := uuid.NewV4()

	// Check if the user exists
	userExists, err := checkers.CheckUserExists(ctx, r.Db, userId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !userExists {
		return fmt.Errorf("%s: user does not exist: %s", op, userId)
	}

	// Check if the stay exists
	stayExists, err := checkers.CheckStayExists(ctx, r.Db, stayId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !stayExists {
		return fmt.Errorf("%s: stay does not exist: %s", op, stayId)
	}

	// Prepare and execute the INSERT query
	stmt, err := r.Db.PrepareContext(ctx, "INSERT INTO stays_reports (id, user_id, stay_id, title, description, report_attach, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id, userId, stayId, title, description, imagePath)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) GetAllStaysReports(ctx context.Context, userId string) ([]staysreports.StaysReportEntity, error) {
	const op = "repo.StaysReports.GetAllStaysReports"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT sr.id, u.name AS user_name, s.name AS stay_name, sr.title, sr.description FROM stays_reports sr INNER JOIN users u ON sr.user_id = u.id INNER JOIN stays s ON sr.stay_id = s.id WHERE sr.user_id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var reports []staysreports.StaysReportEntity

	for rows.Next() {
		var report staysreports.StaysReportEntity
		if err := rows.Scan(&report.ReportID, &report.UserName, &report.StayName, &report.Title, &report.Description); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		reports = append(reports, report)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return reports, nil
}

func (r *Repo) GetStaysReportById(ctx context.Context, userId, id string) (*staysreports.StayReport, error) {
	const op = "repo.StaysReports.GetStaysReportById"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT id, user_id, stay_id, title, description, report_attach, created_at, updated_at FROM stays_reports WHERE user_id = $1 AND id = $2")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, userId, id)

	var report staysreports.StayReport

	err = row.Scan(
		&report.ID,
		&report.UserID,
		&report.StayID,
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

func (r *Repo) UpdateStaysReports(ctx context.Context, userId, reportId, title, description, updatedImagePath string) (*staysreports.StaysReportEntity, error) {
	const op = "repo.StaysReports.UpdateStaysReports"

	// Checking stay for existence
	exists, err := checkers.CheckStaysReportExists(ctx, r.Db, reportId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return nil, fmt.Errorf("%s: report does not exist: %s", op, reportId)
	}

	// Preparing update statement
	updateStmt, err := r.Db.PrepareContext(ctx, "UPDATE stays_reports SET title = $1, description = $2, report_attach = $3, updated_at = CURRENT_TIMESTAMP WHERE user_id = $4 AND id = $5")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer updateStmt.Close()

	// Executing update statement
	_, err = updateStmt.ExecContext(ctx, title, description, updatedImagePath, userId, reportId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Preparing query to find recently updated report
	selectStmt, err := r.Db.PrepareContext(ctx, "SELECT sr.id, u.name AS user_name, s.name AS stay_name, sr.title, sr.description FROM stays_reports sr INNER JOIN users u ON sr.user_id = u.id INNER JOIN stays s ON sr.stay_id = s.id WHERE sr.user_id = $1 AND sr.id = $2")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer selectStmt.Close()

	// Executing it
	row := selectStmt.QueryRowContext(ctx, userId, reportId)

	var report staysreports.StaysReportEntity
	err = row.Scan(&report.ReportID, &report.UserName, &report.StayName, &report.Title, &report.Description)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &report, nil
}

func (r *Repo) DeleteStaysReports(ctx context.Context, userId, reportId string) error {
	const op = "repo.StaysReports.DeleteStaysReports"

	exists, err := checkers.CheckStaysReportExists(ctx, r.Db, reportId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: report does not exist: %s", op, reportId)
	}

	// Prepare the SQL statement to delete based on reportId
	stmt, err := r.Db.PrepareContext(ctx, "DELETE FROM stays_reports WHERE id = $1")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	// Execute the delete statement with reportId as the parameter
	_, err = stmt.ExecContext(ctx, reportId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
