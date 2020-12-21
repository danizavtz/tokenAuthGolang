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
		dbi, err := pgxpool.Connect(context.Background(), "postgres://postgres:postgres@0.0.0.0:5432/login?sslmode=disable")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		dbInstance = dbi
	})
	return dbInstance
}