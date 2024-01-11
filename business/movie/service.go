package movie

import (
	"errors"
)

type MovieRepo interface {
	FindAll() ([]Movie, error)
	InsertMovie(Movie) (Movie, error)
	FindMovieById(id int) (Movie, error)
	UpdateMovieById(int, Movie) (Movie, error)
	DeleteMovieById(id int) error
}

type MovieService interface {
	FindAll() ([]Movie, error)
	InsertMovie(Movie) (Movie, error)
	FindMovieById(id int) (Movie, error)
	UpdateMovieById(int, Movie) (Movie, error)
	DeleteMovieById(id int) error
}

type movieService struct {
	movieRepo MovieRepo
}

func NewMovieService(
	movieRepo MovieRepo,
) MovieService {
	return &movieService{
		movieRepo: movieRepo}
}

func (c *movieService) FindAll() (res []Movie, err error) {
	t, err := c.movieRepo.FindAll()
	if err != nil {
		return res, err
	}
	return t, nil
}

func (c *movieService) InsertMovie(movie Movie) (res Movie, err error) {
	t, err := c.movieRepo.InsertMovie(movie)
	if err != nil {
		return res, err
	}
	return t, nil
}

func (c *movieService) FindMovieById(id int) (res Movie, err error) {
	t, err := c.movieRepo.FindMovieById(id)
	if err != nil {
		return res, err
	}
	return t, nil
}
func (c *movieService) UpdateMovieById(id int, movie Movie) (res Movie, err error) {
	t, _ := c.movieRepo.FindMovieById(id)
	if t.Id == 0 {
		return res, errors.New("id not found")
	}
	update, err := c.movieRepo.UpdateMovieById(id, movie)
	if err != nil {
		return res, err
	}
	update.Id = id
	return update, nil
}

func (c *movieService) DeleteMovieById(id int) error {
	t, _ := c.movieRepo.FindMovieById(id)
	if t.Id == 0 {
		return errors.New("id not found")
	}
	err := c.movieRepo.DeleteMovieById(t.Id)
	if err != nil {
		return errors.New("error deleting")
	}
	return nil
}
