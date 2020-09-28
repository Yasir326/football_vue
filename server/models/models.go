package models

import "time"

//Competition Struct
type Competition struct {
	ID          int       `json:"id"`
	Areas       Area      `json:"area"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Plan        string    `json:"plan"`
	LastUpdated time.Time `json:"lastUpdated"`
}

//Area Struct
type Area struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Season Struct
type Season struct {
	ID              int         `json:"id"`
	StartDate       string      `json:"startDate"`
	EndDate         string      `json:"endDate"`
	CurrentMatchday int         `json:"currentMatchday"`
	Winner          interface{} `json:"winner"`
	Standing        Standings   `json:"standings"`
}

//Standings Struct
type Standings []struct {
	Stage string      `json:"stage"`
	Type  string      `json:"type"`
	Group interface{} `json:"group"`
	Table Table       `json:"table"`
}

//Team Struct
type Team struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CrestURL string `json:"crestUrl"`
}

//Table Struct
type Table []struct {
	Position       int         `json:"position"`
	TeamStanding   Team        `json:"team"`
	PlayedGames    int         `json:"playedGames"`
	Form           interface{} `json:"form"`
	Won            int         `json:"won"`
	Draw           int         `json:"draw"`
	Lost           int         `json:"lost"`
	Points         int         `json:"points"`
	GoalsFor       int         `json:"goalsFor"`
	GoalsAgainst   int         `json:"goalsAgainst"`
	GoalDifference int         `json:"goalDifference"`
}
