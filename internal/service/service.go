package service

//import (
//	"context"
//	"github.com/imperatorofdwelling/Website-backend/internal/db"
//	accountI "github.com/imperatorofdwelling/Website-backend/internal/repo/account/interface"
//	interfaces "github.com/imperatorofdwelling/Website-backend/internal/service/interface"
//	"log/slog"
//	"strconv"
//)
//
//type service struct {
//	rAccount accountI.AccountRepository
//}
//
//func NewService(
//	accountRepository accountI.AccountRepository,
//) interfaces.ServiceUseCase {
//	return &service{
//		rAccount: accountRepository,
//	}
//}
//
//func (s *service) Migrate(ctx context.Context, log *slog.Logger) error {
//	if err := s.rAccount.Migrate(ctx, log); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// TODO check login
//
//func (s *service) checkIdParam(id string) (int64, error) {
//	idInt, err := strconv.ParseInt(id, 10, 64)
//	if err != nil || idInt <= 0 {
//		return 0, db.ErrParamNotFound
//	}
//	return idInt, nil
//}
