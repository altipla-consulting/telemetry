package logging

import "log/slog"

const (
	LevelCritical = slog.Level(12)
	LevelNotice   = slog.Level(2)
	LevelTrace    = slog.Level(-8)
)
