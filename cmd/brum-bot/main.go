package main

import (
	"brum-bot/internal/app/brum"
	"flag"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	brum.Brum(Token)
}
