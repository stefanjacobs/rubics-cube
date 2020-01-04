run:
	go run main.go cube.go cubeColor.go cubeState.go priorityq.go solver.go state.go

run_profile:
	go run main.go cube.go cubeColor.go cubeState.go priorityq.go solver.go state.go -cpuprofile cpu.prof -memprofile mem.prof

profile:
	go tool pprof cpu.prof

build:
	go build rubic.go

clean:
	rm rubic

test:
	go test