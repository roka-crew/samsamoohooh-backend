package main

import (
	"log"

	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/pkg/config"
	"github.com/roka-crew/samsamoohooh-backend/pkg/database/postgres"
)

func main() {
	config, err := config.New("./configs/config.yaml")
	if err != nil {
		log.Panicf("failed to new config: %v\n", err)
	}

	db, err := postgres.New(config)
	if err != nil {
		log.Panicf("failed to connection postgres: %v\n", err)
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Panicf("failed to auto migrate: %v\n", err)
	}
}
