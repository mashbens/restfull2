package movie

import "time"

type Movie struct {
	Id          int
	Title       string
	Description string
	Rating      float64
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
