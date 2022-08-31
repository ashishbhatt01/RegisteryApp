package main

import (
	"log"

	"github.com/ashishbhatt01/registeryApp/app/logging"
	"github.com/ashishbhatt01/registeryApp/app/route"

	gin "github.com/gin-gonic/gin"
)

const listenPort = ":4000"

func main() {
	log.Println("Starting the app")
	gin.SetMode(gin.ReleaseMode)

	ginEngine := gin.Default()

	logging.Initialize()
	log.Println("logger initialized successfully")

	route.Initialize(ginEngine)
	log.Println("routes initialized successfully")

	if err := ginEngine.Run(listenPort); err != nil {
		log.Fatal("gin engine failed to run", err.Error())
	}
}
