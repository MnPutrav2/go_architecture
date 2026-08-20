include .env

install:
	go run ./cmd/_cli install

migrate:
	go run ./cmd/_cli migrate

rollback:
	go run ./cmd/_cli rollback

run:
	go run ./cmd/server

build:
	go run ./cmd/_cli build

start:
	./build/app/server

dev:
	go run ./cmd/_cli dev

template:
	go run ./cmd/_cli make:template name=$(name) type=$(type)

help:
	go run ./cmd/_cli help