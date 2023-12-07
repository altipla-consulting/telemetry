package logging

import (
	"log/slog"
	"os"

	"github.com/altipla-consulting/env"
	"github.com/sirupsen/logrus"

	"github.com/altipla-consulting/telemetry"
	"github.com/altipla-consulting/telemetry/internal/config"
)

func Standard() telemetry.Option {
	return func(settings *config.Settings) {
		if env.IsLocal() {
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}))
			slog.SetDefault(logger)

			logrus.SetFormatter(&logrus.TextFormatter{
				ForceColors: true,
			})
			logrus.SetLevel(logrus.DebugLevel)

			return
		}

		logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
		slog.SetDefault(logger)

		logrus.SetFormatter(new(logrus.JSONFormatter))
	}
}
