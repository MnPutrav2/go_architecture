include .env

install:
	go run ./cmd/_cli install

migrate:
	go run ./cmd/_cli migrate

rollback:
	go run ./cmd/_cli rollback

build:
	go run ./cmd/_cli build

run:
	go run ./cmd/server

start:
	./final/app/server

template:
	go run ./cmd/_cli make:template name=$(name) type=$(type)

help:
	go run ./cmd/_cli help