package infrastructure

import (
	"datawow/book-list/config"

	"gorm.io/gorm"
)

type Infrastructure struct {
	MySql *gorm.DB
}

func NewInfrastructure(config *config.Config) Infrastructure {
	return Infrastructure{
		MySql: InitMySql(config),
	}
}
