package main

import (
	"log"

	"github.com/ahenla/go-blog/cmd/api"
	"github.com/ahenla/go-blog/config"
	"github.com/ahenla/go-blog/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
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
	server := api.NewAPIServer(":8000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
