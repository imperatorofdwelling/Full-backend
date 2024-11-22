package chat

import (
	"database/sql"
	"github.com/google/wire"
	chatHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/chat"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	chatRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/chat"
	chatSvc "github.com/imperatorofdwelling/Full-backend/internal/service/chat"
	"log/slog"
	"sync"
)

var (
	hdl     *chatHdl.Handler
	hdlOnce sync.Once

	svc     *chatSvc.Service
	svcOnce sync.Once

	rep     *chatRepo.Repo
	repOnce sync.Once
)

var ChatProviderSet wire.ProviderSet = wire.NewSet(
	ProvideChatHandler,
	ProvideChatService,
	ProvideChatRepo,

	wire.Bind(new(interfaces.ChatHandler), new(*chatHdl.Handler)),
	wire.Bind(new(interfaces.ChatService), new(*chatSvc.Service)),
	wire.Bind(new(interfaces.ChatRepository), new(*chatRepo.Repo)),
)

func ProvideChatHandler(svc interfaces.ChatService, log *slog.Logger) *chatHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &chatHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideChatService(repo interfaces.ChatRepository) *chatSvc.Service {
	svcOnce.Do(func() {
		svc = &chatSvc.Service{
			Repo: repo,
		}
	})

	return svc
}

func ProvideChatRepo(db *sql.DB) *chatRepo.Repo {
	repOnce.Do(func() {
		rep = &chatRepo.Repo{
			Db: db,
		}
	})

	return rep
}
