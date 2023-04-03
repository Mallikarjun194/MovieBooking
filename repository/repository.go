package repository

import (
	"MovieBooking/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RepositoryI interface {
	Create(value any) error
	QueryAll(value any) error
	Query(value any) error
	Update(value any) error
	Delete(value any) error
	QueryField(value any, field string, fvalue string) error
}

type Repository struct {
	Db *gorm.DB
}

func OpenDBConnection() Repository {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(new(models.Movie), new(models.Theater), new(models.Seat), new(models.Show), new(models.Ticket))

	return Repository{Db: db}
}

func (r *Repository) Create(value any) error {
	return r.Db.Create(value).Error
}
func (r *Repository) QueryAll(value any) error {
	return r.Db.Find(value).Error
}

func (r *Repository) Query(value any) error {
	return r.Db.First(value).Error
}

func (r *Repository) QueryField(value any, field string, fvalue string) error {
	r.Db.Where(fmt.Sprintf("%v == '%v'", field, fvalue)).Find(value)
	return nil
}

func (r *Repository) Update(value any) error {
	return r.Db.Updates(value).Error
}

func (r *Repository) Delete(value any) error {
	return r.Db.Delete(value).Error
}

//
//func (r *Repository) Query(value any, field string, fvalue string) error {
//
//}
