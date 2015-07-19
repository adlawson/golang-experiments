package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	sequencer := NewSequencer()
	dispatcher := NewDispatcher(sequencer)

	router := httprouter.New()
	router.GET("/", dispatcher.Total)
	router.POST("/", dispatcher.Increment)

	log.Fatal(http.ListenAndServe(":3000", router))
}
