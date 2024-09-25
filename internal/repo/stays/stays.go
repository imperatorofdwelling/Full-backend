package stays

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"time"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) CreateStay(ctx context.Context, stay *models.StayEntity) error {
	const op = "repo.stays.CreateStay"

	stmt, err := r.Db.PrepareContext(ctx,
		"INSERT INTO stays(user_id, location_id, name, image_main, images, type, number_of_bedrooms, number_of_beds, number_of_bathrooms, guests, rating, is_smoking_prohibited, square, street, house, entrance, floor, room, price, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, stay.UserID, stay.LocationID, stay.Name, stay.ImageMain, stay.Images, stay.Type, stay.NumberOfBedrooms, stay.NumberOfBeds, stay.NumberOfBathrooms, stay.Guests, stay.Rating, stay.IsSmokingProhibited, stay.Square, stay.Street, stay.House, stay.Entrance, stay.Floor, stay.Room, stay.Price, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) GetStayByID(ctx context.Context, id uuid.UUID) (*models.Stay, error) {
	const op = "repo.stays.getStayByID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT 1 FROM stays WHERE id=$1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var stay models.Stay

	row := stmt.QueryRowContext(ctx, id)
	if row.Err() != nil {
		return nil, fmt.Errorf("%s: %w", op, row.Err())
	}

	err = row.Scan(&stay.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &stay, nil
}

func (r *Repo) GetStays(ctx context.Context) ([]*models.Stay, error) {
	const op = "repo.stays.getStays"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM stays")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	var stays []*models.Stay

	for rows.Next() {
		var stay models.Stay

		err = rows.Scan(&stay.ID, &stay.LocationID, &stay.UserID, &stay.Name, &stay.ImageMain, &stay.Images, &stay.Type, &stay.NumberOfBedrooms, &stay.NumberOfBeds, &stay.NumberOfBathrooms, &stay.Guests, &stay.Rating, &stay.IsSmokingProhibited, &stay.Square, &stay.Street, &stay.House, &stay.Entrance, &stay.Floor, &stay.Room, &stay.Price, &stay.CreatedAt, &stay.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		stays = append(stays, &stay)
	}

	return stays, nil
}

func (r *Repo) UpdateStayByID(ctx context.Context, stay *models.StayEntity, id uuid.UUID) error {
	const op = "repo.stays.updateStayByID"

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE stays SET location_id=$1, name=$2, image_main=$3, images=$4, type=$5, number_of_bedrooms=$6, number_of_beds=$7, number_of_bathrooms=$8, guests=$9, rating=$10, is_smoking_prohibited=$11, square=$12, street=$13, house=$14, entrance=$15, floor=$16, room=$17, price=$18, updated_at=$19 WHERE id=$20")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		stay.LocationID,
		stay.Name,
		stay.ImageMain,
		stay.Images,
		stay.Type,
		stay.NumberOfBedrooms,
		stay.NumberOfBeds,
		stay.NumberOfBathrooms,
		stay.Guests,
		stay.Rating,
		stay.IsSmokingProhibited,
		stay.Square,
		stay.Street,
		stay.House,
		stay.Entrance,
		stay.Floor,
		stay.Room,
		stay.Price,
		time.Now(),
		id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) DeleteStayByID(ctx context.Context, id uuid.UUID) error {
	const op = "repo.stays.DeleteStayByID"

	stmt, err := r.Db.PrepareContext(ctx, "DELETE FROM stays WHERE id=$1")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) CheckStayIfExistsByID(ctx context.Context, id uuid.UUID) (bool, error) {
	const op = "repo.stays.CheckStayIfExistsByID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT EXISTS(SELECT 1 FROM stays WHERE id=$1)")
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var exists bool

	err = stmt.QueryRowContext(ctx, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return exists, nil
}

func (r *Repo) GetStaysByUserID(ctx context.Context, userId uuid.UUID) ([]*models.Stay, error) {
	const op = "repo.stays.GetStaysByUserID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM stays WHERE user_id=$1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	var stays []*models.Stay

	for rows.Next() {
		var stay models.Stay

		err = rows.Scan(
			&stay.ID,
			&stay.LocationID,
			&stay.UserID,
			&stay.Name,
			&stay.ImageMain,
			&stay.Images,
			&stay.Type,
			&stay.NumberOfBedrooms,
			&stay.NumberOfBeds,
			&stay.NumberOfBathrooms,
			&stay.Guests,
			&stay.Rating,
			&stay.IsSmokingProhibited,
			&stay.Square,
			&stay.Street,
			&stay.House,
			&stay.Entrance,
			&stay.Floor,
			&stay.Room,
			&stay.Price,
			&stay.CreatedAt,
			&stay.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		stays = append(stays, &stay)
	}

	return stays, nil
}