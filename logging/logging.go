package logging

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/altipla-consulting/env"
	"github.com/altipla-consulting/errors"
	"github.com/lmittmann/tint"

	"github.com/altipla-consulting/telemetry"
	"github.com/altipla-consulting/telemetry/internal/config"
)

func Standard() telemetry.Option {
	return stdLevel(slog.LevelInfo)
}

func Debug() telemetry.Option {
	return stdLevel(slog.LevelDebug)
}

func stdLevel(level slog.Level) telemetry.Option {
	return func(settings *config.Settings) {
		settings.Collectors = append(settings.Collectors, new(logCollector))

		if env.IsLocal() {
			slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
				Level: slog.LevelDebug,
			})))

			return
		}

		logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			AddSource: true,
			Level:     level,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				switch {
				case a.Key == slog.MessageKey:
					a.Key = "message"

				case a.Key == slog.SourceKey:
					a.Key = "logging.googleapis.com/sourceLocation"

				case a.Key == slog.LevelKey:
					a.Key = "severity"
					switch level := a.Value.Any().(slog.Level); level {
					// Map a different name for warnings.
					case slog.LevelWarn:
						a.Value = slog.StringValue("WARNING")

					// Custom levels.
					case LevelTrace:
						a.Value = slog.StringValue("DEFAULT")
					case LevelNotice:
						a.Value = slog.StringValue("NOTICE")
					case LevelCritical:
						a.Value = slog.StringValue("CRITICAL")
					}
				}
				return a
			},
		}))
		slog.SetDefault(logger)
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
