package main

import (
	"log"

	"github.com/iykeevans/go-social/server/internal/db"
	"github.com/iykeevans/go-social/server/internal/env"
	"github.com/iykeevans/go-social/server/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5433/go_social_db?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
