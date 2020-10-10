package civ

import (
	"fmt"
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
	var _ error

	sMsg := strings.Split(msg, " ")
	var m interface{}
	var n bool
	m, _ = strconv.Atoi(sMsg[2])
	m, n = m.(int)
	if n {
		cpp, _ = strconv.Atoi(sMsg[2])
		sMsg = sMsg[3:]
	} else {
		sMsg = sMsg[2:]
	}

	rand.Seed(time.Now().Unix())

	if (len(sMsg) * cpp) > len(nations) {
		return "Too many players!"
	}

	result := "```"

	players := make(map[string]string)

	for i := 0; i < cpp; i++ {
		for _, player := range sMsg {
			ele := rand.Intn(len(nations))
			nation := nations[ele]
			nations = remove(nations, ele)
			players[player] += nation + "\n"
		}
	}
	for key, value := range players {
		fmt.Print(key, value)
	}

	result += "```"

	return result
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
