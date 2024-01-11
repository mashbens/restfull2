package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mashbens/restfull2/api/v1/movie"
)

type Controller struct {
	Movie *movie.MovieController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	movieRoutes := e.Group("movie")
	movieRoutes.GET("", controller.Movie.FindAll)
	movieRoutes.GET("/:id", controller.Movie.FindMovieById)
	movieRoutes.POST("", controller.Movie.InsertMovie)
	movieRoutes.PUT("/:id", controller.Movie.UpdateMovieById)
	movieRoutes.DELETE("/:id", controller.Movie.DeleteMovieById)

}
