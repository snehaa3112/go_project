package models

// type Rating struct {
// 	Tconst        string  `json:"tconst"`
// 	AverageRating float32 `json:"averageRating"`
// 	NumVotes      int     `json:"numVotes"`
// }

type Movie struct {
	Tconst         string `binding:"required"`
	Titletype      string `binding:"required"`
	Primarytitle   string `binding:"required"`
	Runtimeminutes int    `binding:"required"`
	Genres         string `binding:"required"`
}

type LongestDurationMoviesResponse struct {
	Tconst         string `json:"tconst"`
	PrimaryTitle   string `json:"primary_title"`
	RuntimeMinutes int    `json:"runtime_minutes"`
	Genres         string `json:"genres"`
}

type TopRatedMoviesResponse struct {
	Tconst        string  `json:"tconst"`
	PrimaryTitle  string  `json:"primary_title"`
	Genres        string  `json:"genres"`
	AverageRating float32 `json:"average_rating"`
}

type GenreMoviesWithSubtotals struct {
	Genres       string `json:"genres"`
	PrimaryTitle string `json:"primary_title"`
	Subtotal     int    `json:"subtotal"`
}

type GenreMoviesWithSubtotalsResponse struct {
	Record []GenreMoviesWithSubtotals `json:"movies"`
	Total  int                        `json:"total"`
}
