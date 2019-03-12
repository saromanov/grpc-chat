package config

// Config defines configuration for service
type Config struct {
	DBName     string `yaml:"dbname"`
	DBPassword string `yaml:"dbpassword"`
	DBUser     string `yaml:"dbuser"`
	Port       int    `yaml:"port"`
}
