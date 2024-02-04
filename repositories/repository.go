package repositories

import (
	"github.com/fingerprint/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// M is the model struct, S is the search struct
type repository[M any, S any] interface {
	Get(id string) (*M, error)
	Create(ent *M) error
	Update(id string, ent *M) error
	Upsert(ent *M) error
	Delete(id string) error
	Find(ent *S) (*[]M, error)
}

type repositoryImpl[M any, S any] struct {
	db *gorm.DB
}

func newRepository[M any, S any](db *gorm.DB) *repositoryImpl[M, S] {
	return &repositoryImpl[M, S]{
		db: db,
	}
}
func (r *repositoryImpl[M, S]) Get(id string) (*M, error) {
	ent := new(M)
	if err := r.db.First(ent, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return ent, nil
}

func (r *repositoryImpl[M, S]) Create(ent *M) error {

	if err := r.db.Create(ent).Error; err != nil {
		return err
	}
	return nil
}
func (r *repositoryImpl[M, S]) Update(id string, ent *M) error {
	// if err := r.db.Model(ent).Where("id = ?", id).Updates(ent).Error; err != nil {
	// 	return err
	// }

	if err := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Model(ent).Where("id = ?", id).Updates(ent).Error; err != nil {
		return err
	}

	if err := r.db.Preload(clause.Associations).First(ent, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
func (r *repositoryImpl[M, S]) Upsert(ent *M) error {
	if err := r.db.Save(ent).Error; err != nil {
		return err
	}
	return nil
}

func (r *repositoryImpl[M, S]) Delete(id string) error {
	ent := new(M)

	if err := r.db.First(ent, "id = ?", id).Error; err != nil {
		return err
	}

	if err := r.db.Unscoped().Delete(ent, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *repositoryImpl[M, S]) Find(ent *S) (*[]M, error) {

	ents := &[]M{}

	// Convert payload into map
	entMap, err := utils.TypeConverter[map[string]interface{}](*ent)
	if err != nil {
		return nil, err
	}

	if err := r.db.Find(ents, *entMap).Error; err != nil {
		return nil, err
	}

	return ents, nil
}
