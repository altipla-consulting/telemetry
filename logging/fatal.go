package logging

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/altipla-consulting/telemetry"
	"github.com/getsentry/sentry-go"
)

func Fatal(msg string, err error) {
	slog.Log(context.Background(), LevelCritical, msg, slog.String("err", err.Error()))
	telemetry.ReportError(context.Background(), err)
	if os.Getenv("SENTRY_DSN") != "" {
		sentry.Flush(5 * time.Second)
	}
	os.Exit(1)
}
