package logging

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/altipla-consulting/telemetry"
)

// Fatal logs a critical error and exits the program.
func Fatal(msg string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Log(ctx, LevelCritical, msg, slog.String("err", err.Error()))
	telemetry.ReportError(ctx, err)
	telemetry.Flush()

	os.Exit(1)
}
