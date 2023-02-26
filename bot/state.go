package bot

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bodn/f1-go-bot/ergast"
)

var (
	State = initMyGlobalConstant()
)

type botStateType struct {
	// define your fields here
	CurrentSeason         int
	RaceSchedule          []ergast.Race
	MostRecentRace        ergast.Race
	MostRecentRaceResults []ergast.RaceResult
	NextRace              ergast.Race
}

func initMyGlobalConstant() botStateType {
	var instance botStateType

	// Season is equal to the year
	// currentYear := time.Now().Year()
	currentYear := 2023
	instance.CurrentSeason = currentYear

	// get the race schedule
	instance.RaceSchedule = ergast.GetSeasonSchedule(instance.CurrentSeason)

	// get the most recent race details
	instance.MostRecentRace = findMostRecentRace(instance.RaceSchedule)

	// get the most recent race result
	mostRecentRaceRound, err := strconv.Atoi(instance.MostRecentRace.Round)
	if err != nil {
		log.Println("There was error attempting to convert the race round string to an integer.")
	}

	if mostRecentRaceRound != 0 {
		instance.MostRecentRaceResults = ergast.GetRaceResults(instance.CurrentSeason, mostRecentRaceRound)
	}

	// get the next race details
	instance.NextRace = getNextRace(instance.RaceSchedule)

	return instance
}

func findMostRecentRace(races []ergast.Race) ergast.Race {
	currentTime := time.Now()
	utcTime := currentTime.UTC()

	var mostRecentRace *ergast.Race
	var mostRecentTime time.Time

	for i := range races {
		eventTime, err := time.Parse("2006-01-02 15:04:05Z", races[i].Date+" "+races[i].Time)
		if err != nil {
			fmt.Printf("Error parsing date time for event %v: %v\n", races[i], err)
			continue
		}

		if eventTime.After(mostRecentTime) && eventTime.Before(utcTime) {
			mostRecentRace = &races[i]
			mostRecentTime = eventTime
		}
	}

	if mostRecentRace == nil {
		fmt.Println("No events found between now and the past")
		return ergast.Race{}
	} else {
		fmt.Printf("Most recent event is %+v at %v\n", *mostRecentRace, mostRecentTime)
	}

	return *mostRecentRace
}

func getNextRace(races []ergast.Race) ergast.Race {
	now := time.Now().UTC()

	var nextEvent ergast.Race
	var found bool

	for _, event := range races {
		eventTime, err := time.Parse("2006-01-02 15:04:05Z", event.Date+" "+event.Time)
		if err != nil {
			return ergast.Race{}
		}

		if eventTime.After(now) {
			if !found {
				nextEvent = event
				found = true
			} else {
				nextEventTime, _ := time.Parse("2006-01-02 15:04:05Z", nextEvent.Date+" "+nextEvent.Time)
				if eventTime.Before(nextEventTime) {
					nextEvent = event
				}
			}
		}
	}

	return nextEvent
}
