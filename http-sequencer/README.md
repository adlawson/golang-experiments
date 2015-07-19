# HTTP sequencer

This is a prototype of an HTTP sequencer that simply counts the number of POST
requests received. It's not very interesting, but I'll likely use the concept
as a boilerplate for other experiments.

## How to run
```bash
$shell_a> go get -v ./... # Download dependencies
$shell_a> go run main.go dispatcher.go persistence.go
```
```bash
$shell_b> curl localhost:3000
0
$shell_b> curl -X POST localhost:3000
$shell_b> curl localhost:3000
1
```
