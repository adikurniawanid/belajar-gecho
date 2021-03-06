package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"template-api-gecho/constant"
	"template-api-gecho/utils"
)

var db *pgxpool.Pool

type Database struct{}

func init() {
	var err error

	host := os.Getenv(constant.DBHost)
	dbname := os.Getenv(constant.DBName)
	user := os.Getenv(constant.DBUser)
	password := os.Getenv(constant.DBPassword)
	schema := os.Getenv(constant.DBSchema)

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s", user, password, host, dbname, schema)
	db, err = pgxpool.Connect(context.Background(), psqlInfo)

	if err != nil {
		utils.LogFatal(err)
	}
}

func (Database Database) GetDatabase() *pgxpool.Pool {
	return db
}
