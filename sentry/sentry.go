package sentry

import (
	"context"
	"os"

	"github.com/altipla-consulting/sentry"

	"github.com/altipla-consulting/telemetry"
	"github.com/altipla-consulting/telemetry/internal/config"
)

func Reporter() telemetry.Option {
	return func(settings *config.Settings) {
		if os.Getenv("SENTRY_DSN") != "" {
			settings.Collectors = append(settings.Collectors, &sentryCollector{
				client: sentry.NewClient(os.Getenv("SENTRY_DSN")),
			})
		}
	}
}

type sentryCollector struct {
	client *sentry.Client
}

func (c *sentryCollector) ReportError(ctx context.Context, err error) {
	c.client.Report(ctx, err)
}
