package company

import (
	"context"
	"errors"
	domain "galhub/internal/domain/company"

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
	company *domain.Company,
) error {

	model := ToModel(company)

	return r.db.
		WithContext(ctx).
		Create(model).
		Error
}

func (r *Repository) GetByID(
	ctx context.Context,
	id uint64,
) (*domain.Company, error) {
	var model CompanyModel
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
	company *domain.Company,
) error {

	return r.db.
		WithContext(ctx).
		Model(&CompanyModel{}).
		Where("id = ?", company.ID).
		Updates(map[string]interface{}{
			"name": company.Name,
		}).
		Error
}

func (r *Repository) Delete(
	ctx context.Context,
	id uint64,
) error {

	return r.db.
		WithContext(ctx).
		Delete(&CompanyModel{}, id).
		Error
}

func (r *Repository) List(
	ctx context.Context,
) ([]*domain.Company, error) {

	var models []CompanyModel

	err := r.db.
		WithContext(ctx).
		Order("id desc").
		Find(&models).
		Error

	if err != nil {
		return nil, err
	}

	result := make([]*domain.Company, 0, len(models))

	for _, model := range models {
		result = append(
			result,
			ToDomain(&model),
		)
	}

	return result, nil
}
