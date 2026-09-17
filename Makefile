SERVER_NAME = server
BACKUP_NAME = backup-service
BIN_DIR = ./bin

# _____ chạy trực tiếp _________
run := run-server

run-server:
	go run ./cmd/server/main.go

run-backup:
	go run ./cmd/cronjob/backupsDB.go

dev:
	air

# ______ build binary ______
build: build-server build-backup

build-server:
	go build -o $(BIN_DIR)/$(SERVER_NAME) ./cmd/server

build-backup:
	go build -o $(BIN_DIR)/$(BACKUP_NAME) ./cmd/cronjob

#_________ Docker _________
docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-logs-backup:
	docker compose logs -f backup-service

# _________ Tiện ích khác _________
test:
	go test ./... -v

tidy:
	go mod tidy

help:                    ## Liệt kê các lệnh có sẵn
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'