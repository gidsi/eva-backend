package main

import "time"

type Calendar struct {
	Space	string
	Events	[]Event
}

type Event struct {
	Start         time.Time
	ImportedId    string
	Status        string
	Description   string
	Location      string
	Summary       string
	Rrule         string
	Class         string
	Sequence      int
	WholeDayEvent bool

}