package favorite

import "time"

type FavoriteModel struct {
	ID uint64 `gorm:"primaryKey"`

	UserID uint64 `gorm:"index:idx_favorites_user_id;index:idx_favorites_user_game,unique"`
	GameID uint64 `gorm:"index:idx_favorites_game_id;index:idx_favorites_user_game,unique"`

	CreatedAt time.Time
}

func (FavoriteModel) TableName() string {
	return "favorites"
}
