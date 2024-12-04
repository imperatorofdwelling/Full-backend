package providers

import (
	"github.com/google/wire"
	fileHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/file"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	fileSvc "github.com/imperatorofdwelling/Full-backend/internal/service/file"
	"log/slog"
	"sync"
)

var (
	hdl     *fileHdl.Handler
	hdlOnce sync.Once

	svc     *fileSvc.Service
	svcOnce sync.Once
)

var FileProviderSet wire.ProviderSet = wire.NewSet(
	ProvideFileHandler,
	ProvideFileService,

	wire.Bind(new(interfaces.FileHandler), new(*fileHdl.Handler)),
	wire.Bind(new(interfaces.FileService), new(*fileSvc.Service)),
)

func ProvideFileHandler(svc interfaces.FileService, log *slog.Logger) *fileHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &fileHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideFileService() *fileSvc.Service {
	svcOnce.Do(func() {
		svc = &fileSvc.Service{}
	})

	return svc
}
