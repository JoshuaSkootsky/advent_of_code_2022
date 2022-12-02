day1a:
	go run cmd/day1/day1a.go 

day1b:
	go run cmd/day1/day1b.go

day1a-build:
	mkdir -p bin/
	go build -o bin/day1a.out cmd/day1/day1a.go

