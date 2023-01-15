package main

import (
	"wb/rest-api/internal/config"
	"wb/rest-api/internal/server"
	"wb/rest-api/internal/storage/database"
	"wb/rest-api/pkg/logging"
)

const (
	cfgPath = "./config.json"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("------------------------------------------------------------")
	logger.Info("NEW APPLICATION")

	cfg, err := config.GetConfig(cfgPath, logger)
	if err != nil {
		logger.Fatal(err)
	}

	db, err := database.NewDatabaseConnection(cfg.DB, logger)
	if err != nil {
		logger.Fatal(err)
	}

	srv := server.NewServer(db, logger)

	srv.Run(cfg.Listen)
}
