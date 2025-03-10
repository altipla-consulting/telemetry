package logging

import (
	"context"
	"log/slog"
	"os"

	"github.com/altipla-consulting/env"
	"github.com/altipla-consulting/errors"
	"github.com/lmittmann/tint"

	"github.com/altipla-consulting/telemetry"
	"github.com/altipla-consulting/telemetry/internal/config"
)

func Standard() telemetry.Option {
	return configureLevels(slog.LevelInfo, slog.LevelDebug)
}

func Debug() telemetry.Option {
	return configureLevels(slog.LevelDebug, slog.LevelDebug)
}

func Trace() telemetry.Option {
	return configureLevels(LevelTrace, LevelTrace)
}

func configureLevels(level slog.Level, localLevel slog.Level) telemetry.Option {
	return func(settings *config.Settings) {
		settings.Collectors = append(settings.Collectors, new(logCollector))

		if env.IsLocal() {
			slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
				Level: localLevel,
			})))

			return
		}

		logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level: level,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				switch {
				case a.Key == slog.MessageKey:
					a.Key = "message"

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
	if env.IsLocal() {
		slog.Error(string(errors.Stack(err)))
	}
}
