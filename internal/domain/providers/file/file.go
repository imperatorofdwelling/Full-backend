package providers

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/interfaces"
	service "github.com/imperatorofdwelling/Website-backend/internal/service/file"
	"sync"
)

var (
	svc     *service.FileService
	svcOnce sync.Once
)

var FileProviderSet wire.ProviderSet = wire.NewSet(
	ProvideFileService,

	wire.Bind(new(interfaces.FileService), new(*service.FileService)),
)

func ProvideFileService() *service.FileService {
	svcOnce.Do(func() {
		svc = &service.FileService{}
	})

	return svc
}
