package service

//import (
//	"context"
//	"github.com/imperatorofdwelling/Website-backend/internal/db"
//	"github.com/imperatorofdwelling/Website-backend/internal/domain/account"
//	"log"
//	"strings"
//)
//
//func (s *service) Registration(ctx context.Context, newAccount account.Registration) (*account.Info, error) {
//	//check valid data
//	if !isValidAccountRegister(newAccount) {
//		log.Println("invalid data")
//		return nil, db.ErrValidate
//	}
//
//	result, err := s.rAccount.Registration(ctx, newAccount)
//	if err != nil {
//		return nil, err
//	}
//	return result, nil
//}
//
//func (s *service) Login(ctx context.Context, acc account.Login) (int64, error) {
//	id, err := s.rAccount.Login(ctx, acc)
//	if err != nil {
//		return 0, err
//	}
//	return id, nil
//}
//
//func (s *service) PutAccount(ctx context.Context, id string, updateAcc account.Info) (*account.Info, error) {
//	idInt, err := s.checkIdParam(id)
//	if err != nil {
//		return nil, err
//	}
//
//	// TODO добавить валдацию для account.Info
//	//if !isValidAccountRegister(updateAcc) {
//	//	return nil, db.ErrValidate
//	//}
//
//	result, err := s.rAccount.Put(ctx, idInt, updateAcc)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
//
//// TODO implement delete account
//
//func isValidAccountRegister(newAccountRegistration account.Registration) bool {
//	if newAccountRegistration.Name == "" || strings.TrimSpace(newAccountRegistration.Name) == "" {
//		return false
//	}
//
//	if newAccountRegistration.Email == "" || strings.TrimSpace(newAccountRegistration.Email) == "" {
//		return false
//	}
//
//	if newAccountRegistration.Password == "" || strings.TrimSpace(newAccountRegistration.Password) == "" {
//		return false
//	}
//
//	return true
//}
