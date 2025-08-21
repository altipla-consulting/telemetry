package config

import (
	"context"
)

type Settings struct {
	Collectors []Collector
}

type Collector interface {
	ReportError(ctx context.Context, err error)
	Flush()
}
