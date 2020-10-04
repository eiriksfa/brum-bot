package civ

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var nations = []string{
	"Saladin",
	"Amanitore",
	"Hojo Tokimune",
	"Robert the Bruce",
	"Montezuma",
	"Gorgo",
	"Pericles",
	"Shaka",
	"Suleiman",
	"John Curtin",
	"Matthias",
	"Teddy Roosevelt  ( base game)",
	"Pachacuti",
	"Menelik II",
	"Rough Rider Teddy",
	"Trajan",
	"Wilhelmina",
	"Genghis Khan",
	"Gitarja",
	"Poundmaker",
	"Gilgamesh",
	"Jadwiga",
	"Pedro",
	"Tomyris",
	"Kristina",
	"Mansa Musa",
	"Qin Shi Quang",
	"Alexander",
	"Wilfrid Laurier",
	"Lady Six Sky",
	"Catherine",
	"Lautaro",
	"Cleopatra",
	"Mvemba ",
	"Victoria",
	"Dido",
	"Eleanor (France)",
	"Gaul",
	"Byzantium"}

func Civ(msg string) string {
	sMsg := strings.Split(msg, " ")
	// If the message is "ping" reply with "Pong!"
	sMsg = sMsg[2:]

	rand.Seed(time.Now().Unix())

	if len(sMsg) > len(nations) {
		return "To many players!"
	}

	result := "```"

	for _, player := range sMsg {
		ele := rand.Intn(len(nations))
		nation := nations[ele]
		nations = remove(nations, ele)
		result += fmt.Sprintf("%s - %s\n", player, nation)
	}

	result += "```"

	return result
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
