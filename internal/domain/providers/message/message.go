package message

import (
	"database/sql"
	"github.com/google/wire"
	msgHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/message"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	msgRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/message"
	msgSvc "github.com/imperatorofdwelling/Full-backend/internal/service/message"
	"log/slog"
	"sync"
)

var (
	hdl     *msgHdl.Handler
	hdlOnce sync.Once

	svc     *msgSvc.Service
	svcOnce sync.Once

	rep     *msgRepo.Repo
	repOnce sync.Once
)

var MessageProviderSet wire.ProviderSet = wire.NewSet(
	ProvideMessageHandler,
	ProvideMessageService,
	ProvideMessageRepo,

	wire.Bind(new(interfaces.MessageHandler), new(*msgHdl.Handler)),
	wire.Bind(new(interfaces.MessageService), new(*msgSvc.Service)),
	wire.Bind(new(interfaces.MessageRepository), new(*msgRepo.Repo)),
)

func ProvideMessageHandler(svc interfaces.MessageService, log *slog.Logger) *msgHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &msgHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideMessageService(repo interfaces.MessageRepository) *msgSvc.Service {
	svcOnce.Do(func() {
		svc = &msgSvc.Service{
			Repo: repo,
		}
	})

	return svc
}

func ProvideMessageRepo(db *sql.DB) *msgRepo.Repo {
	repOnce.Do(func() {
		rep = &msgRepo.Repo{
			Db: db,
		}
	})

	return rep
}
