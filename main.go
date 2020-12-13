package main

import (
	"actionTracker/actionTracker"
	"fmt"
)

func main() {
	jump100 := `{"action":"jump", "time":100}`
	jump200 := `{"action":"jump", "time":200}`
	run75 := `{"action":"run", "time":75}`

	at := actionTracker.NewTracker()

	at.AddAction(jump100)
	at.AddAction(jump200)
	at.AddAction(run75)

	fmt.Println(at.GetStats())
}
