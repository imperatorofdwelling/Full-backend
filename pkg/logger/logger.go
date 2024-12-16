package logger

import (
	"errors"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogpretty"
	"log"
	"log/slog"
	"os"
)

var unknownEnv = errors.New("unknown environment (should be local or prod)")

func New() *slog.Logger {
	var logger *slog.Logger
	switch config.GlobalEnv {
	case config.StageEnv:
		logger = setupPrettySlog()
	case config.LocalEnv:
		logger = setupPrettySlog()
	case config.DevEnv:
		logger = setupPrettySlog()
	case config.ProdEnv:
		// TODO change writer(example: file)
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log.Fatal(unknownEnv)
	}
	return logger
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
