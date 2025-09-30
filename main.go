package main

import (
	"fmt"
	"log/slog"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}

	slog.Info("Code executed successfully")
}

func run() error {
	fmt.Println("Hello, World!")
	return nil
}
