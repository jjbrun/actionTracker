package actionTracker

import (
	"testing"
)

var jump100 = `{"action":"jump", "time":100}`
var jump200 = `{"action":"jump", "time":200}`
var jumpNeg50 = `{"action":"jump", "time":-50}`
var run75 = `{"action":"run", "time":75}`
var badJSON = `{"garbage"}`

func TestTracker_AddActionSingleAction(t *testing.T) {
	at := NewTracker()
	actionName := "jump"

	expected := actionStat{
		totalTime: 100,
		count:     1,
	}

	_ = at.AddAction(jump100)

	actual := *at.actionStats[actionName]

	if expected != actual {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}

func TestTracker_AddActionMultipleSameActions(t *testing.T) {
	at := NewTracker()
	actionName := "jump"

	expected := actionStat{
		totalTime: 300,
		count:     2,
	}

	_ = at.AddAction(jump100)
	_ = at.AddAction(jump200)

	actual := *at.actionStats[actionName]

	if expected != actual {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}

func TestTracker_AddActionMultipleSameActionsWithNegativeActionTime(t *testing.T) {
	at := NewTracker()
	actionName := "jump"

	expected := actionStat{
		totalTime: 50,
		count:     2,
	}

	_ = at.AddAction(jump100)
	_ = at.AddAction(jumpNeg50)

	actual := *at.actionStats[actionName]

	if expected != actual {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}

func TestTracker_AddActionMultipleDifferentActions(t *testing.T) {
	at := NewTracker()

	expectedJump := actionStat{
		totalTime: 100,
		count:     1,
	}

	expectedRun := actionStat{
		totalTime: 75,
		count:     1,
	}

	_ = at.AddAction(jump100)
	_ = at.AddAction(run75)

	actualJump := *at.actionStats["jump"]
	actualRun := *at.actionStats["run"]

	if expectedJump != actualJump {
		t.Error("Expected:", expectedJump, "Actual:", actualJump)
	}

	if expectedRun != actualRun {
		t.Error("Expected:", expectedJump, "Actual:", actualRun)
	}
}

func TestTracker_AddActionBadJSON(t *testing.T) {
	at := NewTracker()

	err := at.AddAction(badJSON)
	if err == nil {
		t.Error("Expected error for bad JSON")
	}
}

func TestTracker_GetStats(t *testing.T) {
	at := NewTracker()

	_ = at.AddAction(jump100)
	_ = at.AddAction(jump200)
	_ = at.AddAction(run75)
	_ = at.AddAction(run75)

	actual := `[{"action":"jump","avg":150},{"action":"run","avg":75}]`

	expected := at.GetStats()

	if expected != actual {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}

func TestTracker_GetStatsNoActions(t *testing.T) {
	at := NewTracker()

	actual := ""
	expected := at.GetStats()

	if expected != actual {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}
