package main

import (
	"log/slog"
	"os"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("failed running command", "error", err)
		os.Exit(1)
	}
}
