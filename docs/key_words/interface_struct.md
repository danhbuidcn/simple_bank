## 1. Interface

```go
// Store provides all functions to execute DB queries and transaction
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}
```

- `Interface`: Interface là một tập hợp các phương thức mà một kiểu dữ liệu (struct) phải triển khai. Interface không chứa bất kỳ logic nào, chỉ định nghĩa các phương thức mà một struct phải có.

- `Store Interface`: Interface Store định nghĩa hai thành phần:
    + `Querier`: Đây có thể là một interface khác mà Store kế thừa. Querier có thể chứa các phương thức để thực hiện các truy vấn cơ sở dữ liệu.
    + `TransferTx`: Phương thức này thực hiện một giao dịch chuyển khoản và trả về kết quả của giao dịch hoặc lỗi nếu có.

## 2. Struct

```
// SQLStore provides all functions to execute SQL queries and transaction
type SQLStore struct {
    db *sql.DB
    *Queries
}
```

- `Struct`: Struct là một kiểu dữ liệu tùy chỉnh trong Go, cho phép bạn nhóm các trường dữ liệu lại với nhau.

- `SQLStore Struct`: Struct SQLStore chứa hai trường:
    + `db`: Một con trỏ tới `sql.DB`, đại diện cho kết nối cơ sở dữ liệu.
    + `Queries`: Một con trỏ tới Queries, có thể là một struct khác chứa các phương thức để thực hiện các truy vấn SQL.

## 3. Hàm NewStore

```
// NewStore creates a new store
func NewStore(db *sql.DB) Store {
    return &SQLStore{
        db:      db,
        Queries: New(db),
    }
}
```

- `Hàm`: Hàm là một khối mã thực hiện một nhiệm vụ cụ thể.

- `NewStore Function`: Hàm `NewStore` tạo và trả về một đối tượng `Store`.
    + `Tham số`: Hàm nhận một tham số db là một con trỏ tới sql.DB.
    + `Trả về`: Hàm trả về một đối tượng `Store`, cụ thể là một đối tượng `SQLStore`.

## Tại sao hàm trả về interface nhưng trả về struct

- `Trả về Interface`: Hàm `NewStore` trả về một giá trị có kiểu `Store`, là một interface. Điều này cho phép bạn sử dụng bất kỳ struct nào triển khai interface `Store`.

- `Trả về Struct`: Mặc dù hàm trả về một interface, nhưng giá trị thực sự được trả về là một đối tượng `SQLStore`, một struct triển khai interface `Store`. Điều này cho phép bạn tận dụng tính linh hoạt của interface trong khi vẫn sử dụng một struct cụ thể.

## Tóm tắt

- `Interface Store`: Định nghĩa các phương thức mà một store phải có, bao gồm các phương thức từ `Querier` và `TransferTx`.

- `Struct SQLStore`: Triển khai các phương thức được định nghĩa trong interface `Store` và chứa các trường dữ liệu cần thiết để thực hiện các truy vấn và giao dịch SQL.

- `Hàm NewStore`: Tạo và trả về một đối tượng `SQLStore`, nhưng kiểu trả về là interface `Store` để tận dụng tính linh hoạt của interface.
