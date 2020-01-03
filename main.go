package main

import (
<<<<<<< HEAD
	"github.com/cnych/sample-scheduler-extender/controller"
=======
>>>>>>> 047fdd5ae8b1a6d7fdc0e6d20ce4d70a1d6e7178
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	router := httprouter.New()
<<<<<<< HEAD
	router.GET("/", controller.Index)
	router.POST("/filter", controller.Filter)
	router.POST("/prioritize", controller.Prioritize)

	log.Printf("start up sample-scheduler-extender!\n")
	log.Fatal(http.ListenAndServe(":8888", router))
}
=======
	router.GET("/", Index)
	router.POST("/filter", Filter)
	router.POST("/prioritize", Prioritize)

	log.Fatal(http.ListenAndServe(":8888", router))
}
>>>>>>> 047fdd5ae8b1a6d7fdc0e6d20ce4d70a1d6e7178
