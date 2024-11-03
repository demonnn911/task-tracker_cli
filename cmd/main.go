package main

import (
	"log/slog"
	"os"
	"task-tracker_cli/pkg/handlers"

	"github.com/joho/godotenv"
)

func main() {

	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logHandler)
	arguments := os.Args
	if len(arguments) < 2 {
		logger.Error("incorrect input from user")
	}
	if err := godotenv.Load(); err != nil {
		logger.Error("couldn't load env variables")
	}
	path := os.Getenv("STORAGE_PATH")
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		defer file.Close()
		logger.Error("couldn't open file", "error", err.Error())
		os.Exit(1)
	}
	handler := handlers.NewHandler(file)
	if err := handler.DoAction(arguments); err != nil {
		logger.Error("internal error", "error", err.Error())
	}

}
