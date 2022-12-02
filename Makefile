DAY_1=cmd/day1/day1.go

day1:
	go run $(DAY_1)

day1-build:
	mkdir -p bin/
	go build -o bin/day1.out DAY_1

day2:

test:
	go test ./...

