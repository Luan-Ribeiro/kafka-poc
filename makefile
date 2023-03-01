SHELL := /bin/bash

# =====================================
# Variable

VERSION := 1.0
PROJECT_NAME := kafka-poc

run:
	go run ./producer/main.go
	go run ./consumer/main.go

## Download libraries
tidy:
	go mod tidy
	go mod vendor

# =====================================
# Enviroment

env-up:
	docker-compose up -d

env-setup:
	docker exec broker kafka-topics --create --bootstrap-server \
        localhost:9092 --replication-factor 1 --partitions 1 --topic purchase

env-down:
	docker-compose down