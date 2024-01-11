package movie

import (
	"time"

	"github.com/mashbens/restfull2/business/movie"
)

type Movie struct {
	Id          int
	Title       string
	Description string
	Rating      float64
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Movie) toService() movie.Movie {
	return movie.Movie{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		Rating:      t.Rating,
		Image:       t.Image,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func fromService(t movie.Movie) Movie {
	return Movie{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		Rating:      t.Rating,
		CreatedAt:   t.CreatedAt,
		Image:       t.Image,
		UpdatedAt:   t.UpdatedAt,
	}
}
func toServiceList(data []Movie) []movie.Movie {
	a := []movie.Movie{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
