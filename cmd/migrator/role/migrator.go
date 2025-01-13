package main

import (
	"database/sql"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/db"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/lib/pq"
)

const maxRoles = 3

func main() {
	cfg := config.LoadConfig()
	log := logger.New()

	sqlDB, err := db.ConnectToBD(cfg)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer sqlDB.Close()

	// Вставка объектов adm_objects для роли Tenant
	if err := insertAdmObjectsForTenant(sqlDB); err != nil {
		log.Error("Error inserting adm_objects for tenant: ", err)
		return
	}
	log.Info("Adm_objects for tenant inserted successfully.")

	// Вставка объектов adm_objects для роли Landlord
	if err := insertAdmObjectsForLandlord(sqlDB); err != nil {
		log.Error("Error inserting adm_objects for landlord: ", err)
		return
	}
	log.Info("Adm_objects for landlord inserted successfully.")

	if err := insertRoles(sqlDB); err != nil {
		log.Error("Error inserting roles: %v", err)
		return
	}
	log.Info("Roles inserted successfully.")

	if err := linkRolesToObjects(sqlDB); err != nil {
		log.Error("Error linking roles to objects: %v", err)
		return
	}
	log.Info("Roles linked successfully.")
	return
}

func insertAdmObjectsForTenant(db *sql.DB) error {
	query := `
        INSERT INTO adm_object (route, action) VALUES
            ($1, $2),
            ($3, $4),
            ($5, $6),
            ($7, $8),
            ($9, $10),
            ($11, $12),
            ($13, $14),
            ($15, $16),
            ($17, $18),
            ($19, $20),
            ($21, $22),
            ($23, $24),
            ($25, $26);
    `

	_, err := db.Exec(query,
		"/advantages", pq.Array([]string{"GET", "POST"}),
		"/registration", pq.Array([]string{"POST"}),
		"/login", pq.Array([]string{"POST"}),
		"/contract", pq.Array([]string{"GET", "POST"}),
		"/favourites", pq.Array([]string{"GET", "POST", "DELETE"}),
		"/locations", pq.Array([]string{"GET"}),
		"/reservation", pq.Array([]string{"GET"}),
		"/history", pq.Array([]string{"GET", "POST"}),
		"/stays", pq.Array([]string{"GET"}),
		"/staysadvantage", pq.Array([]string{"GET"}),
		"/report", pq.Array([]string{"GET", "POST", "PUT", "DELETE"}),
		"/staysreviews", pq.Array([]string{"GET", "POST", "PUT", "DELETE"}),
		"/user", pq.Array([]string{"GET", "PUT", "DELETE"}),
		"/user/report", pq.Array([]string{"GET", "POST", "PUT", "DELETE"}),
	)

	return err
}

func insertAdmObjectsForLandlord(db *sql.DB) error {
	query := `
        INSERT INTO adm_object (route, action) VALUES
            ($1, $2),
            ($3, $4),
            ($5, $6),
            ($7, $8),
            ($9, $10),
            ($11, $12),
            ($13, $14),
            ($15, $16),
            ($17, $18),
            ($19, $20),
            ($21, $22),
            ($23, $24),
            ($25, $26);
    `

	_, err := db.Exec(query,
		"/advantages", pq.Array([]string{"GET", "POST", "PUT", "DELETE"}),
		"/registration", pq.Array([]string{"POST"}),
		"/login", pq.Array([]string{"POST"}),
		"/contract", pq.Array([]string{"GET", "PUT", "POST"}),
		"/favourites", pq.Array([]string{"GET", "POST", "PUT", "DELETE"}),
		"/locations", pq.Array([]string{"GET", "DELETE"}),
		"/reservation", pq.Array([]string{"GET", "POST", "PUT", "DELETE"}),
		"/history", pq.Array([]string{"GET", "POST"}),
		"/stays", pq.Array([]string{"GET", "POST", "PUT", "DELETE"}),
		"/staysadvantage", pq.Array([]string{"GET", "POST", "PUT", "DELETE"}),
		"/report", pq.Array([]string{"GET", "POST", "PUT", "DELETE"}),
		"/staysreviews", pq.Array([]string{"GET", "POST", "PUT", "DELETE"}),
		"/user", pq.Array([]string{"GET", "PUT", "DELETE"}),
		"/user/report", pq.Array([]string{"GET", "POST", "PUT", "DELETE"}),
	)

	return err
}
func insertRoles(db *sql.DB) error {
	query := `
        INSERT INTO role (name) VALUES
            ($1),
            ($2),
            ($3);
    `

	_, err := db.Exec(query,
		"tenant",
		"landlord",
	)

	return err
}

func linkRolesToObjects(db *sql.DB) error {
	query := `
        INSERT INTO role_object (role_id, object_id) VALUES
            ($1, $2),
            ($3, $4),
            ($5, $6),
            ($7, $8),
            ($9, $10),
            ($11, $12),
            ($13, $14),
            ($15, $16),
            ($17, $18),
            ($19, $20),
            ($21, $22),
            ($23, $24),
            ($25, $26),
            ($27, $28),
            ($29, $30),
            ($31, $32);
    `

	// Предположим, что id ролей и объектов известны
	// tenant_id = 1, landlord_id = 2
	// object_ids от 1 до 14, как в предыдущих вставках
	_, err := db.Exec(query,
		1, 1, // tenant -> /advantages
		1, 2, // tenant -> /registration
		1, 3, // tenant -> /login
		1, 4, // tenant -> /contract
		1, 5, // tenant -> /favourites
		1, 6, // tenant -> /locations
		1, 7, // tenant -> /reservation
		1, 8, // tenant -> /history
		1, 9, // tenant -> /stays
		1, 10, // tenant -> /staysadvantage
		1, 11, // tenant -> /report
		1, 12, // tenant -> /staysreviews
		1, 13, // tenant -> /user
		1, 14, // tenant -> /user/report
		2, 15, // landlord -> /advantages
		2, 16, // landlord -> /registration
		2, 17, // landlord -> /login
		2, 18, // landlord -> /contract
		2, 19, // landlord -> /favourites
		2, 20, // landlord -> /locations
		2, 21, // landlord -> /reservation
		2, 22, // landlord -> /history
		2, 23, // landlord -> /stays
		2, 24, // landlord -> /staysadvantage
		2, 25, // landlord -> /report
		2, 26, // landlord -> /staysreviews
		2, 27, // landlord -> /user
		2, 28, // landlord -> /user/report
	)

	return err
}
