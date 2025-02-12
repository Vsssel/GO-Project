package main

import (
	"booking-app/internal/db"
	"booking-app/internal/env"
	"booking-app/internal/store"
	"log"
)

func main() {
	config := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgres://assel:password@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIddleTime: env.GetString("DB_MAX_IDLE_TIME", "15min"),

		},
	}

	db, err := db.New(config.db.addr,config.db.maxOpenConns, config.db.maxIdleConns, config.db.maxIddleTime)

	if err != nil {
		log.Panic(err)
	}
	
	store := store.NewStorage(db)

	app := &application{config: config, store: store}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
