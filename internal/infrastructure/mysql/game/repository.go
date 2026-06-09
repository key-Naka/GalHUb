package game

import (
	"context"
	"errors"
	domain "galhub/internal/domain/game"

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
	game *domain.Game,
) error {

	model := ToModel(game)

	return r.db.
		WithContext(ctx).
		Create(model).
		Error
}
func (r *Repository) GetByID(
	ctx context.Context,
	id uint64,
) (*domain.Game, error) {
	var model GameModel
	err := r.db.WithContext(ctx).First(&model, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return ToDomain(&model), nil
}
func (r *Repository) Update(
	ctx context.Context,
	game *domain.Game,
) error {

	return r.db.
		WithContext(ctx).
		Model(&GameModel{}).
		Where("id = ?", game.ID).
		Updates(map[string]interface{}{
			"title":          game.Title,
			"original_title": game.OriginalTitle,
			"cover":          game.Cover,
			"description":    game.Description,
			"release_date":   game.ReleaseDate,
			"status":         game.Status,
		}).
		Error
}
func (r *Repository) Delete(
	ctx context.Context,
	id uint64,
) error {

	return r.db.
		WithContext(ctx).
		Delete(&GameModel{}, id).
		Error
}
func (r *Repository) List(
	ctx context.Context,
	offset int,
	limit int,
) ([]*domain.Game, error) {

	var models []GameModel

	err := r.db.
		WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Order("id desc").
		Find(&models).
		Error

	if err != nil {
		return nil, err
	}

	result := make([]*domain.Game, 0, len(models))

	for _, model := range models {
		result = append(
			result,
			ToDomain(&model),
		)
	}

	return result, nil
}
