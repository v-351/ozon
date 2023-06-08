package main

import (
	"flag"

	"github.com/v-351/ozon/internal/app"
)

func main() {
	postgresFlag := flag.Bool("postgres", false, "using PostgreSQL")
	flag.Parse()
	app.Run(postgresFlag)
}
