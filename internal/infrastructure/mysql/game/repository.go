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

		return nil, err
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
	query domain.ListQuery,
) ([]*domain.Game, int64, error) {

	var models []GameModel
	var total int64

	countDB := r.applyListFilters(
		r.db.
			WithContext(ctx).
			Model(&GameModel{}),
		query,
	)

	if err := countDB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.applyListFilters(
		r.db.
			WithContext(ctx).
			Model(&GameModel{}),
		query,
	).
		Offset(query.Offset).
		Limit(query.Limit).
		Order("id desc").
		Find(&models).
		Error

	if err != nil {
		return nil, 0, err
	}

	result := make([]*domain.Game, 0, len(models))

	for _, model := range models {
		result = append(
			result,
			ToDomain(&model),
		)
	}

	return result, total, nil
}

func (r *Repository) applyListFilters(
	db *gorm.DB,
	query domain.ListQuery,
) *gorm.DB {
	if query.Keyword != "" {
		keyword := "%" + query.Keyword + "%"

		db = db.Where(
			`title LIKE ? OR original_title LIKE ? OR EXISTS (
				SELECT 1
				FROM game_tags gt
				INNER JOIN tags t ON t.id = gt.tag_id
				WHERE gt.game_id = games.id
				AND t.name LIKE ?
			) OR EXISTS (
				SELECT 1
				FROM game_companies gc
				INNER JOIN companies c ON c.id = gc.company_id
				WHERE gc.game_id = games.id
				AND c.name LIKE ?
			)`,
			keyword,
			keyword,
			keyword,
			keyword,
		)
	}

	if query.TagID != nil {
		db = db.Where(
			`EXISTS (
				SELECT 1
				FROM game_tags gt
				WHERE gt.game_id = games.id
				AND gt.tag_id = ?
			)`,
			*query.TagID,
		)
	}

	if query.CompanyID != nil {
		db = db.Where(
			`EXISTS (
				SELECT 1
				FROM game_companies gc
				WHERE gc.game_id = games.id
				AND gc.company_id = ?
			)`,
			*query.CompanyID,
		)
	}

	if query.Status != nil {
		db = db.Where(
			"status = ?",
			*query.Status,
		)
	}

	return db
}
