package main

import (
	"fmt"
	"time"
)

func blockingWorker(i int, sequencer *Sequencer, incr <-chan int) {
	for {
		select {
		case n := <-incr:
			fmt.Printf("Worker %d: starting\n", i)
			time.Sleep(5 * time.Second)
			sequencer.Total += n
			fmt.Printf("Worker %d: done\n", i)
		default:
		}
	}
}

func workerPool(sequencer *Sequencer, notify <-chan int, poolSize int) {
	buffer := []int{} // Buffer used in case all workers are busy
	queue := make(chan int, poolSize)
	defer close(queue)
	for i := 0; i < poolSize; i++ {
		go blockingWorker(i, sequencer, queue)
	}
	for {
		select {
		case n := <-notify:
			buffer = append(buffer, n) // Add notice to the buffer (avoids "pushing in")
			fmt.Println("Worker Pool: add to buffer")
		default:
			if len(buffer) > 0 {
				select {
				case queue <- buffer[0]:
					buffer = buffer[1:]
					fmt.Println("Worker Pool: remove from buffer")
				default: // Nothing put on the queue, buffer remains the same
					fmt.Println("Worker Pool: blocked")
				}
			}
		}
	}
}
