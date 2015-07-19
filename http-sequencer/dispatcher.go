package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Dispatcher struct {
	sequencer *Sequencer
}

func NewDispatcher(sequencer *Sequencer) *Dispatcher {
	return &Dispatcher{sequencer}
}

func (d *Dispatcher) Total(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, d.sequencer.Total)
}

func (d *Dispatcher) Increment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	d.sequencer.Increment(1)
}
