package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	m, err := migrate.New(
		"D:\\projectX\\msrvs\\msrv-bot-tg\\migrations",
		"postgres://postgres:postgres@localhost:5430/msrv-bot-tg?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
