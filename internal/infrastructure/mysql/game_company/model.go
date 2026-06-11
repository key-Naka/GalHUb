package game_company

type GameCompanyModel struct {
	GameID    uint64 `gorm:"primaryKey;autoIncrement:false"`
	CompanyID uint64 `gorm:"primaryKey;autoIncrement:false"`
}

func (GameCompanyModel) TableName() string {
	return "game_companies"
}
