package actionTracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

// actionStat represents statistics of a performed action
type actionStat struct {
	totalTime int
	count int
}

// actionStats is a map containing actionStat with the key being the name of the action
type actionStats map[string]*actionStat

// ActionTracker keeps track of statistics of performed actions
type ActionTracker struct {
	actionStats actionStats
	sync.Mutex
}

func NewTracker() *ActionTracker {
	at := &ActionTracker{
		actionStats: make(map[string]*actionStat),
	}
	return at
}

// AddAction accepts a unencoded JSON representation of the Action struct and performs said action.
// A performed action's statistics are tallied and added to the ActionTracker
func (t *ActionTracker) AddAction(input string) error {
	var action Action
	err := json.Unmarshal([]byte(input), &action)
	if err != nil {
		return errors.New("bad JSON format for action")
	}

	fmt.Printf("adding action: %v\n", action)

	t.tallyAction(action)

	return nil
}

// tallyAction tallies the passed in action's attributes to the ActionTracker's statistics
func (t *ActionTracker) tallyAction(action Action) {
	t.Lock()
	defer t.Unlock()

	 if stat, ok := t.actionStats[action.Name]; ok {
		stat.count++
		stat.totalTime += action.Time
	} else { // Add new action
		t.actionStats[action.Name] = &actionStat{
			totalTime: action.Time,
			count:     1,
		}
	}
}

// GetStats returns a json encoded array fo the Average struct which holds statistics
// for each action that has been performed
func (t *ActionTracker) GetStats() string {
	t.Lock()

	if len(t.actionStats) == 0 {
		return ""
	}

	// Iterate through the actionStats map and make an Average struct for each action
	var averages []Average
	for name, actionStat := range t.actionStats {
		avg := actionStat.totalTime / actionStat.count
		actionAverage := Average{
			Name:    name,
			Average: avg,
		}
		averages = append(averages, actionAverage)
	}

	// Release lock as soon as possible instead of deferring the exit of the function
	t.Unlock()

	averagesJson, err := json.Marshal(averages)
	if err != nil {
		fmt.Printf("could not encode action averages to JSON %v\n", err)
	}

	return string(averagesJson)
}
