// ! config loader/defaults
package config

type Config struct {
	Host string
	Port string
}

func LoadConfig() *Config{
	return &Config{
		Host: "0.0.0.0",
		Port: "6379",
	}
}