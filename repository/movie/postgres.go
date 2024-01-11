package movie

import (
	"github.com/mashbens/restfull2/business/movie"
	"gorm.io/gorm"
)

type MovieRepo struct {
	db *gorm.DB
}

func NewMovieRepo(db *gorm.DB) movie.MovieRepo {
	return &MovieRepo{
		db: db,
	}
}

func (r *MovieRepo) FindAll() (data []movie.Movie, err error) {
	record := []Movie{}
	res := r.db.Find(&record).Debug()
	if res.Error != nil {
		return nil, res.Error
	}
	return toServiceList(record), nil
}

func (r *MovieRepo) InsertMovie(data movie.Movie) (movie.Movie, error) {
	record := fromService(data)
	res := r.db.Create(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}
func (r *MovieRepo) FindMovieById(id int) (movie.Movie, error) {
	var record Movie
	res := r.db.Where("id = ?", id).First(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}
func (r *MovieRepo) UpdateMovieById(id int, data movie.Movie) (movie.Movie, error) {
	record := fromService(data)
	res := r.db.Where("id = ?", id).Updates(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	res = r.db.Where("id = ?", id).Find(&record)
	if res.Error != nil {
		return record.toService(), res.Error
	}
	return record.toService(), nil
}
func (r *MovieRepo) DeleteMovieById(id int) error {
	var record Movie
	res := r.db.Where("id = ?", id).Delete(&record)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
