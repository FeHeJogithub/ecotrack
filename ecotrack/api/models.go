package api

type Habit struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Points int    `json:"points"`
}

var habits []Habit
var idCounter = 1
