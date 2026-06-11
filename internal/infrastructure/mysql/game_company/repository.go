package game_company

import (
	"context"
	gamecompany "galhub/internal/domain/game_company"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) gamecompany.Repository {

	return &Repository{
		db: db,
	}
}

func (r *Repository) ReplaceGameCompanies(
	ctx context.Context,
	gameID uint64,
	companyIDs []uint64,
) error {

	return r.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {

			if err := tx.
				Where("game_id = ?", gameID).
				Delete(&GameCompanyModel{}).
				Error; err != nil {

				return err
			}

			if len(companyIDs) == 0 {
				return nil
			}

			var models []GameCompanyModel

			for _, companyID := range companyIDs {

				models = append(
					models,
					GameCompanyModel{
						GameID:    gameID,
						CompanyID: companyID,
					},
				)
			}

			return tx.Create(
				&models,
			).Error
		})
}

func (r *Repository) GetCompanyIDsByGameID(
	ctx context.Context,
	gameID uint64,
) ([]uint64, error) {

	var companyIDs []uint64

	err := r.db.
		WithContext(ctx).
		Model(&GameCompanyModel{}).
		Where("game_id = ?", gameID).
		Pluck("company_id", &companyIDs).
		Error

	if err != nil {
		return nil, err
	}

	return companyIDs, nil
}
