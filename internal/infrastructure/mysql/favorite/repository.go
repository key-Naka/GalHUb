package favorite

import (
	"context"
	"errors"
	domain "galhub/internal/domain/favorite"

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
	favorite *domain.Favorite,
) error {
	model := ToModel(favorite)

	return r.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(model).Error; err != nil {
				return err
			}

			if err := incrementGameFavoriteCount(
				tx,
				favorite.GameID,
			); err != nil {
				return err
			}

			favorite.ID = model.ID
			favorite.CreatedAt = model.CreatedAt

			return nil
		})
}

func (r *Repository) GetByIDAndUserID(
	ctx context.Context,
	id uint64,
	userID uint64,
) (*domain.Favorite, error) {
	var model FavoriteModel

	err := r.db.
		WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
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

func (r *Repository) GetByUserIDAndGameID(
	ctx context.Context,
	userID uint64,
	gameID uint64,
) (*domain.Favorite, error) {
	var model FavoriteModel

	err := r.db.
		WithContext(ctx).
		Where("user_id = ? AND game_id = ?", userID, gameID).
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
	favorite *domain.Favorite,
) error {
	return r.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			var current FavoriteModel

			if err := tx.
				Where("id = ? AND user_id = ?", favorite.ID, favorite.UserID).
				First(&current).
				Error; err != nil {
				return err
			}

			if err := tx.
				Model(&FavoriteModel{}).
				Where("id = ? AND user_id = ?", favorite.ID, favorite.UserID).
				Update("game_id", favorite.GameID).
				Error; err != nil {
				return err
			}

			if current.GameID == favorite.GameID {
				return nil
			}

			if err := decrementGameFavoriteCount(
				tx,
				current.GameID,
			); err != nil {
				return err
			}

			if err := incrementGameFavoriteCount(
				tx,
				favorite.GameID,
			); err != nil {
				return err
			}

			return nil
		})
}

func (r *Repository) Delete(
	ctx context.Context,
	id uint64,
	userID uint64,
) error {
	return r.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			var current FavoriteModel

			if err := tx.
				Where("id = ? AND user_id = ?", id, userID).
				First(&current).
				Error; err != nil {
				return err
			}

			if err := tx.
				Where("id = ? AND user_id = ?", id, userID).
				Delete(&FavoriteModel{}).
				Error; err != nil {
				return err
			}

			if err := decrementGameFavoriteCount(
				tx,
				current.GameID,
			); err != nil {
				return err
			}

			return nil
		})
}

func (r *Repository) List(
	ctx context.Context,
	query domain.ListQuery,
) ([]*domain.Favorite, int64, error) {
	var models []FavoriteModel
	var total int64

	countDB := r.applyListFilters(
		r.db.
			WithContext(ctx).
			Model(&FavoriteModel{}),
		query,
	)
	if err := countDB.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.applyListFilters(
		r.db.
			WithContext(ctx).
			Model(&FavoriteModel{}),
		query,
	).
		Offset(query.Offset).
		Limit(query.Limit).
		Order("favorites.id desc").
		Find(&models).
		Error
	if err != nil {
		return nil, 0, err
	}

	result := make([]*domain.Favorite, 0, len(models))

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
	db = db.Where("favorites.user_id = ?", query.UserID)

	if query.Keyword != "" {
		keyword := "%" + query.Keyword + "%"

		db = db.
			Joins("INNER JOIN games ON games.id = favorites.game_id").
			Where(
				"games.title LIKE ? OR games.original_title LIKE ?",
				keyword,
				keyword,
			)
	}

	return db
}

func incrementGameFavoriteCount(
	tx *gorm.DB,
	gameID uint64,
) error {
	return tx.
		Table("games").
		Where("id = ?", gameID).
		UpdateColumn(
			"favorite_count",
			gorm.Expr("favorite_count + ?", 1),
		).
		Error
}

func decrementGameFavoriteCount(
	tx *gorm.DB,
	gameID uint64,
) error {
	return tx.
		Table("games").
		Where("id = ?", gameID).
		UpdateColumn(
			"favorite_count",
			gorm.Expr(
				"CASE WHEN favorite_count > 0 THEN favorite_count - 1 ELSE 0 END",
			),
		).
		Error
}
