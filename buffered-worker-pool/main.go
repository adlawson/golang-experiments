package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const poolSize = 4

func main() {
	sequencer := NewSequencer()
	notify := make(chan int)
	defer close(notify)
	go workerPool(sequencer, notify, poolSize)

	dispatcher := NewDispatcher(sequencer, notify)
	router := httprouter.New()
	router.GET("/", dispatcher.Total)
	router.POST("/", dispatcher.Increment)

	log.Fatal(http.ListenAndServe(":3000", router))
}
