DAY_1=cmd/day1/day1.go
DAY_2=cmd/day2/day2.go
DAY_03=cmd/day03/main.go
DAY_04=cmd/day04/main.go

day1:
	go run $(DAY_1)

day1-build:
	mkdir -p bin/
	go build -o bin/day1.out DAY_1

day2:
	go run $(DAY_2)

day03:
	go run $(DAY_03)

day04:
	go run $(DAY_04)

test:
	go test ./...

