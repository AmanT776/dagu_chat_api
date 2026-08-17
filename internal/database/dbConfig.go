package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type dbConfig struct{
	Host string
	Port int
	User string
	Password string
	DbName string
}

func LoadConfig() *dbConfig{
	godotenv.Load()
	
	port,err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil{
		fmt.Println(err)
	}
	dbValues := dbConfig{
		Host: os.Getenv("DB_HOST"),
		Port: port,
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName: os.Getenv("DB_NAME"),
	}

	return &dbValues
}

