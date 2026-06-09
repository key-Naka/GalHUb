package bootstrap

import (
	mysqlDB "galhub/internal/infrastructure/mysql"

	mysqlGame "galhub/internal/infrastructure/mysql/game"
	mysqlUser "galhub/internal/infrastructure/mysql/user"

	domainGame "galhub/internal/domain/game"
	domainUser "galhub/internal/domain/user"
)

type Repositories struct {
	User domainUser.Repository
	Game domainGame.Repository
}

func NewRepositories() *Repositories {

	return &Repositories{
		User: mysqlUser.NewRepository(
			mysqlDB.DB,
		),
		Game: mysqlGame.NewRepository(
			mysqlDB.DB,
		),
	}
}
