.PHONY:

test:
	cat assets/3x3-example-1.txt | go run cmd/puzzle/puzzle.go

benchmark:
	go test -cpuprofile=cpu.out -memprofile=mem.out -bench=. -benchmem -benchtime=5s ./pkg/solver

profile:
	go tool pprof -http -port=8080 cpu.out mem.out
