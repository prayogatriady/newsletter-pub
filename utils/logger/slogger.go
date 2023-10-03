package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

// GeneralLogger exported
var Slogger *slog.Logger

func init() {
	logDir := "log"

	// Create a directory for logs if it doesn't exist
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, os.ModePerm)
	}

	// Specify the absolute path for log files
	absPath, err := filepath.Abs("log/") // Adjust this path as needed
	if err != nil {
		fmt.Println("Error reading given path:", err)
	}

	// Start a goroutine to monitor the date and create new log files when needed
	go dailysLogRotation(absPath)

	// Initialize the loggers
	initSloggers(absPath)
}

func initSloggers(absPath string) {

	// Get the current date
	currentDate := time.Now().Format("20060102")

	// Create separate log files for GeneralLogger and ErrorLogger with the current date in the name
	generalLogFileName := fmt.Sprintf("newsletterPub-%s.log", currentDate)

	dir := filepath.Join(absPath, generalLogFileName)

	generalLog, err := os.OpenFile(dir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	logHandler := slog.NewJSONHandler(generalLog, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}).WithAttrs([]slog.Attr{
		slog.Group("app",
			slog.String("name", Cfg.App.AppName),
			slog.String("version", Cfg.App.AppVersion),
			slog.String("timezone", Cfg.App.Timezone),
		),
	})

	Slogger = slog.New(logHandler)

}

func dailysLogRotation(absPath string) {
	for {
		// Sleep until the next day
		nextDay := time.Now().Add(24 * time.Hour)
		nextDay = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, nextDay.Location())
		sleepDuration := time.Until(nextDay)
		time.Sleep(sleepDuration)

		// Create new log files for the new day
		initSloggers(absPath)
	}
}
