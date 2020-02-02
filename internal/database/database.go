package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mohsentm/event-project-test/config"
)

var db *gorm.DB

func init() {
	config.Init()
}

func connection() (*gorm.DB) {
	config.Get()
	conf := config.Get()
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

	return db
}

func mysqlConnection(mysql *config.Mysql) (db *gorm.DB, err error) {
	return gorm.Open("mysql", mysql.User+":"+mysql.Password+"@("+mysql.Host+":"+mysql.Port+")/"+mysql.Database+"?parseTime=true")
}
