package config

type Config struct {
	ApiPort string
	Database Database
}

type Database struct {
	DbConnection string
	Mysql Mysql
}

type Mysql struct {
	Host     string
	Database string
	User     string
	Password string
	Port     string
}

var defaultConfig = Config{
	ApiPort: "80",
	Database: Database{
		DbConnection: "Mysql",
		Mysql: Mysql{
			Host:     "localhost",
			Database: "event",
			User:     "user",
			Password: "password",
			Port:     "3306",
		},
	},
}
