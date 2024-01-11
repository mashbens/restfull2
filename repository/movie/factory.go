package movie

import (
	"github.com/mashbens/restfull2/business/movie"
	"github.com/mashbens/restfull2/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) movie.MovieRepo {
	var movieRepository movie.MovieRepo

	if dbCon.Driver == util.POSTGRES {
		movieRepository = NewMovieRepo(dbCon.POSTGRES)
		dbCon.POSTGRES.AutoMigrate(&Movie{})
	} else {
		panic("Database driver not supported")
	}
	return movieRepository
}
