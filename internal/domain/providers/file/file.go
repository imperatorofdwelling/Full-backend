package providers

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/service/file"
	"sync"
)

var (
	svc     *file.Service
	svcOnce sync.Once
)

var FileProviderSet wire.ProviderSet = wire.NewSet(
	ProvideFileService,

	wire.Bind(new(interfaces.FileService), new(*file.Service)),
)

func ProvideFileService() *file.Service {
	svcOnce.Do(func() {
		svc = &file.Service{}
	})

	return svc
}
