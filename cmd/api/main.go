package main

import (
	"log"

	"github.com/anandhansreekumar/go-social/internal/env"
)

func main() {
	config := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
