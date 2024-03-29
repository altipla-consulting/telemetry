package logging

import (
	"context"
	"log/slog"
	"os"

	"github.com/altipla-consulting/telemetry"
)

func Fatal(msg string, err error) {
	slog.Log(context.Background(), LevelCritical, msg, slog.String("err", err.Error()))
	telemetry.ReportError(context.Background(), err)
	os.Exit(1)
}
