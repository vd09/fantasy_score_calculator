package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// match_csv_reader
type MatchData1 struct {
	MatchID          string
	Team             string
	Name             string
	FoursAsBatter    int
	SixesAsBatter    int
	BallsAsBatter    int
	DotBallsAsBowler int
	MatchName        string
	MatchResult      string
	NameWDTitles     string
	OnFieldUmpires1  string
	OnFieldUmpires2  string
	OnFieldUmpires3  string
	OutDetails       string
	OversAsBowler    int
	PlayerRole       string
	RunsAsBatter     int
	RunsAsBowler     int
	Venue            string
	WicketAsBowler   int
}

func ReadDataFromCSV1() []MatchData {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return []MatchData{}
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return []MatchData{}
	}

	if len(records) == 0 {
		fmt.Println("No data found")
		return []MatchData{}
	}

	headers := records[0]
	columnIndex := make(map[string]int)
	for i, header := range headers {
		columnIndex[header] = i
	}

	var matchData []MatchData

	for i, record := range records {
		if i == 0 {
			continue
		}

		foursAsBatter, _ := strconv.Atoi(record[columnIndex["4s-as-batter"]])
		sixesAsBatter, _ := strconv.Atoi(record[columnIndex["6s-as-batter"]])
		ballsAsBatter, _ := strconv.Atoi(record[columnIndex["balls-as-batter"]])
		dotBallsAsBowler, _ := strconv.Atoi(record[columnIndex["dotballs-as-bowler"]])
		oversAsBowler, _ := strconv.Atoi(record[columnIndex["overs-as-bowler"]])
		runsAsBatter, _ := strconv.Atoi(record[columnIndex["runs-as-batter"]])
		runsAsBowler, _ := strconv.Atoi(record[columnIndex["runs-as-bowler"]])
		wicketAsBowler, _ := strconv.Atoi(record[columnIndex["wicket-as-bowler"]])

		matchData = append(matchData, MatchData{
			MatchID:          record[columnIndex["match-id"]],
			Team:             record[columnIndex["team"]],
			Name:             record[columnIndex["name"]],
			FoursAsBatter:    foursAsBatter,
			SixesAsBatter:    sixesAsBatter,
			BallsAsBatter:    ballsAsBatter,
			DotBallsAsBowler: dotBallsAsBowler,
			MatchName:        record[columnIndex["match-name"]],
			MatchResult:      record[columnIndex["match-result"]],
			NameWDTitles:     record[columnIndex["name-wd-titles"]],
			OnFieldUmpires1:  record[columnIndex["on-field-umpires1"]],
			OnFieldUmpires2:  record[columnIndex["on-field-umpires2"]],
			OnFieldUmpires3:  record[columnIndex["on-field-umpires3"]],
			OutDetails:       record[columnIndex["out-details"]],
			OversAsBowler:    oversAsBowler,
			PlayerRole:       record[columnIndex["player-role"]],
			RunsAsBatter:     runsAsBatter,
			RunsAsBowler:     runsAsBowler,
			Venue:            record[columnIndex["venue"]],
			WicketAsBowler:   wicketAsBowler,
		})
	}
	return matchData
}

func GroupDataByMatchID1(matchData []MatchData) map[string][]MatchData {
	groupedData := make(map[string][]MatchData)

	for _, data := range matchData {
		groupedData[data.MatchID] = append(groupedData[data.MatchID], data)
	}

	return groupedData
}

func main1() {
	matchData := ReadDataFromCSV()

	// Group data by match ID
	groupedData := GroupDataByMatchID(matchData)

	// Print grouped data
	for matchID, records := range groupedData {
		fmt.Printf("Match ID: %s\n", matchID)
		//fmt.Printf("Len: %s\n", len(records))
		for _, data := range records {
			fmt.Printf("%+v\n", data)
		}
		fmt.Println()
	}
}
