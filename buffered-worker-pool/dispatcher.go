package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Dispatcher struct {
	sequencer *Sequencer
	worker    chan int
}

func NewDispatcher(sequencer *Sequencer, worker chan int) *Dispatcher {
	return &Dispatcher{sequencer, worker}
}

func (d *Dispatcher) Total(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, d.sequencer.Total)
}

func (d *Dispatcher) Increment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	d.worker <- 1
}
