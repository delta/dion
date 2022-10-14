build:
	go build -v .

# start a docker container for postgres
postgres:
	docker compose -f ./docker/postgres.yml up -d

postgres-down:
	docker compose -f ./docker/postgres.yml down

install:
	go get ./...
	cp config/config.example.yaml config/config.yaml
	cp docker/example.env docker/.env

dev:
	air

run:
	go run main.go

clean:
	go mod tidy
	rm -rf build
