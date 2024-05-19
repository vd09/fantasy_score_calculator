package worker

//
//import (
//	"context"
//	"fantasy_score_calculator/domain"
//	"fantasy_score_calculator/match"
//
//	gr_variable "github.com/vd09/gr-variable"
//)
//
//func abcd(ctx context.Context, matchDetails match.Match, matchPlayerDataChannel gr_variable.ReadOnlyGrChannel[domain.MatchPlayerData]) {
//	lMatchPlayerData, ok := matchPlayerDataChannel.ReadValue()
//	if !ok {
//		return
//	}
//	matchDetails := match.NewMatchWithFirstPlayer(lMatchPlayerData)
//
//	for {
//		select {
//		case <-ctx.Done():
//			return
//		case matchPlayerData, ok := <-matchPlayerDataChannel.Receive():
//			if !ok {
//				return
//			}
//			matchDetails.AddPlayer(matchPlayerData)
//		}
//	}
//}
