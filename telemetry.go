package telemetry

import (
	"context"
	"net/http"

	"github.com/altipla-consulting/errors"

	"github.com/altipla-consulting/telemetry/internal/config"
)

type Option func(settings *config.Settings)

var (
	settings   = new(config.Settings)
	configured = false
)

func Configure(opts ...Option) {
	if !configured {
		for _, opt := range opts {
			opt(settings)
		}
		configured = true
	} else {
		panic("telemetry.Configure() must be called only once at the start of the program")
	}
}

func ReportError(ctx context.Context, err error) {
	if !configured {
		panic("telemetry.Configure() must be called before reporting any error")
	}

	for _, collector := range settings.Collectors {
		collector.ReportError(ctx, err)
	}
}

// Deprecated: Use ReportError() instead, the request is detected automatically.
func ReportErrorRequest(r *http.Request, err error) {
	ReportError(r.Context(), err)
}

// ReportPanics report any panic that can be recovered if it happens. It should be called with defer before any code
// that should be protected.
func ReportPanics(ctx context.Context) {
	if err := errors.Recover(recover()); err != nil { // revive:disable-line:defer
		ReportError(ctx, err)
	}
}

// Deprecated: Use ReportError() instead.
func ReportPanic(ctx context.Context, err error) {
	ReportError(ctx, err)
}

// DefaultReporter merging all the methods in a single struct that can be passed to external interfaces.
var DefaultReporter = defaultReporter{}

type defaultReporter struct{}

func (defaultReporter) ReportError(ctx context.Context, err error) {
	ReportError(ctx, err)
}

// Deprecated: Use ReportError() instead.
func (defaultReporter) ReportErrorRequest(r *http.Request, err error) {
	ReportErrorRequest(r, err)
}

// Deprecated: Use ReportError() instead.
func (defaultReporter) ReportPanics(ctx context.Context) {
	ReportPanics(ctx)
}
