package civ

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Ranking struct {
	Rating      string
	Description string
	Victory     string
}

type Leader struct {
	Name       string
	Country    string
	Overall    string
	Domination string
	Science    string
	Culture    string
	Religion   string
	Diplomacy  string
}

type NationList struct {
	Rankings []Ranking
	Leaders  []Leader
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func FilterLeaders(leaders []Leader, ranks []string) []Leader {
	vsf := make([]Leader, 0)
	for _, leader := range leaders {
		if stringInSlice(leader.Overall, ranks) {
			vsf = append(vsf, leader)
		}
	}
	return vsf
}

func FilterLeadersOnName(leaders []Leader, name string) []Leader {
	vsf := make([]Leader, 0)
	for _, leader := range leaders {
		if leader.Name == name {
			vsf = append(vsf, leader)
		}
	}
	return vsf
}

func Rankings() string {
	dat, _ := ioutil.ReadFile("./assets/nations.json")
	var nationList NationList
	_ = json.Unmarshal(dat, &nationList)
	result := "```"
	for _, ranking := range nationList.Rankings {
		result = result + "Rating: " + ranking.Rating + "\n"
		result = result + "Description: " + ranking.Description + "\n"
		result = result + "Victory: " + ranking.Victory + "\n"
		result = result + "\n"
	}
	return result + "```"
}

func Leaders(r []string) string {
	dat, _ := ioutil.ReadFile("./assets/nations.json")
	var nationList NationList
	_ = json.Unmarshal(dat, &nationList)
	result := "```"
	var leaders []Leader
	ranks := strings.Join(r[:], ",")
	ranks = strings.ReplaceAll(ranks, ",", " ")
	if len(ranks) == 1 {
		rankings := strings.Split(ranks, ",")
		leaders = FilterLeaders(nationList.Leaders, rankings)
	} else {
		leaders = FilterLeadersOnName(nationList.Leaders, ranks)
	}

	for _, leader := range leaders {
		result = result + "Name: " + leader.Name + "\n"
		result = result + "Country: " + leader.Country + "\n"
		result = result + "Rank: " + leader.Overall + "\n"
		result = result + "Domination: " + leader.Domination + "\n"
		result = result + "Science: " + leader.Science + "\n"
		result = result + "Culture: " + leader.Culture + "\n"
		result = result + "Religion: " + leader.Religion + "\n"
		result = result + "Diplomacy: " + leader.Diplomacy + "\n"
		result = result + "\n"
	}
	return result + "```"
}

func Assign(players []string, ranks string) string {

	rand.Seed(time.Now().Unix())

	dat, _ := ioutil.ReadFile("./assets/nations.json")
	var nationList NationList
	_ = json.Unmarshal(dat, &nationList)

	rankings := strings.Split(ranks, "")

	leaders := FilterLeaders(nationList.Leaders, rankings)

	if len(players) > len(leaders) {
		return "To many players!"
	}

	result := "```"

	for _, player := range players {
		ele := rand.Intn(len(leaders))
		nation := leaders[ele]
		leaders = remove(leaders, ele)
		result += fmt.Sprintf("%s - %s\n", player, nation.Name)
	}

	result += "```"

	return result
}

func remove(s []Leader, i int) []Leader {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
