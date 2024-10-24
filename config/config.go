package config

type Config struct {
	Port string
	Env  string
}
type Payload struct{
	Message string `json:"message"`
	Status string `json:"status"`
	Data string `json:"data"`
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
	var db Database
	db.Username = "root"
	db.Database = "pure_stitch_dev"
	db.Password = "admin@123"
	db.Protocol = "tcp"
	db.Ip_address = "127.0.0.1"
	db.Port = "3306"
	return db
}
