package civ

import (
	"math/rand"
	"strconv"
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

//Civ ...
func Civ(msg string) string {
	var cpp int = 1
	rand.Seed(time.Now().Unix())

	sMsg := strings.Split(msg, " ")
	_, err := strconv.Atoi(sMsg[2])
	if err == nil {
		cpp, _ = strconv.Atoi(sMsg[2])
		sMsg = sMsg[3:]
	} else {
		sMsg = sMsg[2:]
	}

	if (len(sMsg) * cpp) > len(nations) {
		return "Too many players!"
	}

	players := make([]string, len(sMsg), len(sMsg))
	for i := 0; i < len(sMsg); i++ {
		players[i] = sMsg[i] + "\t"
	}

	for i := 0; i < cpp; i++ {
		for j := 0; j < len(players); j++ {
			ele := rand.Intn(len(nations))
			nation := nations[ele]
			nations = remove(nations, ele)
			players[j] += nation
			if j < len(players) && i < (cpp-1) {
				players[j] += " - "
			}
			if i == (cpp-1) && j < (len(players)-1) {
				players[j] += "\n"
			}
		}
	}

	result := "```"
	for i := 0; i < len(sMsg); i++ {
		result += players[i]
	}
	result += "```"

	return result
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
