package main

import (
	"fmt"

	"github.com/bodn/f1-go-bot/bot"
)

func main() {
	fmt.Println(bot.State.CurrentSeason)
	fmt.Println(bot.State.MostRecentRace)
	fmt.Println(len(bot.State.MostRecentRaceResults))
	fmt.Println(len(bot.State.MostRecentQualifyingResults))
	fmt.Println(len(bot.State.MostRecentSprintResults))
}
