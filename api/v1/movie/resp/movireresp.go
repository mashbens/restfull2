package resp

import (
	"time"

	"github.com/mashbens/restfull2/business/movie"
)

type MovieResp struct {
	Id          int
	Title       string
	Description string
	Rating      float64
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func FromService(t movie.Movie) MovieResp {
	return MovieResp{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		Rating:      t.Rating,
		Image:       t.Image,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func FromeServiceSlice(data []movie.Movie) []MovieResp {
	var movieArray []MovieResp
	for key := range data {
		movieArray = append(movieArray, FromService(data[key]))
	}
	return movieArray
}
