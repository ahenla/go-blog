package main

import (
	"database/sql"
	"log"

	"github.com/ahenla/go-blog/cmd/api"
	"github.com/ahenla/go-blog/config"
	"github.com/ahenla/go-blog/db"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := db.NewSQLDB(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAdress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	connectDB(db)

	server := api.NewAPIServer(":8000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func connectDB(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("DB: connection successful!")
}
