package main

func GetCoverStations(stations map[string][]string, statesNeed []string) []string {
	var finalStations []string

	for len(statesNeed) > 0 {
		var bestStation string
		var statesCovered []string

		for station, states := range stations {
			covered := equalData(statesNeed, states)
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}

		statesNeed = removeData(statesNeed, statesCovered)
		finalStations = append(finalStations, bestStation)
	}

	return finalStations
}

func equalData(statesNeed []string, states []string) []string {
	var covered []string
	for _, stateNeed := range statesNeed {
		for _, state := range states {
			if state == stateNeed {
				covered = append(covered, state)
			}
		}
	}
	return covered
}

func removeData(statesNeed []string, statesCovered []string) []string {
	for _, covered := range statesCovered {
		for i, stateNeed := range statesNeed {
			if covered == stateNeed {
				statesNeed = append(statesNeed[:i], statesNeed[i+1:]...)
			}
		}
	}
	return statesNeed
}
