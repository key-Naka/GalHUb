package bootstrap

import (
	mysqlDB "galhub/internal/infrastructure/mysql"

	mysqlGame "galhub/internal/infrastructure/mysql/game"
	mysqlGameTag "galhub/internal/infrastructure/mysql/game_tag"
	mysqlTag "galhub/internal/infrastructure/mysql/tag"
	mysqlUser "galhub/internal/infrastructure/mysql/user"

	domainGame "galhub/internal/domain/game"
	domainGameTag "galhub/internal/domain/game_tage"
	domainTag "galhub/internal/domain/tag"
	domainUser "galhub/internal/domain/user"
)

type Repositories struct {
	User    domainUser.Repository
	Game    domainGame.Repository
	Tag     domainTag.Repository
	GameTag domainGameTag.Repository
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
	}
}
