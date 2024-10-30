package staysreports

import (
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreports"
	"github.com/imperatorofdwelling/Full-backend/pkg/staysChecker"
	"golang.org/x/net/context"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) CreateStaysReports(ctx context.Context, userId, stayId, title, description string) error {
	const op = "repo.StaysReports.CreateStaysReports"

	id, _ := uuid.NewV4()
	// Checking stay for existence
	exists, err := staysChecker.CheckStayExists(ctx, r.Db, stayId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: stay does not exist: %s", op, stayId)
	}

	stmt, err := r.Db.PrepareContext(ctx, "INSERT INTO stays_reports (id, user_id, stay_id, title, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, id, userId, stayId, title, description)
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

func (r *Repo) UpdateStaysReports(ctx context.Context, userId, reportId, title, description string) error {
	const op = "repo.StaysReports.UpdateStaysReports"

	// Checking stay for existence
	exists, err := staysChecker.CheckStaysReportExists(ctx, r.Db, reportId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: report does not exist: %s", op, reportId)
	}

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE stays_reports SET title = $1, description = $2, updated_at = CURRENT_TIMESTAMP WHERE user_id = $3 AND id = $4")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, title, description, userId, reportId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) DeleteStaysReports(ctx context.Context, userId, reportId string) error {
	const op = "repo.StaysReports.DeleteStaysReports"

	exists, err := staysChecker.CheckStaysReportExists(ctx, r.Db, reportId)
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
