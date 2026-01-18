DB_URL=postgres://igormilosavljevic:@localhost:5432/gator?sslmode=disable

migrate-up:
	goose -dir sql/schema postgres "$(DB_URL)" up

migrate-down:
	goose -dir sql/schema postgres "$(DB_URL)" down

sqlc:
	sqlc generate

run:
	go run .
