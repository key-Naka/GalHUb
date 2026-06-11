package game_tag

type GameTagModel struct {
	GameID uint64 `gorm:"primaryKey"`
	TagID  uint64 `gorm:"primaryKey"`
}

func (GameTagModel) TableName() string {
	return "game_tags"
}
