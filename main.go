package main

import (
	"github.com/ltakato/gopportunities/config"
	"github.com/ltakato/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
