api:
	go run . api

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
	go run . migrate up

rollback:
	go run . migrate down

parsing_bol_com:
	go run . parse bol.com
	
parsing_amazon_de:
	go run . parse amazon.de

.PHONY: api test unittest build start restart stop down migrate rollback parsing_bol api