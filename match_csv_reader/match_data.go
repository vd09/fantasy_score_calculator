package main

import (
	"context"
	"errors"
	"fantasy_score_calculator/match"
	"fantasy_score_calculator/player"
	"fmt"
	"sync"

	"github.com/vd09/csvutils"
	"github.com/vd09/gr_worker/worker_pool"
)

type MatchPlayerDataProcessor struct {
	worker_pool.WorkerPool
	*match.Match
}

func NewMatchPlayerDataProcessor(ctx context.Context, matchPlayerData *player.MatchPlayer) *MatchPlayerDataProcessor {
	adapter, err := worker_pool.NewWorkerPoolAdapter(
		worker_pool.WithContext(ctx),
		worker_pool.WithMaxWorkers(1),
		worker_pool.WithMaxTasks(25),
	)
	if err != nil {
		return nil
	}

	return &MatchPlayerDataProcessor{
		WorkerPool: adapter,
		Match:      match.NewMatch(matchPlayerData),
	}
}

func handleRecord(
	ctx context.Context,
	rwMx *sync.RWMutex,
	matchIdToProcessor map[string]*MatchPlayerDataProcessor,
) csvutils.RecordHandler {
	return func(record interface{}) error {
		matchPlayerData, ok := record.(*player.MatchPlayer)
		if !ok {
			return errors.New("invalid record type")
		}

		rwMx.RLock()
		processor, ok := matchIdToProcessor[matchPlayerData.MatchID]
		rwMx.RUnlock()
		if !ok {
			processor = NewMatchPlayerDataProcessor(ctx, matchPlayerData)
			if processor == nil {
				return fmt.Errorf("not able to init MatchPlayerDataProcessor")
			}
			rwMx.Lock()
			matchIdToProcessor[matchPlayerData.MatchID] = processor
			rwMx.Unlock()
		}
		processor.AddTask(processor.AddPlayer, matchPlayerData)
		return nil
	}
}

func printHandleRecord(record interface{}) error {
	matchPlayerData, ok := record.(*player.MatchPlayer)
	if !ok {
		return errors.New("invalid record type")
	}
	fmt.Printf("----%#v\n", matchPlayerData)
	fmt.Printf("--------%#v\n", matchPlayerData.PlayerStats)
	fmt.Printf("------------%#v\n", matchPlayerData.PlayerStats.BattingStats)
	fmt.Printf("------------%#v\n", matchPlayerData.PlayerStats.BowlingStats)
	fmt.Printf("------------%#v\n", matchPlayerData.PlayerStats.FieldingStats)
	return nil
}

func main() {
	ctx := context.Background()
	rwMx := sync.RWMutex{}
	matchIdToProcessor := make(map[string]*MatchPlayerDataProcessor)

	csvRecordHandler := csvutils.WithHandler(handleRecord(ctx, &rwMx, matchIdToProcessor))
	//csvRecordHandler := csvutils.WithHandler(printHandleRecord)
	if err := csvutils.ReadCSV("resource/data.csv", &player.MatchPlayer{}, csvRecordHandler); err != nil {
		fmt.Println("Error:", err)
	}

	for _, processor := range matchIdToProcessor {
		processor.AddTask(processor.UpdateScores)
	}

	for _, processor := range matchIdToProcessor {
		processor.WaitAndStop()
	}

	for _, processor := range matchIdToProcessor {
		//fmt.Printf("%-25s | %#v\n", "Name", "Score")
		//fmt.Println("----------------------------------------")
		//for name, mp := range processor.MatchPlayerMap {
		//	fmt.Printf("%-25s | %5d\n", name, mp.PlayerStats.Score)
		//}

		mpSlice := make([]*player.MatchPlayer, 0)
		for _, mp := range processor.MatchPlayerMap {
			mpSlice = append(mpSlice, mp)
		}

		csvutils.WriteCSV("resource/final_output.csv", mpSlice)
	}
}
