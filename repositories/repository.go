package repositories

import (
	"context"

	"github.com/fingerprint/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// M is the model struct, V is the validation struct
type repository[M any, V any] interface {
	Get(ctx context.Context, id string) (*M, error)
	Create(ctx context.Context, ent *M) error
	Update(ctx context.Context, id string, ent *M) error
	Upsert(ctx context.Context, ent *M) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, ent *V) (*[]M, error)
}

type repositoryImpl[M any, V any] struct {
	db *gorm.DB
}

func newRepository[M any, V any](db *gorm.DB) *repositoryImpl[M, V] {
	return &repositoryImpl[M, V]{
		db: db,
	}
}
func (r *repositoryImpl[M, V]) Get(ctx context.Context, id string) (*M, error) {
	ent := new(M)
	if err := r.db.Preload(clause.Associations).First(ent, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return ent, nil
}

func (r *repositoryImpl[M, V]) Create(ctx context.Context, ent *M) error {

	if err := r.db.Create(ent).Error; err != nil {
		return err
	}
	return nil
}
func (r *repositoryImpl[M, V]) Update(ctx context.Context, id string, ent *M) error {
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
func (r *repositoryImpl[M, V]) Upsert(ctx context.Context, ent *M) error {
	if err := r.db.Save(ent).Error; err != nil {
		return err
	}
	return nil
}

func (r *repositoryImpl[M, V]) Delete(ctx context.Context, id string) error {
	ent := new(M)

	if err := r.db.First(ent, "id = ?", id).Error; err != nil {
		return err
	}

	if err := r.db.Unscoped().Delete(ent, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *repositoryImpl[M, V]) Search(ctx context.Context, ent *V) (*[]M, error) {

	ents := &[]M{}

	// Convert payload into map
	entMap, err := utils.TypeConverter[map[string]interface{}](*ent)
	if err != nil {
		return nil, err
	}

	// I follow this example: https://stackoverflow.com/a/42849112
	// entJson, err := json.Marshal(ent)
	// if err != nil {
	// 	return nil, err
	// }
	// var entMap map[string]interface{}
	// json.Unmarshal(entJson, &entMap)
	// if entMap == nil {
	// 	return nil, errors.New("invalid json")
	// }

	if err := r.db.Find(ents, *entMap).Error; err != nil {
		return nil, err
	}

	return ents, nil
}
