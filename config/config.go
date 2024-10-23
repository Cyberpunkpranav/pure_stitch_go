package config

type Config struct{
	Port string
	Env string
}
func NewConfig() *Config {
    return &Config{
        Port: "8080",        // Default port
        Env:  "development", // Default environment
    }
}