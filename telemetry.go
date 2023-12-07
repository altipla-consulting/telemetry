package telemetry

import (
	"context"
	"net/http"

	"github.com/altipla-consulting/telemetry/internal/config"
)

type Option func(settings *config.Settings)

var settings = new(config.Settings)

func Configure(opts ...Option) {
	for _, opt := range opts {
		opt(settings)
	}
}

func ReportError(ctx context.Context, err error) {
	for _, reporter := range settings.ErrorReporters {
		reporter.Report(ctx, err)
	}
}

func ReportErrorRequest(r *http.Request, err error) {
	for _, reporter := range settings.ErrorReporters {
		reporter.ReportRequest(r, err)
	}
}
