package location

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/location"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) FindByNameMatch(ctx context.Context, match string) (*[]models.Location, error) {
	const op = "repo.location.FindByNameMatch"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM locations WHERE LOWER(city) LIKE LOWER($1)")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var locations []models.Location

	rows, err := stmt.QueryContext(ctx, fmt.Sprintf("%%%s%%", match))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var loc models.Location

		if err := rows.Scan(
			&loc.ID,
			&loc.City,
			&loc.FederalDistrict,
			&loc.FiasID,
			&loc.KladrID,
			&loc.Lat,
			&loc.Lon,
			&loc.Okato,
			&loc.Oktmo,
			&loc.Population,
			&loc.RegionIsoCode,
			&loc.RegionName,
			&loc.CreatedAt,
			&loc.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		locations = append(locations, loc)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &locations, nil
}

func (r *Repo) GetByID(ctx context.Context, id uuid.UUID) (*models.Location, error) {
	const op = "repo.location.GetByID"

	stmt, err := r.Db.PrepareContext(ctx, `
        SELECT id, city, federal_district, fias_id, kladr_id, lat, lon, okato, oktmo, 
               population, region_iso_code, region_name, created_at, updated_at 
        FROM locations WHERE id = $1
    `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var location models.Location

	row := stmt.QueryRowContext(ctx, id)
	if row.Err() != nil {
		return nil, fmt.Errorf("%s: %w", op, row.Err())
	}

	if err = row.Scan(
		&location.ID,
		&location.City,
		&location.FederalDistrict,
		&location.FiasID,
		&location.KladrID,
		&location.Lat,
		&location.Lon,
		&location.Okato,
		&location.Oktmo,
		&location.Population,
		&location.RegionIsoCode,
		&location.RegionName,
		&location.CreatedAt,
		&location.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &location, nil
}

func (r *Repo) GetAll(ctx context.Context) (*[]models.Location, error) {
	const op = "repo.location.GetAll"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM locations")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var locations []models.Location

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var loc models.Location

		err := rows.Scan(
			&loc.ID,
			&loc.City,
			&loc.FederalDistrict,
			&loc.FiasID,
			&loc.KladrID,
			&loc.Lat,
			&loc.Lon,
			&loc.Okato,
			&loc.Oktmo,
			&loc.Population,
			&loc.RegionIsoCode,
			&loc.RegionName,
			&loc.CreatedAt,
			&loc.UpdatedAt)
		if err != nil {
			return nil, err
		}

		locations = append(locations, loc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &locations, nil
}

func (r *Repo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	const op = "repo.location.DeleteByID"

	stmt, err := r.Db.PrepareContext(ctx, "DELETE FROM locations WHERE id = $1")
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

func (r *Repo) UpdateByID(ctx context.Context, id uuid.UUID, location models.LocationEntity) error {
	const op = "repo.location.UpdateByID"

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE locations SET city=$2, federal_district=$3, fias_id=$4, kladr_id=$5, lat=$6, lon=$7, okato=$8, oktmo=$9, population=$10, region_iso_code=$11, region_name=$12 WHERE id=$1")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id, location.City, location.FederalDistrict, location.FiasID, location.KladrID, location.Lat, location.Lon, location.Okato, location.Oktmo, location.Population, location.RegionIsoCode, location.RegionName)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
