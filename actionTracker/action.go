package actionTracker

// Action represents a single occurrence an action was performed
type Action struct {
	// Name of action
	Name string `json:"action"`
	// Time in minutes the action was performed
	Time int `json:"time"`
}

// Average represents the average time in minutes an action has been performed
type Average struct {
	// Name of action
	Name string `json:"action"`
	// Average time in minutes the action was performed
	Average int `json:"avg"`
}
