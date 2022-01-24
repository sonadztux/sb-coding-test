package models

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Movie struct {
	ImdbID string `json:imdbID`
	Title  string `json:Title`
	Year   string `json:Year`
	Type   string `json:Type`
	Poster string `json:Poster`
}

type MovieList struct {
	Movie []Movie `json:"Search"`
}

type MovieDetail struct {
	ImdbID     string  `json:"imdbID"`
	Title      string  `json:"Title"`
	Year       string  `json:"Year"`
	Rated      string  `json:"Rated"`
	Released   string  `json:"Released"`
	Runtime    string  `json:"Runtime"`
	Genre      string  `json:"Genre"`
	Director   string  `json:"Director"`
	Writer     string  `json:"Writer"`
	Actors     string  `json:"Actors"`
	Plot       string  `json:"Plot"`
	Language   string  `json:"Language"`
	Country    string  `json:"Country"`
	Awards     string  `json:"Awards"`
	Poster     string  `json:"Poster"`
	Ratings    Ratings `json:"Ratings"`
	Metascore  string  `json:"Metascore"`
	ImdbRating string  `json:"imdbRating"`
	ImdbVotes  string  `json:"imdbVotes"`
	Type       string  `json:"Type"`
	DVD        string  `json:"DVD"`
	BoxOffice  string  `json:"BoxOffice"`
	Production string  `json:"Production"`
	Website    string  `json:"Website"`
}

type Ratings []struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

type LogCall struct {
	ID         int
	CallType   string
	ParamQuery string
	Status     int
	CreatedAt  time.Time
}
