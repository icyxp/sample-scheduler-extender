package main

import (
	"log"
	"math/rand"
	"net/http"
	"sample-scheduler-extender/controller"
	"time"

	"github.com/julienschmidt/httprouter"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	router := httprouter.New()
	router.GET("/", controller.Index)
	router.POST("/filter", controller.Filter)
	router.POST("/prioritize", controller.Prioritize)

	log.Printf("start up sample-scheduler-extender!\n")
	log.Fatal(http.ListenAndServe(":8888", router))
}
