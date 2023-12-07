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

// ReportPanics report any panic that can be recovered if it happens. It should be called with defer before any code
// that should be protected.
func ReportPanics(ctx context.Context) {
	if r := recover(); r != nil {
		for _, reporter := range settings.ErrorReporters {
			reporter.ReportPanic(ctx, r)
		}
	}
}
