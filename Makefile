DOMAIN_NAME=$(shell uname -a | awk '{print $$2}')

all:
	make build
	make up

build:
	docker compose build

up:
	docker compose up

down:
	docker compose down

clean:
	docker compose down --rmi all --volumes --remove-orphans

.PHONY: all build up down clean