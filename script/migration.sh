#!/bin/bash

migrate create -ext sql -dir migrations/posgres create_table_to_do

migrate -database "postgres://postgres:postgres@localhost:5432/portfolio?sslmode=disable" -path migrations/posgres up

migrate -database "postgres://postgres:postgres@localhost:5432/portfolio?sslmode=disable" -path migrations/posgres down

# force dirty state to clear from version
migrate -database "postgres://postgres:postgres@localhost:5432/portfolio?sslmode=disable" -path migrations/posgres force ver

# check version 
migrate -database "postgres://postgres:postgres@localhost:5432/portfolio?sslmode=disable" -path migrations/posgres version