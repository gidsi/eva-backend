package main

import "net/http"

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}


type Routes []Route

var IndexRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"SpaceDataIndex",
		"GET",
		"/spaces",
		SpaceDataIndex,
	},
	Route{
		"SpaceUrlIndex",
		"GET",
		"/urls",
		SpaceUrlIndex,
	},
	Route{
		"CalendarIndex",
		"GET",
		"/calendar",
		CalendarIndex,
	},
	Route{
		"SpaceUrlAdd",
		"POST",
		"/urls",
		SpaceUrlAdd,
	},
	Route{
		"SpaceUrlUpdate",
		"PUT",
		"/urls/{SharedSecret}",
		SpaceUrlUpdate,
	},
	Route{
		"RefreshData",
		"GET",
		"/refresh",
		refreshData,
	},
}