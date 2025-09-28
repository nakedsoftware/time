package main

import (
	"context"
	"log"

	"github.com/nakedsoftware/time/internal/cli"
)

func main() {
	if err := cli.Execute(context.Background()); err != nil {
		log.Fatalf("The program failed to execute successfully: %v", err)
	}
}
