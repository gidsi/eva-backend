package main

import "time"

type Calendar struct {
	Space	string
	Events	[]Event
}

type Event struct {
	Start         time.Time `json:"start"`
	ImportedId    string `json:"importId"`
	Status        string `json:"status"`
	Description   string `json:"description"`
	Location      string `json:"location"`
	Summary       string `json:"summary"`
	Rrule         string `json:"rrule"`
	Class         string `json:"class"`
	Url           string `json:"url"`
	Sequence      int `json:"sequence"`
	WholeDayEvent bool `json:"whileDayEvent"`

}