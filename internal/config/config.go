package config

import (
	"context"
	"net/http"
)

type Settings struct {
	Collectors []Collector
}

type Collector interface {
	ReportError(ctx context.Context, err error)
	ReportErrorRequest(r *http.Request, err error)
	ReportPanic(ctx context.Context, err error)
}
