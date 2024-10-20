// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/imperatorofdwelling/Full-backend/internal/api"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/db"
	providers2 "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/advantage"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/providers/auth"
	user2 "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/favourite"
	providers3 "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/file"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/providers/location"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/providers/reservation"
	providers4 "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/stays"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/providers/staysadvantage"
	providers5 "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/staysreviews"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/providers/user"
	"log/slog"
)

// Injectors from wire.go:

func InitializeAPI(cfg *config.Config, log *slog.Logger) (*api.ServerHTTP, error) {
	sqlDB, err := db.ConnectToBD(cfg)
	if err != nil {
		return nil, err
	}
	repository := auth.ProvideAuthRepository(sqlDB)
	userRepository := user.ProvideUserRepository(sqlDB)
	service := auth.ProvideAuthService(repository, userRepository)
	authHandler := auth.ProvideAuthHandler(service, log)
	userService := user.ProvideUserService(userRepository)
	userHandler := user.ProvideUserHandler(userService, log)
	repo := providers.ProvideLocationRepository(sqlDB)
	locationService := providers.ProvideLocationService(repo)
	handler := providers.ProvideLocationHandler(locationService, log)
	advantageRepo := providers2.ProvideAdvantageRepository(sqlDB)
	fileService := providers3.ProvideFileService()
	advantageService := providers2.ProvideAdvantageService(advantageRepo, fileService)
	advantageHandler := providers2.ProvideAdvantageHandler(advantageService, log)
	staysRepo := providers4.ProvideStaysRepo(sqlDB)
	staysService := providers4.ProvideStaysService(staysRepo, locationService)
	staysHandler := providers4.ProvideStaysHandler(staysService, log)
	staysadvantageRepo := staysadvantage.ProvideStaysAdvantageRepo(sqlDB)
	staysadvantageService := staysadvantage.ProvideStaysAdvantageService(staysadvantageRepo, staysService, advantageService)
	staysadvantageHandler := staysadvantage.ProvideStaysAdvantageHandler(staysadvantageService, log)
	reservationRepo := reservation.ProvideReservationRepository(sqlDB)
	reservationService := reservation.ProvideReservationService(reservationRepo)
	reservationHandler := reservation.ProvideReservationHandler(reservationService, log)
	staysreviewsRepo := providers5.ProvideStaysReviewsRepository(sqlDB)
	staysreviewsService := providers5.ProvideStaysReviewsService(staysreviewsRepo)
	staysreviewsHandler := providers5.ProvideStaysReviewsHandler(staysreviewsService, log)
	favouriteRepo := user2.ProvideFavouriteRepository(sqlDB)
	favouriteService := user2.ProvideFavouriteService(favouriteRepo)
	favHandler := user2.ProvideFavouriteHandler(favouriteService, log)
	serverHTTP := api.NewServerHTTP(cfg, authHandler, userHandler, handler, advantageHandler, staysHandler, staysadvantageHandler, reservationHandler, staysreviewsHandler, favHandler)
	return serverHTTP, nil
}
