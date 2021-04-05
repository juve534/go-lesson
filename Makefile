container=app

.PHONY: build run
build:
	docker-compose build

run: build
	docker-compose up -d