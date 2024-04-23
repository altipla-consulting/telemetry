package logging

import (
	"context"
	"testing"
)

func TestLogCollectorReportPanic(t *testing.T) {
	var collector logCollector
	defer func() {
		collector.ReportPanic(context.Background(), recover())
	}()
	panic("test failure")
}
