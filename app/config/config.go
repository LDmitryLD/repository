package config

import (
	"os"
)

type AppConf struct {
	DB DB `yaml:"db"`
}

type DB struct {
	Driver   string
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

func NewAppConf() AppConf {

	// if err := godotenv.Load("../.env"); err != nil {
	// 	log.Fatal("ошибка при загрузке переменных окружения: ", err.Error())
	// }

	return AppConf{
		DB: DB{
			Driver:   os.Getenv("DB_DRIVER"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
	}
}
