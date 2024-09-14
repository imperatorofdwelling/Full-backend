package account

//import (
//	"context"
//	"database/sql"
//	"errors"
//	"github.com/imperatorofdwelling/Website-backend/internal/db"
//	"github.com/imperatorofdwelling/Website-backend/internal/domain/account"
//	interfaces "github.com/imperatorofdwelling/Website-backend/internal/repo/account/interface"
//	"github.com/jackc/pgconn"
//	"log/slog"
//)
//
//type accountDataBase struct {
//	db *sql.DB
//}
//
//func NewAccountDataBase(db *sql.DB) interfaces.AccountRepository {
//	return &accountDataBase{
//		db: db,
//	}
//}
//
//func (r *accountDataBase) Migrate(ctx context.Context, log *slog.Logger) error {
//	accQuery := `
//		CREATE TABLE IF NOT EXISTS account(
//		id SERIAL PRIMARY KEY,
//		name text NOT NULL,
//		email text NOT NULL,
//		password text NOT NULL,
//		phone text,
//		birth_date date,
//		national text,
//		gender text
//	);
//    `
//	_, err := r.db.ExecContext(ctx, accQuery)
//	if err != nil {
//		message := db.ErrMigrate.Error() + " account"
//		log.Error("%q: %s\n", message, err.Error())
//		return db.ErrMigrate
//	}
//	log.Info("account table created")
//	return err
//}
//
//func (r *accountDataBase) Registration(ctx context.Context, newAccount account.Registration) (*account.Info, error) {
//	// Check if an account with the same email already exists
//	var existingCount int
//	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM account WHERE email = $1", newAccount.Email).Scan(&existingCount)
//	if err != nil {
//		return nil, err
//	}
//
//	if existingCount > 0 {
//		return nil, db.ErrDuplicate
//	}
//
//	var id int64
//
//	err = r.db.QueryRowContext(ctx,
//		"INSERT INTO account(name, email, password, phone, birth_date, national, gender) values($1, $2, $3, $4, $5, $6, $7) RETURNING id",
//		newAccount.Name, newAccount.Email, newAccount.Password, "", newAccount.Name, "", "").Scan(&id)
//	// Check if a user with the same email already exists
//	if err != nil {
//		var pgxError *pgconn.PgError
//		if errors.As(err, &pgxError) {
//			if pgxError.Code == "23505" {
//				return nil, db.ErrDuplicate
//			}
//		}
//		return nil, err
//	}
//
//	// Add the new account
//	requestAccount := &account.Info{
//		ID:        id,
//		Name:      newAccount.Name,
//		Email:     newAccount.Email,
//		Phone:     "",
//		BirthDate: "",
//		National:  "",
//		Gender:    "",
//	}
//
//	return requestAccount, nil
//}
//
//func (r *accountDataBase) Login(ctx context.Context, acc account.Login) (int64, error) {
//	var id int64
//	row := r.db.QueryRowContext(ctx, "SELECT id FROM account WHERE email = $1 and password = $2", acc.Email, acc.Password)
//
//	if err := row.Scan(&id); err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return 0, db.ErrNotExist
//		}
//		return 0, err
//	}
//	return id, nil
//}
//
//// TODO изменить отдельно на изменение пароля
//func (r *accountDataBase) Put(ctx context.Context, id int64, updateAcc account.Info) (*account.Info, error) {
//	res, err := r.db.ExecContext(ctx, "UPDATE account SET name = $1, email = $2, phone = $3, birth_date = $4, national = $5, gender = $6 WHERE account.id = $7",
//		updateAcc.Name, updateAcc.Email, updateAcc.Phone, updateAcc.BirthDate, updateAcc.National, updateAcc.Gender, id)
//	if err != nil {
//		var pgxError *pgconn.PgError
//		if errors.As(err, &pgxError) {
//			if pgxError.Code == "23505" {
//				return nil, db.ErrDuplicate
//			}
//		}
//		return nil, err
//	}
//
//	result := &account.Info{
//		ID:        id,
//		Name:      updateAcc.Name,
//		Email:     updateAcc.Email,
//		Phone:     updateAcc.Phone,
//		BirthDate: updateAcc.BirthDate,
//		National:  updateAcc.National,
//		Gender:    updateAcc.Gender,
//	}
//
//	rowsAffected, err := res.RowsAffected()
//	if err != nil {
//		return nil, err
//	}
//
//	if rowsAffected == 0 {
//		return nil, db.ErrUpdateFailed
//	}
//
//	return result, nil
//}
//func (r *accountDataBase) Delete(ctx context.Context, id int64) error {
//	res, err := r.db.ExecContext(ctx, "DELETE FROM account WHERE id = $1", id)
//	if err != nil {
//		return err
//	}
//
//	rowsAffected, err := res.RowsAffected()
//	if err != nil {
//		return err
//	}
//
//	if rowsAffected == 0 {
//		return db.ErrNotExist
//	}
//
//	return nil
//}
