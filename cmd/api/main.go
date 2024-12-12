package main

import (
	"log"
	"github.com/n4vxn/social/internal/store"
)

func main() {
	cfg := config{
		addr: ":8080",
	}
	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
