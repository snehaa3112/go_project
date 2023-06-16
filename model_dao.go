package models

import (
	"project/connection"
)

func GetLongestDurationMovies() (*[]LongestDurationMoviesResponse, error) {
	var movies []LongestDurationMoviesResponse

	query := "SELECT tconst, primaryTitle as primary_title, runtimeMinutes as runtime_minutes, genres FROM movies ORDER BY runtimeMinutes DESC LIMIT 10"
	res := connection.DB.Raw(query).Order("runtimeMinutes DESC").Find(&movies)
	if res.Error != nil {
		return nil, res.Error
	}

	return &movies, nil
}

func GetTopRatedMovies() (*[]TopRatedMoviesResponse, error) {
	var movies []TopRatedMoviesResponse

	query := "SELECT movies.tconst, movies.primaryTitle as primary_title, movies.genres, ratings.averageRating as average_rating FROM movies INNER JOIN ratings ON movies.tconst = ratings.tconst WHERE ratings.averageRating > 6.0 ORDER BY ratings.averageRating ASC;"
	res := connection.DB.Raw(query).Order("averageRating ASC").Find(&movies)
	if res.Error != nil {
		return nil, res.Error
	}

	return &movies, nil
}

func GetGenreMoviesWithSubtotals() (*[]GenreMoviesWithSubtotals, error) {
	var movies []GenreMoviesWithSubtotals

	query := "SELECT m.genres, m.primaryTitle as primary_title, SUM(r.numVotes) AS subtotal FROM movies m INNER JOIN ratings r ON m.tconst = r.tconst GROUP BY m.genres, m.primaryTitle WITH ROLLUP ORDER BY m.genres, m.primaryTitle;"
	res := connection.DB.Raw(query).Order("averageRating ASC").Find(&movies)
	if res.Error != nil {
		return nil, res.Error
	}

	return &movies, nil
}

func AddNewMovie(body *Movie) error {
	res := connection.DB.Table("movies").Create(&body)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func UpdateRunTime() error {
	query := "UPDATE movies SET runtimeMinutes = CASE WHEN genres = 'Documentary' THEN runtimeMinutes + 15 WHEN genres = 'Animation' THEN runtimeMinutes + 30 ELSE runtimeMinutes + 45 END;"
	res := connection.DB.Exec(query)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
