package telemetry

import (
	"context"
	"net/http"

	"github.com/altipla-consulting/telemetry/internal/config"
)

type Option func(settings *config.Settings)

var (
	settings    = new(config.Settings)
	initialized = false
)

func Configure(opts ...Option) {
	if !initialized {
		for _, opt := range opts {
			opt(settings)
		}
		initialized = true
	} else {
		panic("Configure must be called only once")
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
	if r := recover(); r != nil { // revive:disable-line:defer
		for _, reporter := range settings.ErrorReporters {
			reporter.ReportPanic(ctx, r)
		}
	}
}
