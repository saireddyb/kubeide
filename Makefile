.PHONY: run

build:
	go build -o myprogram main.go

run:
	cd modules/api && go run main.go