package player

import (
	"fantasy_score_calculator/domain"
	"regexp"
	"strings"
)

const (
	caseInsensitiveFlag = `(?i)`
)

var (
	patternParts = []string{
		`(?P<runOut>run out \((?P<runOutDirect>[^/()]+)(?:/(?P<runOutIndirect>[^()]+))?\))`, // Run out
		`(?P<hitWicket>Hit Wicket(?: b (?P<hitWicketBowler>.*))?)`,                          // Hit wicket
		`(?P<lbw>lbw (?P<lbwBowler>.*))`,                                                    // LBW
		`(?P<stumped>st (?P<stumpedBy>.*?) b (?P<stumpedBowler>.*))`,                        // Stumped
		`(?P<caughtAndBowled>c & b (?P<caughtAndBowledBowler>.*))`,                          // Caught and bowled
		`(?P<caught>c (?P<caughtBy>.*?) b (?P<caughtBowler>.*))`,                            // Caught
		`(?P<bowled>b (?P<bowledBy>.*))`,                                                    // Bowled
	}
	wicketPattern      = strings.Join(patternParts, "|")
	wicketPatternRegex = regexp.MustCompile(caseInsensitiveFlag + wicketPattern)
)

// ParseWicketInfo parses the wicket information from the given line.
func ParseWicketInfo(wicketDetailsLine string) domain.WicketInfo {
	return getWicketInfo(extractWicketNamedToValue(wicketDetailsLine))
}

// extractWicketNamedToValue returns a map of named capture groups to their values.
func extractWicketNamedToValue(line string) map[string]string {
	matches := make(map[string]string)
	if match := wicketPatternRegex.FindStringSubmatch(line); match != nil {
		groupNames := wicketPatternRegex.SubexpNames()
		for i, name := range groupNames {
			if i != 0 && name != "" {
				matches[name] = match[i]
			}
		}
	}
	return matches
}

// getWicketInfo returns a WicketInfo struct populated based on the matched capture groups.
func getWicketInfo(matches map[string]string) domain.WicketInfo {
	var info domain.WicketInfo

	switch {
	case matches["runOut"] != "":
		info.RunOutDirect = matches["runOutDirect"]
		info.RunOutIndirect = matches["runOutIndirect"]
	case matches["hitWicket"] != "":
		info.HitWicket = true
		info.Bowler = matches["hitWicketBowler"]
	case matches["lbw"] != "":
		info.LBW = true
		info.Bowler = matches["lbwBowler"]
	case matches["stumped"] != "":
		info.Stumped = matches["stumpedBy"]
		info.Bowler = matches["stumpedBowler"]
	case matches["caughtAndBowled"] != "":
		info.Caught = matches["caughtAndBowledBowler"]
		info.Bowler = matches["caughtAndBowledBowler"]
	case matches["caught"] != "":
		info.Caught = matches["caughtBy"]
		info.Bowler = matches["caughtBowler"]
	case matches["bowled"] != "":
		info.Bowler = matches["bowledBy"]
	}

	return info
}
