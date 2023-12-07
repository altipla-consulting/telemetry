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
)

func Standard() telemetry.Option {
	return func(settings *config.Settings) {
		settings.ErrorReporters = append(settings.ErrorReporters, &logReporter{})

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

type logReporter struct{}

// Report implements config.ErrorReporter.
func (*logReporter) Report(ctx context.Context, err error) {
	// empty
}

// ReportPanic implements config.ErrorReporter.
func (*logReporter) ReportPanic(ctx context.Context, panicErr any) {
	rec := errors.Recover(panicErr)
	slog.Error("Panic recovered", "error", rec.Error())
}

// ReportRequest implements config.ErrorReporter.
func (*logReporter) ReportRequest(r *http.Request, err error) {
	// empty
}
