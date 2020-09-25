package main

import (
	"fmt"
	"short_url/app/config"
	"short_url/app/internal/db"
	"short_url/app/internal/server"
)

func main() {
	// init config
	conf := config.New()
	// init database
	fmt.Print(conf.Mongo.Url())
	mongo, err := db.New(conf)
	if err != nil {
		fmt.Print(err)
	}
	// init server
	server.New(conf, mongo)
}
