package ergast

import "fmt"

const BASE_URL = "https://ergast.com/api/f1"
const RESPONSE_TYPE = ".json"

func GetSeasonSchedule(year int) []Race {
	// BASE_URL/2023.json
	requestUrl := fmt.Sprintf("%s/%d%s", BASE_URL, year, RESPONSE_TYPE)

	var data ErgastSeasonScheduleResponse
	SendRequest(requestUrl, &data)

	return data.MRData.RaceTable.Races
}

func GetRaceResults(year int, round int) []RaceResult {
	// BASE_URL/2023/4/results.json
	requestUrl := fmt.Sprintf("%s/%d/%d/results%s", BASE_URL, year, round, RESPONSE_TYPE)

	var data ErgastRaceResultsResponse
	SendRequest(requestUrl, &data)

	// For some reason Ergast returns an array of Races in the response.
	// Pick the first (only) element from that array
	return data.MRData.RaceTable.Races[0].Results
}

func GetQualifyingResults(year int, round int) []QualifyingResult {
	// BASE_URL/2023/4/qualifying.json
	requestUrl := fmt.Sprintf("%s/%d/%d/qualifying%s", BASE_URL, year, round, RESPONSE_TYPE)

	var data ErgastQualifyingResultsResponse
	SendRequest(requestUrl, &data)

	// For some reason Ergast returns an array of Races in the response.
	// Pick the first (only) element from that array
	return data.MRData.RaceTable.Races[0].QualifyingResults
}

func GetSprintResults(year int, round int) []SprintResult {
	// BASE_URL/2023/4/sprint.json
	requestUrl := fmt.Sprintf("%s/%d/%d/sprint%s", BASE_URL, year, round, RESPONSE_TYPE)

	var data ErgastSprintResultsResponse
	SendRequest(requestUrl, &data)

	// For some reason Ergast returns an array of Races in the response.
	// Pick the first (only) element from that array
	return data.MRData.RaceTable.Races[0].SprintResults
}
