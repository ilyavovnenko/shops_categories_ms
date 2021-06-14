api:
	go run cmd/main.go 

test: 
	go test -v -cover -covermode=atomic ./...

unittest:
	go test -short  ./...

build:
	docker-compose up --build -d
	docker-compose start

start:
	docker-compose start

restart:
	docker-compose restart

stop:
	docker-compose stop

down:
	docker-compose down

migrate:
	go run cmd/migrations/migrate.go up

rollback:
	go run cmd/migrations/migrate.go down

.PHONY: api test unittest build start restart stop down migrate rollback