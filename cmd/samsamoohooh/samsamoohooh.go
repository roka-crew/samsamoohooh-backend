package main

import (
	"fmt"
	"log"

	"github.com/roka-crew/samsamoohooh-backend/pkg/config"
)

func main() {
	config, err := config.New("configs/config.yaml")
	if err != nil {
		log.Panicf("failed to new config")
	}

	fmt.Println("parsed config: ", config)
}
