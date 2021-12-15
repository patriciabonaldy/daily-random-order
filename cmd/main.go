package main

import (
	bootstrap "daily-random-order/cmd/boostrap"
	"daily-random-order/cmd/config"
	"log"
)

func main() {
	config.Load()
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
