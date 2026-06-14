package favorite

import domain "galhub/internal/domain/favorite"

func ToDomain(
	m *FavoriteModel,
) *domain.Favorite {
	return &domain.Favorite{
		ID: m.ID,

		UserID: m.UserID,
		GameID: m.GameID,

		CreatedAt: m.CreatedAt,
	}
}

func ToModel(
	f *domain.Favorite,
) *FavoriteModel {
	return &FavoriteModel{
		ID: f.ID,

		UserID: f.UserID,
		GameID: f.GameID,

		CreatedAt: f.CreatedAt,
	}
}
