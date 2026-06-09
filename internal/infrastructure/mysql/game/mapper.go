package game

import domain "galhub/internal/domain/game"

func ToDomain(
	m *GameModel,
) *domain.Game {

	return &domain.Game{
		ID: m.ID,

		Title:         m.Title,
		OriginalTitle: m.OriginalTitle,

		Cover:       m.Cover,
		Description: m.Description,

		ReleaseDate: m.ReleaseDate,

		ViewCount:     m.ViewCount,
		FavoriteCount: m.FavoriteCount,

		Status: m.Status,

		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
func ToModel(
	g *domain.Game,
) *GameModel {

	return &GameModel{
		ID: g.ID,

		Title:         g.Title,
		OriginalTitle: g.OriginalTitle,

		Cover:       g.Cover,
		Description: g.Description,

		ReleaseDate: g.ReleaseDate,

		ViewCount:     g.ViewCount,
		FavoriteCount: g.FavoriteCount,

		Status: g.Status,
	}
}
