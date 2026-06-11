package game_tag

import (
	"context"
	gametag "galhub/internal/domain/game_tage"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) gametag.Repository {

	return &Repository{
		db: db,
	}
}

func (r *Repository) ReplaceGameTags(
	ctx context.Context,
	gameID uint64,
	tagIDs []uint64,
) error {

	return r.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("game_id = ?", gameID).Delete(&GameTagModel{}).Error; err != nil {
				return err
			}

			if len(tagIDs) == 0 {
				return nil
			}

			var models []GameTagModel

			for _, tagID := range tagIDs {

				models = append(
					models,
					GameTagModel{
						GameID: gameID,
						TagID:  tagID,
					},
				)
			}

			return tx.Create(
				&models,
			).Error
		})
}

func (r *Repository) GetTagIDsByGameID(
	ctx context.Context,
	gameID uint64,
) ([]uint64, error) {

	var tagIDs []uint64

	err := r.db.
		WithContext(ctx).
		Model(&GameTagModel{}).
		Where("game_id = ?", gameID).
		Pluck("tag_id", &tagIDs).
		Error

	if err != nil {
		return nil, err
	}

	return tagIDs, nil
}
