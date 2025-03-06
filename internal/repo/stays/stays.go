package stays

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"github.com/lib/pq"
	"sort"
	"strings"
	"time"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) CreateStay(ctx context.Context, stay *models.StayEntity) error {
	const op = "repo.stays.CreateStay"

	stmt, err := r.Db.PrepareContext(ctx,
		"INSERT INTO stays(user_id, location_id, name, type, number_of_bedrooms, number_of_beds, number_of_bathrooms, guests, is_smoking_prohibited, square, street, house, entrance, floor, room, price, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, stay.UserID, stay.LocationID, stay.Name, stay.Type, stay.NumberOfBedrooms, stay.NumberOfBeds, stay.NumberOfBathrooms, stay.Guests, stay.IsSmokingProhibited, stay.Square, stay.Street, stay.House, stay.Entrance, stay.Floor, stay.Room, stay.Price, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) GetStayByID(ctx context.Context, id uuid.UUID) (*models.Stay, error) {
	const op = "repo.stays.getStayByID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM stays WHERE id=$1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var stay models.Stay

	row := stmt.QueryRowContext(ctx, id)
	if row.Err() != nil {
		return nil, fmt.Errorf("%s: %w", op, row.Err())
	}

	var amenitiesData []byte // Переменная для хранения JSONB данных

	err = row.Scan(&stay.ID, &stay.LocationID, &stay.UserID, &stay.Name, &stay.Type, &stay.NumberOfBedrooms, &stay.NumberOfBeds, &stay.NumberOfBathrooms, &stay.Guests, &stay.Rating, &amenitiesData, &stay.IsSmokingProhibited, &stay.Square, &stay.Street, &stay.House, &stay.Entrance, &stay.Floor, &stay.Room, &stay.Price, &stay.CreatedAt, &stay.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	// Десериализуем JSONB данные в []Amenity
	if err = json.Unmarshal(amenitiesData, &stay.Amenities); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &stay, nil
}

func (r *Repo) GetStays(ctx context.Context) ([]models.StayResponse, error) {
	const op = "repo.stays.getStays"

	stmt, err := r.Db.PrepareContext(ctx, `
		SELECT *
		FROM stays
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	var stays []models.StayResponse

	for rows.Next() {
		var stay models.StayResponse

		var amenitiesData []byte // Переменная для хранения JSONB данных

		err = rows.Scan(&stay.ID, &stay.UserID, &stay.LocationID, &stay.Name, &stay.Type, &stay.NumberOfBedrooms, &stay.NumberOfBeds, &stay.NumberOfBathrooms, &stay.Guests, &stay.Rating, &amenitiesData, &stay.IsSmokingProhibited, &stay.Square, &stay.Street, &stay.House, &stay.Entrance, &stay.Floor, &stay.Room, &stay.Price, &stay.CreatedAt, &stay.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		// Десериализуем JSONB данные в []Amenity
		if err = json.Unmarshal(amenitiesData, &stay.Amenities); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		images, err := r.GetImagesByStayID(ctx, stay.ID)
		if err != nil {
			return nil, err
		}

		stay.Images = images

		stays = append(stays, stay)
	}

	return stays, nil
}

func (r *Repo) UpdateStayByID(ctx context.Context, stay *models.StayEntity, id uuid.UUID) error {
	const op = "repo.stays.updateStayByID"

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE stays SET location_id=$1, name=$2, type=$5, number_of_bedrooms=$6, number_of_beds=$7, number_of_bathrooms=$8, guests=$9, is_smoking_prohibited=$10, square=$11, street=$12, house=$13, entrance=$14, floor=$15, room=$16, price=$17, updated_at=$18 WHERE id=$19")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		stay.LocationID,
		stay.Name,
		stay.Type,
		stay.NumberOfBedrooms,
		stay.NumberOfBeds,
		stay.NumberOfBathrooms,
		stay.Guests,
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
	var amenitiesData []byte // Переменная для хранения JSONB данных
	for rows.Next() {
		var stay models.Stay

		err = rows.Scan(
			&stay.ID,
			&stay.LocationID,
			&stay.UserID,
			&stay.Name,
			&stay.Type,
			&stay.NumberOfBedrooms,
			&stay.NumberOfBeds,
			&stay.NumberOfBathrooms,
			&stay.Guests,
			&stay.Rating,
			&amenitiesData,
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
		// Десериализуем JSONB данные в []Amenity
		if err = json.Unmarshal(amenitiesData, &stay.Amenities); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		stays = append(stays, &stay)
	}

	return stays, nil
}

func (r *Repo) GetImagesByStayID(ctx context.Context, id uuid.UUID) ([]models.StayImage, error) {
	const op = "repo.stays.GetImagesByStayID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM stays_images WHERE stay_id=$1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	rowsImg, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var stayImages []models.StayImage

	for rowsImg.Next() {
		var stayImage models.StayImage
		err = rowsImg.Scan(&stayImage.ID, &stayImage.StayID, &stayImage.ImageName, &stayImage.IsMain, &stayImage.CreatedAt, &stayImage.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		stayImages = append(stayImages, stayImage)
	}

	return stayImages, nil
}

func (r *Repo) GetMainImageByStayID(ctx context.Context, id uuid.UUID) (models.StayImage, error) {
	const op = "repo.stays.GetMainImageByStayID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM stays_images WHERE stay_id=$1 AND is_main=$2")
	if err != nil {
		return models.StayImage{}, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var stayImage models.StayImage

	row := stmt.QueryRowContext(ctx, id, true)

	err = row.Scan(&stayImage.ID, &stayImage.StayID, &stayImage.ImageName, &stayImage.IsMain, &stayImage.UpdatedAt, &stayImage.CreatedAt)
	if err != nil {
		return models.StayImage{}, fmt.Errorf("%s: %w", op, err)
	}

	return stayImage, nil
}

func (r *Repo) GetStayImageByID(ctx context.Context, imageID uuid.UUID) (models.StayImage, error) {
	const op = "repo.stays.GetStayImageByID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM stays_images WHERE id=$1")
	if err != nil {
		return models.StayImage{}, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var stayImage models.StayImage

	row := stmt.QueryRowContext(ctx, imageID)
	err = row.Scan(
		&stayImage.ID,
		&stayImage.StayID,
		&stayImage.ImageName,
		&stayImage.IsMain,
		&stayImage.UpdatedAt,
		&stayImage.CreatedAt)
	if err != nil {
		return models.StayImage{}, fmt.Errorf("%s: %w", op, err)
	}

	return stayImage, nil
}

func (r *Repo) CreateStayImage(ctx context.Context, fileName string, isMain bool, stayID uuid.UUID) error {
	const op = "repo.stays.CreateStayImage"

	stmt, err := r.Db.PrepareContext(ctx, "INSERT INTO stays_images(image_name, stay_id, is_main, created_at, updated_at) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, fileName, stayID, isMain, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) DeleteStayImage(ctx context.Context, imageId uuid.UUID) error {
	const op = "repo.stays.DeleteStayImage"

	stmt, err := r.Db.PrepareContext(ctx, "DELETE FROM stays_images WHERE id=$1")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, imageId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) GetStaysByLocationID(ctx context.Context, id uuid.UUID) (*[]models.Stay, error) {
	const op = "repo.stays.GetStaysByLocationID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM stays WHERE location_id=$1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var stays []models.Stay
	var amenitiesData []byte // Переменная для хранения JSONB данных
	for rows.Next() {
		var stay models.Stay

		err = rows.Scan(
			&stay.ID,
			&stay.LocationID,
			&stay.UserID,
			&stay.Name,
			&stay.Type,
			&stay.NumberOfBedrooms,
			&stay.NumberOfBeds,
			&stay.NumberOfBathrooms,
			&stay.Guests,
			&stay.Rating,
			&amenitiesData,
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

		// Десериализуем JSONB данные в []Amenity
		if err = json.Unmarshal(amenitiesData, &stay.Amenities); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		stays = append(stays, stay)
	}

	return &stays, nil
}

func (this *Repo) Filtration(ctx context.Context, search models.Filtration, locationIDS []uuid.UUID) ([]models.Stay, error) {
	const op = "repo.stays.Filtration"

	sort.Float64s(search.Rating)

	// Находим минимальный и максимальный рейтинг
	var minRating, maxRating float64
	if len(search.Rating) > 0 {
		minRating = search.Rating[0]
		maxRating = search.Rating[0]
		for _, r := range search.Rating {
			if r < minRating {
				minRating = r
			}
			if r > maxRating {
				maxRating = r
			}
		}
	}

	query := `
		SELECT * FROM stays
		WHERE location_id = ANY($1) 
	`

	args := []interface{}{
		pq.Array(locationIDS),
	}
	count := 1
	if search.PriceMin != -1 && search.PriceMax != -1 {
		count++
		nextCount := count + 1
		query += fmt.Sprintf(" AND price BETWEEN $%d AND $%d", count, nextCount)
		args = append(args, fmt.Sprintf("%f", search.PriceMin), fmt.Sprintf("%f", search.PriceMax))
		count++
	}
	if len(search.NumberOfBedrooms) > 0 {
		count++
		query += fmt.Sprintf(" AND number_of_bedrooms = ANY($%d)", count)
		args = append(args, pq.Array(search.NumberOfBedrooms))
	}
	if len(search.Rating) > 0 {
		count++
		nextCount := count + 1
		query += fmt.Sprintf(" AND rating BETWEEN $%d AND $%d", count, nextCount)
		args = append(args, fmt.Sprintf("%f", search.Rating[0]), fmt.Sprintf("%f", search.Rating[1]))
		count++
	}

	if len(search.Amenities) > 0 {
		//Создаем временный массив для условий
		var conditions []string
		for key, value := range search.Amenities {
			// Используем индекс для параметров
			conditions = append(conditions, fmt.Sprintf("amenities ->> '%s' = $%d", key, len(args)+1))
			args = append(args, fmt.Sprintf("%t", value))
		}
		// Объединяем условия с помощью AND
		query += " AND (" + strings.Join(conditions, " AND ") + ")"
	}

	stmt, err := this.Db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	// Выполняем запрос
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var amenitiesData []byte // Переменная для хранения JSONB данных

	// Обработка результатов
	var stays []models.Stay
	for rows.Next() {
		var stay models.Stay
		if err = rows.Scan(
			&stay.ID,
			&stay.LocationID,
			&stay.UserID,
			&stay.Name,
			&stay.Type,
			&stay.NumberOfBedrooms,
			&stay.NumberOfBeds,
			&stay.NumberOfBathrooms,
			&stay.Guests,
			&stay.Rating,
			&amenitiesData,
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
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		// Десериализуем JSONB данные в []Amenity
		if err = json.Unmarshal(amenitiesData, &stay.Amenities); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		stays = append(stays, stay)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return stays, nil
}
