package main

import (
	"github.com/joho/godotenv"

	"github.com/phuwn/smart-schedule/pkg/server"
	"github.com/phuwn/smart-schedule/pkg/store"
	"github.com/phuwn/tools/log"
	"github.com/phuwn/tools/util"
)

// init server stuff
func init() {
	env := util.Getenv("RUN_MODE", "")
	if env == "local" || env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	store := store.New()
	server.NewServerCfg(store)
}