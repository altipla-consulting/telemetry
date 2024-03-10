package telemetry

import (
	"context"
	"net/http"

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

func ReportErrorRequest(r *http.Request, err error) {
	if !configured {
		panic("telemetry.Configure() must be called before reporting any error")
	}

	for _, collector := range settings.Collectors {
		collector.ReportErrorRequest(r, err)
	}
}

// ReportPanics report any panic that can be recovered if it happens. It should be called with defer before any code
// that should be protected.
func ReportPanics(ctx context.Context) {
	if !configured {
		panic("telemetry.Configure() must be called before reporting any panic")
	}

	if r := recover(); r != nil { // revive:disable-line:defer
		for _, reporter := range settings.Collectors {
			reporter.ReportPanic(ctx, r)
		}
	}
}

// DefaultReporter merging all the methods in a single struct that can be passed to external interfaces.
var DefaultReporter = defaultReporter{}

type defaultReporter struct{}

func (defaultReporter) ReportError(ctx context.Context, err error) {
	ReportError(ctx, err)
}

func (defaultReporter) ReportErrorRequest(r *http.Request, err error) {
	ReportErrorRequest(r, err)
}

func (defaultReporter) ReportPanic(ctx context.Context, panicErr any) {
	ReportPanics(ctx)
}
