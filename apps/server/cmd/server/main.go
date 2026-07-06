// @title		Server API
// @version		1.0
// @description	Modular monolith REST API for React web and native clients.
// @host		localhost:18080
// @BasePath	/api/v1
// @schemes		http
package main

import (
	"flag"
	"log"

	_ "shopwise/apps/server/docs"
	"shopwise/apps/server/internal/bootstrap"
)

func main() {
	runMigrations := flag.Bool("migrate", false, "Run database migrations")
	runSeed := flag.Bool("seed", false, "Seed database with startup metadata")
	flag.Parse()

	if *runMigrations {
		if err := bootstrap.Migrate(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		return
	}

	if *runSeed {
		if err := bootstrap.Seed(); err != nil {
			log.Fatalf("Seeding failed: %v", err)
		}
		return
	}

	if err := bootstrap.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
