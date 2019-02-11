package config

type DbConfig struct {
	DbType string `json:"DB_TYPE"`
	Name string `json:"DB_NAME"`
	Port int	`json:"DB_PORT"`
	User string `json:"DB_USER"`
	Pass string `json:"DB_PASS"`
}
