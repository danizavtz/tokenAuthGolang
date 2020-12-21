package services

import (
	"context"
	"sync"
	"github.com/jackc/pgx/v4/pgxpool"
	"fmt"
	"os"
)
var once sync.Once
var dbInstance *pgxpool.Pool

func GetDatabaseInstance()(*pgxpool.Pool) {
	once.Do(func (){
		dbi, err := pgxpool.Connect(context.Background(), assembleConnectionStringForDatabase())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		dbInstance = dbi
	})
	return dbInstance
}

func assembleConnectionStringForDatabase() string {
	connectionStringDatabase := "postgres://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_ADDRESS")+":"+os.Getenv("DB_PORT")+"/"+os.Getenv("DB_NAME")+"?sslmode=disable"
	return connectionStringDatabase
}