package service

import (
	models "project/model"
)

func GetLongestDurationMovies() (*[]models.LongestDurationMoviesResponse, error) {
	movies, err := models.GetLongestDurationMovies()
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func GetTopRatedMovies() (*[]models.TopRatedMoviesResponse, error) {
	movies, err := models.GetTopRatedMovies()
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func GetGenreMoviesWithSubtotals() (*[]models.GenreMoviesWithSubtotalsResponse, error) {
	movies, err := models.GetGenreMoviesWithSubtotals()
	if err != nil {
		return nil, err
	}

	groupedResults := make(map[string][]models.GenreMoviesWithSubtotals)
	for _, result := range *movies {
		group := result.Genres
		groupedResults[group] = append(groupedResults[group], result)
	}

	delete(groupedResults, "")

	var response []models.GenreMoviesWithSubtotalsResponse

	for _, groupResults := range groupedResults {
		var temp models.GenreMoviesWithSubtotalsResponse
		for _, result := range groupResults {
			if result.PrimaryTitle == "" {
				temp.Total = result.Subtotal
			} else {
				temp.Record = append(temp.Record, result)
			}
		}

		response = append(response, temp)
	}

	return &response, nil
}

func AddNewMovie(body *models.Movie) error {
	if err := models.AddNewMovie(body); err != nil {
		return err
	}

	return nil
}

func UpdateRunTime() error {
	if err := models.UpdateRunTime(); err != nil {
		return err
	}

	return nil
}
