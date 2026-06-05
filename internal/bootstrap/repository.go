package bootstrap

import (
	mysqlDB "galhub/internal/infrastructure/mysql"

	mysqlUser "galhub/internal/infrastructure/mysql/user"

	domainUser "galhub/internal/domain/user"
)

type Repositories struct {
	User domainUser.Repository
}

func NewRepositories() *Repositories {

	return &Repositories{
		User: mysqlUser.NewRepository(
			mysqlDB.DB,
		),
	}
}
