package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "exiting: %v", err)
	}
}

func run() error {
	s := http.Server{
		Addr:    ":9090",
		Handler: Routes(),
	}

	if err := s.ListenAndServe(); err != nil {
		return fmt.Errorf("server failed: %w", err)
	}

	return nil
}
