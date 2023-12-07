package config

import (
	"context"
	"net/http"
)

type Settings struct {
	ErrorReporters []ErrorReporter
}

type ErrorReporter interface {
	Report(ctx context.Context, err error)
	ReportRequest(r *http.Request, err error)
	ReportPanic(ctx context.Context, panicErr any)
}
