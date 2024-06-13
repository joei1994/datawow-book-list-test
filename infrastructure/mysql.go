package infrastructure

import (
	"fmt"

	"datawow/book-list/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySql(config *config.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MySql.User,
		config.MySql.Password,
		config.MySql.Host,
		config.MySql.Port,
		config.MySql.Name)

	fmt.Println(dataSourceName)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
