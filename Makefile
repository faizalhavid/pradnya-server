DB_URL=postgres://postgres:postgres@localhost:5433/pradnya?sslmode=disable

dev:
	air

run:
	go run cmd/api/main.go

seed:
	go run cmd/seed/main.go
	go run cmd/seed/main.go

docker-up:
	docker compose up -d

docker-down:
	docker compose down

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

redocs:
	swag init -g cmd/api/main.go