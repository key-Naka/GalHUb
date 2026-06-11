package bootstrap

import (
	mysqlDB "galhub/internal/infrastructure/mysql"

	mysqlCompany "galhub/internal/infrastructure/mysql/company"
	mysqlGame "galhub/internal/infrastructure/mysql/game"
	mysqlGameCompany "galhub/internal/infrastructure/mysql/game_company"
	mysqlGameTag "galhub/internal/infrastructure/mysql/game_tag"
	mysqlTag "galhub/internal/infrastructure/mysql/tag"
	mysqlUser "galhub/internal/infrastructure/mysql/user"

	domainCompany "galhub/internal/domain/company"
	domainGame "galhub/internal/domain/game"
	domainGameCompany "galhub/internal/domain/game_company"
	domainGameTag "galhub/internal/domain/game_tage"
	domainTag "galhub/internal/domain/tag"
	domainUser "galhub/internal/domain/user"
)

type Repositories struct {
	User        domainUser.Repository
	Game        domainGame.Repository
	Tag         domainTag.Repository
	GameTag     domainGameTag.Repository
	Company     domainCompany.Repository
	GameCompany domainGameCompany.Repository
}

func NewRepositories() *Repositories {

	return &Repositories{
		User: mysqlUser.NewRepository(
			mysqlDB.DB,
		),
		Game: mysqlGame.NewRepository(
			mysqlDB.DB,
		),
		Tag: mysqlTag.NewRepository(
			mysqlDB.DB,
		),
		GameTag: mysqlGameTag.NewRepository(
			mysqlDB.DB,
		),
		Company: mysqlCompany.NewRepository(
			mysqlDB.DB,
		),
		GameCompany: mysqlGameCompany.NewRepository(
			mysqlDB.DB,
		),
	}
}
