package repositories

import (
	"gorm.io/gorm"
)

type repository[M any] interface {
	Get(id string) (*M, error)
	Create(ent *M) error
	Update(id string, ent *M) error
	Upsert(ent *M) error
	Delete(id string) error
}

type repositoryImpl[M any] struct {
	db *gorm.DB
}

func newRepository[M any](db *gorm.DB) *repositoryImpl[M] {
	return &repositoryImpl[M]{
		db: db,
	}
}
func (r *repositoryImpl[M]) Get(id string) (*M, error) {
	ent := new(M)
	if err := r.db.First(ent, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return ent, nil
}

func (r *repositoryImpl[M]) Create(ent *M) error {

	if err := r.db.Create(ent).Error; err != nil {
		return err
	}
	return nil
}
func (r *repositoryImpl[M]) Update(id string, ent *M) error {
	if err := r.db.Model(ent).Where("id = ?", id).Updates(ent).Error; err != nil {
		return err
	}
	return nil
}
func (r *repositoryImpl[M]) Upsert(ent *M) error {
	if err := r.db.Save(ent).Error; err != nil {
		return err
	}
	return nil
}

func (r *repositoryImpl[M]) Delete(id string) error {
	ent := new(M)
	if err := r.db.Delete(ent, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
