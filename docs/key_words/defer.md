# `defer`

## Cách `defer` Hoạt Động

Trong Go, `defer` được sử dụng để trì hoãn việc thực thi một hàm cho đến khi hàm bao quanh nó kết thúc. Điều này rất hữu ích để đảm bảo rằng các tài nguyên như kết nối cơ sở dữ liệu, file, hoặc khóa được giải phóng một cách an toàn.

### Ví dụ:

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
```

**Output:**

```
hello
world
```

## Tại Sao `defer` Không Làm Ứng Dụng Bất Ổn Định

Sử dụng `defer` để đóng kết nối cơ sở dữ liệu không nhất thiết làm cho ứng dụng không ổn định. Nó đảm bảo rằng kết nối được đóng đúng cách khi không còn cần thiết, ngăn ngừa rò rỉ tài nguyên. Điều quan trọng là quản lý kết nối hiệu quả và đảm bảo rằng chúng được tái sử dụng khi có thể.

### Ví dụ:

```go
func queryDatabase(db *sql.DB) {
    conn, err := db.Conn(context.Background())
    if (err != nil) {
        log.Fatal(err)
    }
    defer conn.Close()
    // Thực hiện các thao tác cơ sở dữ liệu
}
```

Trong ví dụ này, `defer conn.Close()` đảm bảo rằng kết nối được đóng sau khi các thao tác cơ sở dữ liệu hoàn tất, giúp quản lý tài nguyên hiệu quả.

## Tại Sao `defer` Là Thực Hành Tốt

Sử dụng `defer` là một thực hành tốt vì nó giúp mã nguồn dễ đọc hơn và giảm thiểu lỗi bằng cách đảm bảo rằng các tài nguyên được giải phóng đúng cách. Điều này đặc biệt quan trọng trong các ứng dụng lớn và phức tạp, nơi việc quản lý tài nguyên thủ công có thể dẫn đến lỗi.

# `maxIdleConns`

`maxIdleConns` đề cập đến số lượng kết nối tối đa trong pool kết nối nhàn rỗi. Duy trì các kết nối nhàn rỗi có thể giúp cải thiện hiệu suất bằng cách giảm bớt chi phí thiết lập các kết nối mới.

## Cách `maxIdleConns` Hoạt Động

Khi một kết nối không còn được sử dụng, nó có thể được giữ lại trong pool kết nối nhàn rỗi thay vì bị đóng ngay lập tức. Điều này giúp giảm thiểu chi phí thiết lập kết nối mới khi có yêu cầu kết nối tiếp theo. Tuy nhiên, việc duy trì quá nhiều kết nối nhàn rỗi có thể tiêu tốn tài nguyên hệ thống, do đó cần phải cấu hình `maxIdleConns` một cách hợp lý.

# `maxOpenConns`

`maxOpenConns` xác định số lượng kết nối tối đa có thể mở cùng một lúc. Điều này giúp kiểm soát số lượng kết nối đồng thời tới cơ sở dữ liệu, ngăn ngừa việc sử dụng quá nhiều tài nguyên và đảm bảo rằng hệ thống không bị quá tải.

## Cách `maxOpenConns` Hoạt Động

Khi số lượng kết nối mở đạt đến giới hạn `maxOpenConns`, các yêu cầu kết nối mới sẽ phải chờ cho đến khi một kết nối hiện tại được đóng. Điều này giúp duy trì hiệu suất ổn định và ngăn ngừa tình trạng quá tải hệ thống.

# `connMaxLifetime`

`connMaxLifetime` xác định thời gian tối đa mà một kết nối có thể tồn tại. Sau khi đạt đến giới hạn này, kết nối sẽ bị đóng và loại bỏ khỏi pool, ngay cả khi nó vẫn đang nhàn rỗi.

## Cách `connMaxLifetime` Hoạt Động

Thiết lập `connMaxLifetime` giúp đảm bảo rằng các kết nối cũ không tồn tại quá lâu, ngăn ngừa các vấn đề tiềm ẩn liên quan đến kết nối lâu dài như rò rỉ bộ nhớ hoặc các lỗi kết nối không mong muốn. Điều này giúp duy trì hiệu suất và độ tin cậy của ứng dụng.
