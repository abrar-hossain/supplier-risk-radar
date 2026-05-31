package config

import "os"

type Config struct {
	CompaniesHouseAPIKey string
	ServerPort           string
}

func Load() Config {
	return Config{
		CompaniesHouseAPIKey: os.Getenv("COMPANIES_HOUSE_API_KEY"),
		ServerPort:           os.Getenv("SERVER_PORT"),
	}
}
