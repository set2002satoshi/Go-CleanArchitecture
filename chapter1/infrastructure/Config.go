package infrastructure

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
		Test struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading: %w", err)
	}
	
	c := new(Config)

	c.DB.Production.Host     = os.Getenv("DB_HOST")
	c.DB.Production.Username = os.Getenv("DB_USER")
	c.DB.Production.Password = os.Getenv("DB_PASSWORD")
	c.DB.Production.DBName   = os.Getenv("DB_NAME")

	c.DB.Test.Host     = os.Getenv("DB_HOST")
	c.DB.Test.Username = os.Getenv("DB_USER")
	c.DB.Test.Password = os.Getenv("DB_PASSWORD")
	c.DB.Test.DBName   = os.Getenv("DB_NAME")

	c.Routing.Port = ":8080"

	return c
}
