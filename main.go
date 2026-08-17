package main

import (
	"database/sql"
	"os"

	"github.com/AmanT776/dagu_chat_api/internal/database"
	"github.com/AmanT776/dagu_chat_api/internal/logger"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	db *sql.DB
}

func main(){
	godotenv.Load()
	logType := os.Getenv("LOG_TYPE")
	logger := logger.New(logType)

	dbConfig := database.LoadConfig()
	db ,err := database.NewConnection(dbConfig)
	if err != nil{
		logger.Error(err.Error())
	}
	apiCfg := apiConfig{
		db: db,
	}

	defer apiCfg.db.Close()
	logger.Info("Database connection successfully established!")
}