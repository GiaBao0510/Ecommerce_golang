# câu lệnh FROM này để xác định Image gốc mà Container sẽ dựa vào đó để tạo nên
FROM golang:1.25-alpine AS builder

#   2) Tạo thư mục làm việc bên trong container
# Từ đây các lệnh tiếp theo sẽ chạy trong /build
WORKDIR /build

#   COPY câu lệnh này để sao chép các file từ máy tính của bạn vào Container, 
#trong trường hợp này nó sẽ sao chép tất cả các file trong thư mục hiện tại (.) vào thư mục /build trong Container
COPY . .

# Câu lệnh này dùng để tải tất cả các dependencies được liệt kê trong file go.mod và go.sum, giúp đảm bảo rằng tất cả các thư viện cần thiết để xây dựng ứng dụng đều có sẵn trong Container.
RUN go mod download

#   Lệnh này giúp chạy quá trình build ứng dụng Go, nó sẽ trỏ đến đúng thư mục chứa file main.go 
# (trong trường hợp này là ./cmd/server) và tạo ra một file thực thi có tên e-commerce.com trong thư mục /app/e-commerce-golang của Container.
RUN go build -o e-commerce.com ./cmd/server

# Câu lệnh này giúp kích thước của image cuối cùng nhỏ hơn bằng cách chỉ lấy file thực thi đã được biên dịch từ bước trước và bỏ qua tất cả các file nguồn và thư viện không cần thiết
FROM scratch

# Câu lệnh này để sao chép file thực thi đã được biên dịch từ bước trước vào thư mục /app/e-commerce-golang trong Container,
COPY ./configs /configs

COPY --from=builder /build/e-commerce.com /

# EXPOSE câu lệnh này để chỉ định cổng mà ứng dụng trong Container sẽ lắng nghe, trong trường hợp này là cổng 8009
EXPOSE 8080

# 7) Lệnh chạy chính khi container start
CMD ["/e-commerce.com", "config/local.yaml"]