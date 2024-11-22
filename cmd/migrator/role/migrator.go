package main

import (
	"database/sql"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/db"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
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

	err = checkRolesCount(sqlDB)
	if err == nil {
		log.Info("Role count is correct")
		return
	} else {
		log.Warn(err.Error())
		log.Warn("Resetting all tables and reinitializing roles and permissions")

		if err = resetTables(sqlDB); err != nil {
			log.Error(err.Error())
			return
		}
	}

	if err := addRoles(sqlDB); err != nil {
		log.Error(err.Error())
		return
	}

	if err := addPermissions(sqlDB); err != nil {
		log.Error(err.Error())
		return
	}

	if err := assignPermissions(sqlDB); err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Roles and permissions successfully initialized!")

}

func checkRolesCount(db *sql.DB) error {
	stmt, err := db.Prepare("SELECT COUNT(*) FROM role")
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow().Scan(&count)
	if err != nil {
		return fmt.Errorf("error counting roles: %v", err)
	}

	if count != maxRoles {
		return fmt.Errorf("invalid roles count: expected %d, got %d", maxRoles, count)
	}

	return nil
}

func addRoles(db *sql.DB) error {
	roles := []string{"tenant", "landlord", "operator"}
	query := "INSERT INTO role (description) VALUES ($1)"
	for _, role := range roles {
		_, err := db.Exec(query, role)
		if err != nil {
			return fmt.Errorf("error adding role '%s': %v", role, err)
		}
	}
	return nil
}

func addPermissions(db *sql.DB) error {
	permissions := []string{
		"browse",
		"rent",
		"chat",
		"list own properties",
		"view own metrics",
	}
	query := "INSERT INTO permission (description) VALUES ($1)"
	for _, perm := range permissions {
		_, err := db.Exec(query, perm)
		if err != nil {
			return fmt.Errorf("error adding permission '%s': %v", perm, err)
		}
	}
	return nil
}

func assignPermissions(db *sql.DB) error {
	rolePermissions := map[string][]string{
		"tenant":   {"browse", "rent", "chat"},
		"landlord": {"browse", "chat", "list own properties", "view own metrics"},
		"operator": {"browse", "chat"},
	}

	for role, perms := range rolePermissions {
		for _, perm := range perms {
			query := `
				INSERT INTO role_has_permission (role_id, perm_id)
				SELECT r.id, p.id
				FROM role r, permission p
				WHERE r.description = $1 AND p.description = $2
			`
			_, err := db.Exec(query, role, perm)
			if err != nil {
				return fmt.Errorf("error assigning permission '%s' to role '%s': %v", perm, role, err)
			}
		}
	}
	return nil
}
func resetTables(db *sql.DB) error {
	tables := []string{"role_has_permission", "permission", "role"}

	// Используем TRUNCATE для очистки таблиц
	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table)
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("error truncating table '%s': %v", table, err)
		}
	}

	return nil
}
