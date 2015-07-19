# Buffered Worker Pool

This is a prototype of an HTTP server with a buffered (non-blocking) worker
pool. It's based on the [HTTP Sequencer][http-seq] prototype but rather than
directly increment the sequencer, workers do the job asynchronously. The workers
wait for 5 seconds before incrementing the sequencer to simulate "work". The
current `poolSize` is `4` but this can be changed to demonstrate blocked input.

## How to run
```bash
$shell_a> go get -v ./... # Download dependencies
$shell_a> go run main.go dispatcher.go sequencer.go worker.go
```
```bash
$shell_b> curl localhost:3000
0
$shell_b> curl -X POST localhost:3000
# ... < 5 seconds later
$shell_b> curl localhost:3000
0
# ... > 5 seconds later
$shell_b> curl localhost:3000
1
```

You'll also see logging output in `shell_a` so you can see the status of "work"
as it's being done. For example:
```
Worker Pool: add to buffer
Worker Pool: remove from buffer
Worker 2: starting
Worker Pool: add to buffer
Worker Pool: remove from buffer
Worker 4: starting
Worker 2: done
Worker Pool: add to buffer
Worker Pool: remove from buffer
Worker 5: starting
Worker 4: done
Worker 5: done
```


[http-seq]: /http-sequencer
