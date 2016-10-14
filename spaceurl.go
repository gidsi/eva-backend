package main

type SpaceUrl struct {
	Url string `json:"url"`
	Validated bool `json:"validated"`
	LastUpdated int64 `json:"lastUpdated"`
}