package main

import (
	"fantasy_score_calculator/csv_reader"
	"fantasy_score_calculator/domain"
	"fmt"
	"strconv"

	gr_variable "github.com/vd09/gr-variable"
)

func handleRecord(matchIdToChannel map[string]gr_variable.GrChannel[domain.MatchPlayerData]) csv_reader.RecordHandler {
	return func(record []string, columnIndex map[string]int) error {
		foursAsBatter, _ := strconv.Atoi(record[columnIndex["4s-as-batter"]])
		sixesAsBatter, _ := strconv.Atoi(record[columnIndex["6s-as-batter"]])
		ballsAsBatter, _ := strconv.Atoi(record[columnIndex["balls-as-batter"]])
		dotBallsAsBowler, _ := strconv.Atoi(record[columnIndex["dotballs-as-bowler"]])
		oversAsBowler, _ := strconv.ParseFloat(record[columnIndex["overs-as-bowler"]], 64)
		runsAsBatter, _ := strconv.Atoi(record[columnIndex["runs-as-batter"]])
		runsAsBowler, _ := strconv.Atoi(record[columnIndex["runs-as-bowler"]])
		wicketAsBowler, _ := strconv.Atoi(record[columnIndex["wicket-as-bowler"]])

		matchData := domain.MatchPlayerData{
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
		}

		ch, ok := matchIdToChannel[matchData.MatchID]
		if !ok {
			ch = gr_variable.NewGrChannelWithLength[domain.MatchPlayerData](25)
			matchIdToChannel[matchData.MatchID] = ch
		}
		ch.WriteValue(matchData)
		// You can now use matchData as needed, e.g., store it in a slice or process it further.
		fmt.Println(matchData)

		return nil
	}
}

func main() {
	//ctx := context.Background()
	matchIdToChannel := map[string]gr_variable.GrChannel[domain.MatchPlayerData]{}
	if err := csv_reader.ReadCSV("data.csv", handleRecord(matchIdToChannel)); err != nil {
		fmt.Println("Error:", err)
	}
}
