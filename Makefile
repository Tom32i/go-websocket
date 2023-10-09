start:
	go run main.go

build:
	go build -o curvygo main.go

run: build
	./curvygo
