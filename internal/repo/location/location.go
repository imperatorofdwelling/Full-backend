package location

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/location"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) FindByNameMatch(ctx context.Context, match string) (*[]location.Location, error) {
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
		var loc location.Location

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
