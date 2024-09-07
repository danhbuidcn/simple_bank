# Sử dụng hình ảnh Go chính thức
FROM golang:1.22-alpine

# Cài đặt make và postgresql-client
RUN apk add --no-cache make postgresql-client gcc musl-dev

# Cài đặt Go migration tool
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Đảm bảo thư mục Go bin đã có trong PATH
ENV PATH=$PATH:/go/bin

# Thiết lập thư mục làm việc bên trong container
WORKDIR /app

# Sao chép tệp go.mod và go.sum trước để tận dụng cache
COPY go.mod go.sum ./

# Cài đặt các phụ thuộc
RUN go mod tidy

# Sao chép toàn bộ mã nguồn vào container
COPY . .

# Biên dịch ứng dụng
RUN go build -o /app/main .

# Chỉ định cổng mà ứng dụng sẽ sử dụng
EXPOSE 8080

# Chạy ứng dụng khi container khởi động
CMD ["/app/main"]
