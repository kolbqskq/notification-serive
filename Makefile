include .env.shared
export

migrate-up:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir migrations postgres ${DATABASE_URL} up