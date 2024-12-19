package confirmEmail

import (
	"database/sql"
	"github.com/google/wire"
	emailConfHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/confirmEmail"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	emailConfRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/confirmEmail"
	emailConfService "github.com/imperatorofdwelling/Full-backend/internal/service/confirmEmail"
	"log/slog"
	"sync"
)

var (
	hdl     *emailConfHdl.Handler
	hdlOnce sync.Once

	svc     *emailConfService.Service
	svcOnce sync.Once

	repo     *emailConfRepo.Repo
	repoOnce sync.Once
)

var ProvideSet = wire.NewSet(
	ProvideConfirmEmailHandler,
	ProvideConfirmEmailService,
	ProvideConfirmEmailRepo,

	wire.Bind(new(interfaces.ConfirmEmailHandler), new(*emailConfHdl.Handler)),
	wire.Bind(new(interfaces.ConfirmEmailService), new(*emailConfService.Service)),
	wire.Bind(new(interfaces.ConfirmEmailRepository), new(*emailConfRepo.Repo)),
)

func ProvideConfirmEmailHandler(svc interfaces.ConfirmEmailService, log *slog.Logger) *emailConfHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &emailConfHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideConfirmEmailService(confirmEmailRepo interfaces.ConfirmEmailRepository, userRepo interfaces.UserRepository) *emailConfService.Service {
	svcOnce.Do(func() {
		svc = &emailConfService.Service{
			ConfirmEmailRepo: confirmEmailRepo,
			UserRepo:         userRepo,
		}
	})

	return svc
}

func ProvideConfirmEmailRepo(db *sql.DB) *emailConfRepo.Repo {
	repoOnce.Do(func() {
		repo = &emailConfRepo.Repo{
			DB: db,
		}
	})

	return repo
}
