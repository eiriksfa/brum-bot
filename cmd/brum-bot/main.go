package main

import (
	"brum-bot/internal/app/brum"
	"flag"
	"os"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", os.Getenv("DISCORD_KEY"), "Bot Token")
	flag.Parse()
}

func main() {
	brum.Brum(Token)
}
