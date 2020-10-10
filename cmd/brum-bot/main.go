package main

import (
	"brum-bot/internal/app/brum"
	"io/ioutil"
	"os"
)

func main() {
	token := os.Getenv("DISCORD_KEY")
	if token == "" {
		dat, _ := ioutil.ReadFile("./.discord-key")
		token = string(dat)
	}

	brum.Brum(token)
}
