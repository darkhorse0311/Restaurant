package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/reynld/carbtographer/pkg/database"
	"github.com/reynld/carbtographer/pkg/server"
	"github.com/reynld/carbtographer/pkg/utils"
)

func main() {
	godotenv.Load()
	s := server.Server{}

	if err := utils.CheckEnviroment(); err != nil {
		log.Fatal(err)
	}

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s.InitializeServer()
	s.Run()
}
