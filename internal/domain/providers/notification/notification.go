package notification

import (
	"database/sql"
	"github.com/google/wire"
	ntfHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/notification"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	ntfRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/notification"
	ntfSvc "github.com/imperatorofdwelling/Full-backend/internal/service/notification"
	"github.com/imperatorofdwelling/Full-backend/pkg/pushNotification"
	"log/slog"
	"sync"
)

var (
	svcPush     *pushNotification.GorushClient
	svcPushOnce sync.Once

	hdl     *ntfHdl.Handler
	hdlOnce sync.Once

	svc     *ntfSvc.Service
	svcOnce sync.Once

	rep     *ntfRepo.Repo
	repOnce sync.Once
)

var PushClientProviderSet = wire.NewSet(
	ProvidePushNotificationSender,

	wire.Bind(new(interfaces.PushNotificationSender), new(*pushNotification.GorushClient)),
)

func ProvidePushNotificationSender(baseUrl string) *pushNotification.GorushClient {
	svcPushOnce.Do(func() {
		svcPush = &pushNotification.GorushClient{
			BaseURL: baseUrl,
		}
	})
	return svcPush
}

var NotificationProviderSet wire.ProviderSet = wire.NewSet(
	ProvideNotificationHandler,
	ProvideNotificationService,
	ProvideNotificationRepo,

	wire.Bind(new(interfaces.NotificationHandler), new(*ntfHdl.Handler)),
	wire.Bind(new(interfaces.NotificationService), new(*ntfSvc.Service)),
	wire.Bind(new(interfaces.NotificationRepository), new(*ntfRepo.Repo)),
)

func ProvideNotificationHandler(svc interfaces.NotificationService, log *slog.Logger) *ntfHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &ntfHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideNotificationService(repo interfaces.NotificationRepository, push interfaces.PushNotificationSender) *ntfSvc.Service {
	svcOnce.Do(func() {
		svc = &ntfSvc.Service{
			Repo:       repo,
			PushClient: push,
		}
	})

	return svc
}

func ProvideNotificationRepo(db *sql.DB) *ntfRepo.Repo {
	repOnce.Do(func() {
		rep = &ntfRepo.Repo{
			Db: db,
		}
	})

	return rep
}
