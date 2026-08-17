package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewConnection(cfg *dbConfig) (*sql.DB,error){
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName,
	)
	fmt.Println(dsn)
	db,err := sql.Open("postgres",dsn)
	if err != nil{
		return nil, err
	}

	return db,nil
}