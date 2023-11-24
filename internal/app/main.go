package main

import (
	"log"
	"server/internal/config"
	"server/internal/model"
	"server/internal/router"
)

func main() {
	var err error

	// load configuration
	err = config.LoadConfig()
	if err != nil {
		log.Println(err)
		log.Fatalln("Failed loading configuration")
	}

	if config.Conf.General.Debug {
		log.Println("Running in DEBUG mode")
	} else {
		log.Println("Running in RELEASE mode")
	}

	// connect to database
	err = model.ConnectDatabase()
	if err != nil {
		log.Println(err)
		log.Fatalln("Failed connecting to the database")
	}

	// run server
	err = router.RunServer()
	if err != nil {
		log.Fatalln(err)
	}
}
