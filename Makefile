fmt:
	go fmt ./...

run:
	go run cmd/main/main.go

migrate.up:
	go run cmd/migrate/main.go

up:
	docker compose up -d

stop:
	docker compose stop

down:
	docker compose down

restart:
	docker compose restart

logs:
	docker compose logs -n 30 -f

docs-update:
	rm -rf swagger/v1
	swag fmt
	swag init -d ./cmd/main/,./  -o swagger
