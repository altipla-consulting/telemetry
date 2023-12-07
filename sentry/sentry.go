package sentry

import (
	"os"

	"github.com/altipla-consulting/sentry"

	"github.com/altipla-consulting/telemetry"
	"github.com/altipla-consulting/telemetry/internal/config"
)

func Reporter() telemetry.Option {
	return func(settings *config.Settings) {
		if os.Getenv("SENTRY_DSN") != "" {
			settings.ErrorReporters = append(settings.ErrorReporters, sentry.NewClient(os.Getenv("SENTRY_DSN")))
		}
	}
}
