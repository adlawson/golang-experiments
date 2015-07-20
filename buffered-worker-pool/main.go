package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const poolSize = 4

func main() {
	sequencer := NewSequencer()
	worker := NewWorker(sequencer)
	pool := NewWorkerPool(worker, poolSize)

	queue := make(chan int)
	defer close(queue)
	go pool.ListenAndServe(queue)

	dispatcher := NewDispatcher(sequencer, queue)
	router := httprouter.New()
	router.GET("/", dispatcher.Total)
	router.POST("/", dispatcher.Increment)

	log.Fatal(http.ListenAndServe(":3000", router))
}
