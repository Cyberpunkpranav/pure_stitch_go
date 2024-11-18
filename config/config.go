package config

type Config struct {
	Port string
	Env  string
}
type Payload struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    any    `json:"data"`
}

func NewConfig() *Config {
	return &Config{
		Port: "8080",        // Default port
		Env:  "development", // Default environment
	}
}

type Database struct {
	Username   string
	Database   string
	Password   string
	Protocol   string
	Port       string
	Ip_address string
}
type Database_Redis struct {
	Username string
	Password string
	Address  string
}

func Db_Config() Database {
	// var db Database
	// db.Username = "avnadmin"
	// db.Database = "pure_stitch_dev"
	// db.Password = "AVNS_XGcJHSv7MOBNThblLe2"
	// db.Protocol = "tcp"
	// db.Ip_address = "mysql-3c770044-pranavsharma733902-da3e.h.aivencloud.com"
	// db.Port = "17898"
	// return db
}
