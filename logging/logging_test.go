package logging

import (
	"context"
	"testing"

	"github.com/altipla-consulting/errors"
)

func TestLogCollectorReportPanic(t *testing.T) {
	var collector logCollector
	defer func() {
		collector.ReportError(context.Background(), errors.Recover(recover()))
	}()
	panic("test failure")
}
