package backend

import (
	"context"
	"fmt"
	"log"
	"time-tracker/backend/sources"
)

type App struct {
	ctx     context.Context
	sources *sources.Sources
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	src, err := sources.InitiateSources()
	if err != nil {
		log.Fatal("Failed to initiate sources:", err)
	}
	a.sources = src
}

// Greet for testing
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
