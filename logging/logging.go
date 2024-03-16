package logging

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/altipla-consulting/env"
	"github.com/altipla-consulting/errors"
	"github.com/sirupsen/logrus"

	"github.com/altipla-consulting/telemetry"
	"github.com/altipla-consulting/telemetry/internal/config"
	"github.com/lmittmann/tint"
)

func Standard() telemetry.Option {
	return func(settings *config.Settings) {
		settings.Collectors = append(settings.Collectors, new(logCollector))

		if env.IsLocal() {
			slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
				Level: slog.LevelDebug,
			})))

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

type logCollector struct{}

func (*logCollector) ReportError(ctx context.Context, err error) {
	// empty
}

func (*logCollector) ReportErrorRequest(r *http.Request, err error) {
	// empty
}

func (*logCollector) ReportPanic(ctx context.Context, panicErr any) {
	rec := errors.Recover(panicErr)
	slog.Error("Panic recovered", "error", rec.Error())
}
