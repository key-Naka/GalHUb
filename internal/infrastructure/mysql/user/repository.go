package user

import (
	"context"
	"errors"
	domain "galhub/internal/domain/user"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.Repository {
	return &Repository{db: db}
}
func (r *Repository) Create(ctx context.Context, user *domain.User) error {
	model := ToModel(user)
	return r.db.WithContext(ctx).Create(model).Error
}
func (r *Repository) GetByID(ctx context.Context, id uint64) (*domain.User, error) {

	var model UserModel

	err := r.db.
		WithContext(ctx).
		First(&model, id).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return ToDomain(&model), nil
}
func (r *Repository) GetByEmail(
	ctx context.Context,
	email string,
) (*domain.User, error) {

	var model UserModel

	err := r.db.
		WithContext(ctx).
		Where("email = ?", email).
		First(&model).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return ToDomain(&model), nil
}
func (r *Repository) GetByUsername(
	ctx context.Context,
	username string,
) (*domain.User, error) {

	var model UserModel

	err := r.db.
		WithContext(ctx).
		Where("username = ?", username).
		First(&model).
		Error

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
	user *domain.User,
) error {

	return r.db.
		WithContext(ctx).
		Save(ToModel(user)).
		Error
}
func (r *Repository) Delete(
	ctx context.Context,
	id uint64,
) error {

	return r.db.
		WithContext(ctx).
		Delete(&UserModel{}, id).
		Error
}
