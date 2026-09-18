# Dockerfile riêng cho backup-service — TÁCH BIỆT hoàn toàn khỏi
# Dockerfile của API server chính (image nhẹ, chỉ chứa đúng thứ cần).

# ------ Giai đoạn 1: Build ứng dụng Go ------
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /backup-service ./cmd/backup

# ------ Giai đoạn 2: Tạo image nhẹ cho backup-service ------
FROM alpine:3.20

# Cài pg_dump/pg_restore — client tools của PostgreSQL.
# [LƯU Ý] Nên chọn version client KHỚP với version server 
RUN apk add --no-cache postgresql15-client tzdata

WORKDIR /app
COPY --from=builder /backup-service .
COPY --from=builder /app/configs ./configs

# Thư mục lưu backup — sẽ được mount ra Docker volume trong docker-compose
RUN mkdir -p /backups

ENTRYPOINT ["/app/backup-service"]