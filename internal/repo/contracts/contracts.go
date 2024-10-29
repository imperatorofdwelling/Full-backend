package contracts

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/contracts"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"github.com/imperatorofdwelling/Full-backend/pkg/staysChecker"
	"time"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) AddContract(ctx context.Context, userId, stayId string, dateStart, dateEnd time.Time) error {
	const op = "repo.Contracts.AddContract"

	// Checking stay for existence
	exists, err := staysChecker.CheckStayExists(ctx, r.Db, stayId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: stay does not exist: %s", op, stayId)
	}

	// Check if the user already has a contract for this stay
	var count int
	err = r.Db.QueryRowContext(ctx, "SELECT COUNT(*) FROM contracts WHERE user_id = $1 AND stay_id = $2", userId, stayId).Scan(&count)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if count > 0 {
		return fmt.Errorf("%s: user %s already has a contract for stay %s", op, userId, stayId)
	}

	// Check if the selected dates are in the past
	now := time.Now()
	if dateStart.Before(now) || dateEnd.Before(now) || dateEnd.Before(dateStart) {
		return fmt.Errorf("%s: start date and end date must be in the future", op)
	}

	// Preparing stay query to get price
	selectStmt, err := r.Db.PrepareContext(ctx, "SELECT price, square, street, house, entrance, floor, room FROM stays WHERE id = $1")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer selectStmt.Close()

	var stay models.Stay
	err = selectStmt.QueryRowContext(ctx, stayId).Scan(
		&stay.Price, &stay.Square, &stay.Street, &stay.House,
		&stay.Entrance, &stay.Floor, &stay.Room,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Calculate the number of days between dateStart and dateEnd
	duration := dateEnd.Sub(dateStart)
	days := int(duration.Hours() / 24)

	// Calculate the total price
	totalPrice := float32(days) * stay.Price

	// Preparing contract query to insert new contract
	insertStmt, err := r.Db.PrepareContext(ctx, "INSERT INTO contracts (user_id, stay_id, price, date_start, date_end, square, street, house, entrance, floor, room, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer insertStmt.Close()

	// Executing contract query
	_, err = insertStmt.ExecContext(ctx, userId, stayId, totalPrice, dateStart, dateEnd,
		stay.Square, stay.Street, stay.House, stay.Entrance, stay.Floor, stay.Room,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
func (r *Repo) UpdateContract(ctx context.Context, userId, stayId string, dateStart, dateEnd time.Time) error {
	const op = "repo.Contracts.UpdateContract"

	// Checking stay for existence
	exists, err := staysChecker.CheckStayExists(ctx, r.Db, stayId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: stay does not exist: %s", op, stayId)
	}

	// Check if the selected dates are in the past
	now := time.Now()
	if dateStart.Before(now) || dateEnd.Before(now) {
		return fmt.Errorf("%s: start date and end date must be in the future", op)
	}

	// Check if dateEnd is before dateStart
	if dateEnd.Before(dateStart) {
		return fmt.Errorf("%s: end date must be after start date", op)
	}

	// Preparing query to get the current price for the stay
	var price float64
	priceStmt, err := r.Db.PrepareContext(ctx, "SELECT price FROM stays WHERE id = $1")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer priceStmt.Close()

	// Executing price query
	err = priceStmt.QueryRowContext(ctx, stayId).Scan(&price)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Calculate the number of days between dateStart and dateEnd
	duration := dateEnd.Sub(dateStart)
	days := int(duration.Hours() / 24)

	// Calculate the total price
	totalPrice := float64(days) * price

	// Preparing query for updating the contract
	updateStmt, err := r.Db.PrepareContext(ctx, "UPDATE contracts SET price = $1, date_start = $2, date_end = $3, updated_at = CURRENT_TIMESTAMP WHERE user_id = $4 AND stay_id = $5")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer updateStmt.Close()

	// Executing update query
	_, err = updateStmt.ExecContext(ctx, totalPrice, dateStart, dateEnd, userId, stayId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) GetAllContracts(ctx context.Context, userId string) ([]contracts.ContractEntity, error) {
	const op = "repo.Contracts.GetContractEntities"

	// Preparing query
	stmt, err := r.Db.PrepareContext(ctx, "SELECT u.name AS user_name, s.name AS stay_name, c.price, c.date_start, c.date_end FROM contracts c INNER JOIN users u ON c.user_id = u.id INNER JOIN stays s ON c.stay_id = s.id WHERE c.user_id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	// Executing query
	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var contractEntities []contracts.ContractEntity

	// Getting results
	for rows.Next() {
		var entity contracts.ContractEntity
		err := rows.Scan(&entity.UserName, &entity.StayName, &entity.Price, &entity.DateStart, &entity.DateEnd)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		contractEntities = append(contractEntities, entity)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return contractEntities, nil
}
