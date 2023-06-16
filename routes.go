package routes

import (
	"project/controller"

	"github.com/gin-gonic/gin"
)

func Handler(router *gin.Engine) {
	v1Version := router.Group("/api/v1")

	v1Version.GET("/longest-duration-movies", controller.FetchLongestDurationMovies)
	v1Version.GET("/top-rated-movies", controller.FetchTopRatedMovies)
	v1Version.GET("/genre-movies-with-subtotals", controller.FetchGenreMoviesWithSubtotals)
	v1Version.POST("/new-movie", controller.AddNewMovie)
	v1Version.POST("/update-runtime-minutes", controller.UpdateRunTime)
}
