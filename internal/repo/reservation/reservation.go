package reservation

import "database/sql"

type Repo struct {
	Db *sql.DB
}
