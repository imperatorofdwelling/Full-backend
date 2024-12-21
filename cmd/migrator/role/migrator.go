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
	log := logger.New(logger.EnvLocal)

	sqlDB, err := db.ConnectToBD(cfg)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer sqlDB.Close()

	// Вставка объектов adm_objects
	if err := insertAdmObjects(sqlDB); err != nil {
		log.Error("Error inserting adm_objects: ", err)
		return
	}
	log.Info("Adm_objects inserted successfully.")

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

func insertAdmObjects(db *sql.DB) error {
	query := `
        INSERT INTO adm_object (route, action) VALUES
            ($1, $2),
            ($3, $4),
            ($5, $6),
            ($7, $8),
            ($9, $10),
            ($11, $12),
            ($13, $14);
    `

	_, err := db.Exec(query,
		"/advantages", pq.Array([]string{"GET"}),
		"/locations", pq.Array([]string{"GET"}),
		"/stays", pq.Array([]string{"GET"}),
		"/staysadvantage", pq.Array([]string{"GET"}),
		"/report", pq.Array([]string{"GET"}),
		"/staysreviews", pq.Array([]string{"GET"}),
		"/user", pq.Array([]string{"GET"}),
	)

	return err
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
		"guest",
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
	// tenant_id = 1, landlord_id = 2, guest_id = 3
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
		2, 1, // landlord -> /advantages
		2, 2, // landlord -> /registration
		2, 3, // landlord -> /login
		2, 4, // landlord -> /contract
		2, 5, // landlord -> /favourites
		2, 6, // landlord -> /locations
		2, 7, // landlord -> /reservation
		2, 8, // landlord -> /history
		2, 9, // landlord -> /stays
		2, 10, // landlord -> /staysadvantage
		2, 11, // landlord -> /report
		2, 12, // landlord -> /staysreviews
		2, 13, // landlord -> /user
		2, 14, // landlord -> /user/report
		3, 1, // guest -> /advantages
		3, 2, // guest -> /registration
		3, 3, // guest -> /login
		3, 4, // guest -> /contract
		3, 5, // guest -> /favourites
		3, 6, // guest -> /locations
		3, 7, // guest -> /reservation
		3, 8, // guest -> /history
		3, 9, // guest -> /stays
		3, 10, // guest -> /staysadvantage
		3, 11, // guest -> /report
		3, 12, // guest -> /staysreviews
		3, 13, // guest -> /user
		3, 14, // guest -> /user/report
	)

	return err
}
