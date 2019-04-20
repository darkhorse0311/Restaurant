package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/reynld/carbtographer/pkg/cache"
	"github.com/reynld/carbtographer/pkg/database"
	"github.com/reynld/carbtographer/pkg/scraper"
	"github.com/reynld/carbtographer/pkg/server"
	"github.com/reynld/carbtographer/pkg/utils"
)

func main() {
	godotenv.Load()
	if err := utils.CheckEnviroment(); err != nil {
		log.Fatal(err)
	}

	scrape := flag.Bool("scrape", false, "scrapes default restaurant macro info")
	migrate := flag.Bool("migrate", false, "migrates database")
	seed := flag.Bool("seed", false, "seeds database")
	seedCache := flag.Bool("seed-cache", false, "seeds redis cache")
	serve := flag.Bool("serve", false, "runs server")
	flag.Parse()

	if len(os.Args) > 1 {
		if flag.NFlag() != 1 {
			fmt.Println("pass just one argument")
			flag.Usage()
			os.Exit(1)
		}

		s := server.Server{}
		s.Initialize()

		if *scrape {
			scraper.RunScraper()
		}
		if *serve {
			s.Run()
		}
		if *migrate {
			database.RunMigrations(s.DB)
		}
		if *seed {
			database.RunSeeds(s.DB)
		}
		if *seedCache {
			cache.RunSeeds(s.Cache)
		}

	} else {
		fmt.Println("pass at least one argument")
		flag.Usage()
	}
}
