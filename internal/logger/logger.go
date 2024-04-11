package logger

import (
	"log/slog"
	"os"
)

func SetupLogger() *slog.Logger {
	var log slog.Logger
	logData, err := os.OpenFile(os.Getenv("LOG_PATH"), os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic("could not open log file" + err.Error())
	}
	log = *slog.New(slog.NewTextHandler(logData, &slog.HandlerOptions{AddSource: false, Level: slog.LevelInfo}))
	return &log
}
