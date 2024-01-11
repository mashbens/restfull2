package request

import "github.com/mashbens/restfull2/business/movie"

type MovieRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float64 `json:"rating" validate:"required"`
	Image       string  `json:"image"`
}

func ToService(t MovieRequest) movie.Movie {
	return movie.Movie{
		Title:       t.Title,
		Description: t.Description,
		Rating:      t.Rating,
		Image:       t.Image,
	}
}
