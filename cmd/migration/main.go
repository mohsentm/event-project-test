package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mohsentm/event-project-test/config"
	"github.com/mohsentm/event-project-test/internal/event/model"
)

func init() {
	config.Init()
}

func main() {
	conf := config.Get()
	var db *gorm.DB
	var err error

	switch conf.Database.DbConnection {
	case "mysql":
		db, err = mysqlConnection(&conf.Database.Mysql)
		if err != nil {
			panic(err)
		}
	default:
		panic("undefined database default connection")
	}

	db.AutoMigrate(&model.Event{})

	defer db.Close()

}

func mysqlConnection(mysql *config.Mysql) (db *gorm.DB, err error) {
	return gorm.Open("mysql", mysql.User+":"+mysql.Password+"@("+mysql.Host+":"+mysql.Port+")/"+mysql.Database)
}
