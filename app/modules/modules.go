package modules

import (
	"github.com/mashbens/restfull2/api"
	movie "github.com/mashbens/restfull2/api/v1/movie"
	movieservice "github.com/mashbens/restfull2/business/movie"
	"github.com/mashbens/restfull2/config"
	movierepo "github.com/mashbens/restfull2/repository/movie"
	"github.com/mashbens/restfull2/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	movieRepo := movierepo.RepositoryFactory(dbCon)

	movieService := movieservice.NewMovieService(movieRepo)
	controller := api.Controller{
		Movie: movie.NewMovieController(movieService),
	}
	return controller
}
