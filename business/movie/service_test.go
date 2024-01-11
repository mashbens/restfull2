package movie_test

import (
	"errors"
	"testing"

	"github.com/mashbens/restfull2/business/movie"
	"github.com/stretchr/testify/assert" // Import testify for assertions
)

// Mock MovieService for testing
type MockMovieService struct {
	FindAllFunc         func() ([]movie.Movie, error)
	InsertMovieFunc     func(movie.Movie) (movie.Movie, error)
	FindMovieByIdFunc   func(int) (movie.Movie, error)
	UpdateMovieByIdFunc func(int, movie.Movie) (movie.Movie, error)
	DeleteMovieByIdFunc func(int) error
}

func (m *MockMovieService) FindAll() ([]movie.Movie, error) {
	return m.FindAllFunc()
}

func (m *MockMovieService) InsertMovie(movie movie.Movie) (movie.Movie, error) {
	return m.InsertMovieFunc(movie)
}

func (m *MockMovieService) FindMovieById(id int) (movie.Movie, error) {
	return m.FindMovieByIdFunc(id)
}

func (m *MockMovieService) UpdateMovieById(id int, movie movie.Movie) (movie.Movie, error) {
	return m.UpdateMovieByIdFunc(id, movie)
}

func (m *MockMovieService) DeleteMovieById(id int) error {
	return m.DeleteMovieByIdFunc(id)
}

func TestFindAll(t *testing.T) {
	mockService := &MockMovieService{
		FindAllFunc: func() ([]movie.Movie, error) {
			return []movie.Movie{{Id: 1, Title: "Movie 1"}, {Id: 2, Title: "Movie 2"}}, nil
		},
	}
	service := movie.NewMovieService(mockService)

	movies, err := service.FindAll()

	assert.NoError(t, err, "Unexpected error")
	assert.Len(t, movies, 2, "Expected 2 movies")
}
func TestInsertMovie(t *testing.T) {
	mockService := &MockMovieService{
		InsertMovieFunc: func(movie movie.Movie) (movie.Movie, error) {
			if movie.Title == "Movie X" {
				return movie, nil // Return the inserted movie if the title is "Movie X"
			}
			return movie, errors.New("failed to insert") // Simulate an error for other titles
		},
	}
	service := movie.NewMovieService(mockService)

	// Test case 1: Successful insertion
	newMovie := movie.Movie{Title: "Movie X"}
	insertedMovie, err := service.InsertMovie(newMovie)
	assert.NoError(t, err, "Unexpected error on successful insertion")
	assert.Equal(t, "Movie X", insertedMovie.Title, "Expected inserted movie title to be 'Movie X'")

	// Test case 2: Failed insertion
	invalidMovie := movie.Movie{Title: "Invalid Movie"}
	_, err = service.InsertMovie(invalidMovie)
	assert.Error(t, err, "Expected an error for failed insertion")
}

func TestFindMovieById(t *testing.T) {
	mockService := &MockMovieService{
		FindMovieByIdFunc: func(id int) (movie.Movie, error) {
			if id == 1 {
				return movie.Movie{Id: 1, Title: "Movie 1"}, nil // Return movie with ID 1
			}
			return movie.Movie{}, errors.New("movie not found")
		},
	}
	service := movie.NewMovieService(mockService)

	// Test case 1: Find existing movie by ID
	foundMovie, err := service.FindMovieById(1)
	assert.NoError(t, err, "Unexpected error while finding existing movie")
	assert.Equal(t, 1, foundMovie.Id, "Expected movie with ID 1")

	// Test case 2: Find non-existing movie by ID
	_, err = service.FindMovieById(2)
	assert.Error(t, err, "Expected error for non-existing movie")
}

func TestUpdateMovieById(t *testing.T) {
	mockService := &MockMovieService{
		FindMovieByIdFunc: func(id int) (movie.Movie, error) {
			if id == 1 {
				return movie.Movie{Id: 1, Title: "Old Title"}, nil // Return movie with ID 1
			}
			return movie.Movie{}, errors.New("movie not found")
		},
		UpdateMovieByIdFunc: func(id int, updatedMovie movie.Movie) (movie.Movie, error) {
			if id == 1 {
				updatedMovie.Id = id
				return updatedMovie, nil // Return updated movie
			}
			return movie.Movie{}, errors.New("failed to update")
		},
	}
	service := movie.NewMovieService(mockService)

	// Test case 1: Update existing movie by ID
	updatedMovie, err := service.UpdateMovieById(1, movie.Movie{Title: "New Title"})
	assert.NoError(t, err, "Unexpected error while updating existing movie")
	assert.Equal(t, "New Title", updatedMovie.Title, "Expected updated movie title to be 'New Title'")

	// Test case 2: Update non-existing movie by ID
	_, err = service.UpdateMovieById(2, movie.Movie{Title: "Another Title"})
	assert.Error(t, err, "Expected error for updating non-existing movie")
}

func TestDeleteMovieById(t *testing.T) {
	mockService := &MockMovieService{
		FindMovieByIdFunc: func(id int) (movie.Movie, error) {
			if id == 1 {
				return movie.Movie{Id: 1, Title: "Existing Movie"}, nil // Return movie with ID 1
			}
			return movie.Movie{}, errors.New("movie not found")
		},
		DeleteMovieByIdFunc: func(id int) error {
			if id == 1 {
				return nil // Successful deletion
			}
			return errors.New("failed to delete")
		},
	}
	service := movie.NewMovieService(mockService)

	// Test case 1: Delete existing movie by ID
	err := service.DeleteMovieById(1)
	assert.NoError(t, err, "Unexpected error while deleting existing movie")

	// Test case 2: Delete non-existing movie by ID
	err = service.DeleteMovieById(2)
	assert.Error(t, err, "Expected error for deleting non-existing movie")
}
