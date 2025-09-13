package main

import (
	"fmt"
	"log"

	"github.com/zeuge/hw-go/03-env/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	fmt.Printf("Loaded config: %+v\n", cfg)
}
