package stays

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	filtrationSort "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays/sort"
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
		"INSERT INTO stays(user_id, location_id, name, type, guests, amenities, house, entrance, address, rooms_count, beds_count, price, period, owners_rules, cancellation_policy, describe_property, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	amenitiesJSON, err := json.Marshal(stay.Amenities)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, stay.UserID, stay.LocationID, stay.Name, stay.Type, stay.Guests, amenitiesJSON, stay.House, stay.Entrance, stay.Address, stay.RoomsCount, stay.BedsCount, stay.Price, stay.Period, stay.OwnersRules, stay.CancellationPolicy, stay.DescribeProperty, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) GetStatistics(ctx context.Context, userId string) (*models.Statistics, error) {
	const op = "repo.stays.GetStatistics"

	query := `SELECT 
				COUNT(*) AS total_stays,
				COUNT(*) FILTER (WHERE NOT EXISTS (
					SELECT 1 FROM reservations r 
					WHERE r.stay_id = s.id 
					AND r.arrived <= NOW() 
					AND r.departure > NOW()
				)) AS stay_free,
				COUNT(*) FILTER (WHERE EXISTS (
					SELECT 1 FROM reservations r 
					WHERE r.stay_id = s.id 
					AND r.arrived <= NOW() 
					AND r.departure > NOW()
				)) AS stay_occupied
			FROM stays s
			WHERE s.user_id = $1;
			`

	row := r.Db.QueryRowContext(ctx, query, userId)

	var statistics models.Statistics

	err := row.Scan(&statistics.StayTotal, &statistics.StayFree, &statistics.StayOccupied)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &statistics, nil
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

	var amenitiesData []byte

	err = row.Scan(
		&stay.ID,
		&stay.UserID,
		&stay.LocationID,
		&stay.Name,
		&stay.Type,
		&stay.Guests,
		&stay.Rating,
		&amenitiesData,
		&stay.House,
		&stay.Entrance,
		&stay.CreatedAt,
		&stay.UpdatedAt,
		&stay.Address,
		&stay.RoomsCount,
		&stay.BedsCount,
		&stay.Price,
		&stay.Period,
		&stay.OwnersRules,
		&stay.CancellationPolicy,
		&stay.DescribeProperty,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
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

		var amenitiesData []byte

		err = rows.Scan(
			&stay.ID,
			&stay.UserID,
			&stay.LocationID,
			&stay.Name,
			&stay.Type,
			&stay.Guests,
			&stay.Rating,
			&amenitiesData,
			&stay.House,
			&stay.Entrance,
			&stay.CreatedAt,
			&stay.UpdatedAt,
			&stay.Address,
			&stay.RoomsCount,
			&stay.BedsCount,
			&stay.Price,
			&stay.Period,
			&stay.OwnersRules,
			&stay.CancellationPolicy,
			&stay.DescribeProperty,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

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

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE stays SET location_id=$1, name=$2, type=$3, guests=$4, amenities=$5, house=$6, entrance=$7, address=$8, rooms_count=$9, beds_count=$10, price=$11, period=$12, owners_rules=$13, cancellation_policy=$14, describe_property=$15, updated_at=$16 WHERE id=$17")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	amenitiesJSON, err := json.Marshal(stay.Amenities)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(
		ctx,
		stay.LocationID,
		stay.Name,
		string(stay.Type),
		stay.Guests,
		amenitiesJSON,
		stay.House,
		stay.Entrance,
		stay.Address,
		stay.RoomsCount,
		stay.BedsCount,
		stay.Price,
		stay.Period,
		stay.OwnersRules,
		stay.CancellationPolicy,
		stay.DescribeProperty,
		time.Now(),
		id,
	)

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
	var amenitiesData []byte

	for rows.Next() {
		var stay models.Stay

		err = rows.Scan(
			&stay.ID,
			&stay.UserID,
			&stay.LocationID,
			&stay.Name,
			&stay.Type,
			&stay.Guests,
			&stay.Rating,
			&amenitiesData,
			&stay.House,
			&stay.Entrance,
			&stay.CreatedAt,
			&stay.UpdatedAt,
			&stay.Address,
			&stay.RoomsCount,
			&stay.BedsCount,
			&stay.Price,
			&stay.Period,
			&stay.OwnersRules,
			&stay.CancellationPolicy,
			&stay.DescribeProperty,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

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
	var amenitiesData []byte

	for rows.Next() {
		var stay models.Stay

		err = rows.Scan(
			&stay.ID,
			&stay.UserID,
			&stay.LocationID,
			&stay.Name,
			&stay.Type,
			&stay.Guests,
			&stay.Rating,
			&amenitiesData,
			&stay.House,
			&stay.Entrance,
			&stay.CreatedAt,
			&stay.UpdatedAt,
			&stay.Address,
			&stay.RoomsCount,
			&stay.BedsCount,
			&stay.Price,
			&stay.Period,
			&stay.OwnersRules,
			&stay.CancellationPolicy,
			&stay.DescribeProperty,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		if err = json.Unmarshal(amenitiesData, &stay.Amenities); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		stays = append(stays, stay)
	}

	return &stays, nil
}

func (this *Repo) Filtration(ctx context.Context, search models.Filtration) ([]models.Stay, error) {
	const op = "repo.stays.Filtration"

	sort.Float64s(search.Rating)

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
		WHERE location_id = $1
	`

	args := []interface{}{
		search.LocationID,
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
		query += fmt.Sprintf(" AND rooms_count = ANY($%d)", count)
		roomsStr := make([]string, len(search.NumberOfBedrooms))
		for i, val := range search.NumberOfBedrooms {
			roomsStr[i] = fmt.Sprintf("%d", val)
		}
		args = append(args, pq.Array(roomsStr))
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
		query += " AND (" + strings.Join(conditions, " AND ") + ")"
	}

	switch search.SortBy {
	case filtrationSort.Old:
		query += " ORDER BY created_at ASC"
		break
	case filtrationSort.New:
		query += " ORDER BY created_at DESC"
		break
	case filtrationSort.HighlyRecommended:
		query += " ORDER BY rating DESC, updated_at DESC"
		break
	case filtrationSort.LowlyRecommended:
		query += " ORDER BY rating ASC, updated_at ASC"
		break
	default:
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

	var amenitiesData []byte

	var stays []models.Stay
	for rows.Next() {
		var stay models.Stay
		if err = rows.Scan(
			&stay.ID,
			&stay.UserID,
			&stay.LocationID,
			&stay.Name,
			&stay.Type,
			&stay.Guests,
			&stay.Rating,
			&amenitiesData,
			&stay.House,
			&stay.Entrance,
			&stay.CreatedAt,
			&stay.UpdatedAt,
			&stay.Address,
			&stay.RoomsCount,
			&stay.BedsCount,
			&stay.Price,
			&stay.Period,
			&stay.OwnersRules,
			&stay.CancellationPolicy,
			&stay.DescribeProperty,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

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
