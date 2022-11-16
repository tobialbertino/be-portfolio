migrate create -ext sql -dir migrations/posgres create_table_to_do

migrate -database "postgres://postgres:postgres@localhost:5432/portfolio?sslmode=disable" -path migrations/posgres up

migrate -database "postgres://postgres:postgres@localhost:5432/portfolio?sslmode=disable" -path migrations/posgres down

# force dirty state to cer...
migrate -database "postgres://postgres:postgres@localhost:5432/portfolio?sslmode=disable" -path migrations/posgres force ver

# check version 
migrate -database "postgres://postgres:postgres@localhost:5432/portfolio?sslmode=disable" -path migrations/posgres version

# additional task
migrate create -ext sql -dir migrations/posgres create_table_first
migrate create -ext sql -dir migrations/posgres create_table_second
migrate create -ext sql -dir migrations/posgres create_table_third

migrate create -ext sql -dir migrations/posgres sample_dirty_state