package main

import (
	"fmt"
	"time"
)

type Worker struct {
	sequencer *Sequencer
}

func NewWorker(sequencer *Sequencer) *Worker {
	return &Worker{sequencer}
}

func (w *Worker) Apply(n int) {
	time.Sleep(5 * time.Second)
	w.sequencer.Total += n
}

type WorkerPool struct {
	worker *Worker
	size   int
}

func NewWorkerPool(worker *Worker, poolSize int) *WorkerPool {
	return &WorkerPool{worker, poolSize}
}

func (w *WorkerPool) ListenAndServe(receive <-chan int) {
	buffer := []int{} // Buffer used in case all workers are busy
	send := make(chan int, poolSize)
	defer close(send)
	for i := 0; i < poolSize; i++ {
		go runWorker(i, w.worker, send)
	}
	for {
		select {
		case n := <-receive: // Take from `receive` and add to FIFO `buffer`
			buffer = append(buffer, n)
			fmt.Println("Worker Pool: add to buffer")
		default:
		}
		if len(buffer) > 0 {
			select {
			case send <- buffer[0]: // Take `buffer` head and send to a worker
				buffer = buffer[1:]
				fmt.Println("Worker Pool: remove from buffer")
			default:
				fmt.Println("Worker Pool: blocked")
			}
		}
	}
}

func runWorker(i int, worker *Worker, receive <-chan int) {
	for {
		select {
		case n := <-receive:
			fmt.Printf("Worker %d: starting\n", i)
			worker.Apply(n)
			fmt.Printf("Worker %d: done\n", i)
		default:
		}
	}
}
