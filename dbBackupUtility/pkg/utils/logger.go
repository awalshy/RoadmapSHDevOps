package utils

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func GetLogger() *slog.Logger {
	if logger == nil {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	return logger
}
