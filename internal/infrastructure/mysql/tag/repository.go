package tag

import (
	"context"
	"errors"
	domain "galhub/internal/domain/tag"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) domain.Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(
	ctx context.Context,
	tag *domain.Tag,
) error {

	model := ToModel(tag)

	return r.db.
		WithContext(ctx).
		Create(model).
		Error
}

func (r *Repository) GetByID(
	ctx context.Context,
	id uint64,
) (*domain.Tag, error) {
	var model TagModel
	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return ToDomain(&model), nil
}

func (r *Repository) Update(
	ctx context.Context,
	tag *domain.Tag,
) error {

	return r.db.
		WithContext(ctx).
		Model(&TagModel{}).
		Where("id = ?", tag.ID).
		Updates(map[string]interface{}{
			"name": tag.Name,
		}).
		Error
}

func (r *Repository) Delete(
	ctx context.Context,
	id uint64,
) error {

	return r.db.
		WithContext(ctx).
		Delete(&TagModel{}, id).
		Error
}

func (r *Repository) List(
	ctx context.Context,
) ([]*domain.Tag, error) {

	var models []TagModel

	err := r.db.
		WithContext(ctx).
		Order("id desc").
		Find(&models).
		Error

	if err != nil {
		return nil, err
	}

	result := make([]*domain.Tag, 0, len(models))

	for _, model := range models {
		result = append(
			result,
			ToDomain(&model),
		)
	}

	return result, nil
}
