package controller

import (
	models "project/model"
	"project/service"

	"github.com/gin-gonic/gin"
)

func FetchLongestDurationMovies(c *gin.Context) {
	movies, err := service.GetLongestDurationMovies()
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, movies)
}

func FetchTopRatedMovies(c *gin.Context) {
	movies, err := service.GetTopRatedMovies()
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, movies)
}

func FetchGenreMoviesWithSubtotals(c *gin.Context) {
	movies, err := service.GetGenreMoviesWithSubtotals()
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, movies)
}

func AddNewMovie(c *gin.Context) {
	var body models.Movie
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, "invalid body")
		return
	}

	if err := service.AddNewMovie(&body); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "Successfully added new movie")
}

func UpdateRunTime(c *gin.Context) {
	if err := service.UpdateRunTime(); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "Successfully updated run time for all movies")
}
