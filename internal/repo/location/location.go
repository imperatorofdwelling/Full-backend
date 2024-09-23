package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/location"
)

type LocationRepo struct {
	Db *sql.DB
}

func (r *LocationRepo) FindByNameMatch(ctx context.Context, match string) (*[]location.Location, error) {
	const op = "repo.location.FindByNameMatch"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM locations WHERE LOWER(city) LIKE LOWER($1)")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var locations []location.Location

	rows, err := stmt.QueryContext(ctx, fmt.Sprintf("%%%s%%", match))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var location location.Location

		if err := rows.Scan(
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

		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &locations, nil
}
