run:
	go run cmd/main.go
build:
	go build cmd/main.go
run-docker:
	docker-compose -f docker-compose.yml up