## 1. **Context là gì?**
- **`context.Context`** là một interface trong Go, giúp bạn quản lý thời gian sống (lifecycle) của các tác vụ (tasks) trong chương trình.
- Nó được sử dụng để truyền thông tin, đặt deadline, timeouts, và hủy bỏ các tác vụ một cách đồng bộ.

## 2. **Các loại context chính:**
   Go có 4 loại context phổ biến, mỗi loại có cách sử dụng khác nhau:

   - **`context.Background()`**:
     - Đây là context mặc định, rỗng và không có thông tin gì.
     - Thường dùng ở phần đầu của ứng dụng hoặc khi không cần hủy hoặc cài deadline.
     - Ví dụ: `ctx := context.Background()`

   - **`context.TODO()`**:
     - Được sử dụng khi bạn không chắc chắn sẽ sử dụng context gì, thường để placeholder khi phát triển ứng dụng.
     - Không có thêm thông tin hoặc điều khiển.
     - Ví dụ: `ctx := context.TODO()`

   - **`context.WithCancel(parent context.Context)`**:
     - Tạo ra một context mới dựa trên `parent`, có khả năng bị hủy (cancel).
     - Khi gọi hàm `cancel()`, tất cả các tác vụ con sử dụng context này sẽ bị hủy.
     - Dùng khi cần dừng một tác vụ bất kỳ lúc nào.
     - Ví dụ: 
       ```go
       ctx, cancel := context.WithCancel(context.Background())
       go func() {
           // tác vụ bất đồng bộ
       }()
       cancel() // hủy context và dừng tác vụ
       ```

   - **`context.WithTimeout(parent context.Context, timeout time.Duration)`**:
     - Tạo context với một khoảng thời gian chờ nhất định.
     - Khi hết thời gian, context tự động bị hủy.
     - Dùng để thiết lập thời gian tối đa cho các tác vụ (thường là API hoặc database).
     - Ví dụ:
       ```go
       ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
       defer cancel() // hủy context sau khi tác vụ hoàn tất hoặc hết thời gian
       ```

   - **`context.WithDeadline(parent context.Context, d time.Time)`**:
     - Tạo context với một deadline (thời điểm kết thúc cụ thể).
     - Giống như `WithTimeout` nhưng bạn có thể đặt một thời điểm cụ thể thay vì thời lượng.
     - Ví dụ:
       ```go
       deadline := time.Now().Add(5 * time.Second)
       ctx, cancel := context.WithDeadline(context.Background(), deadline)
       defer cancel()
       ```

## 3. **Cách sử dụng `Context` trong thực tế:**
   Khi bạn có một tác vụ chạy trong một khoảng thời gian dài như gọi API, truy vấn database, hoặc thực hiện các công việc nền (background jobs), bạn cần `context` để:

   - **Quản lý thời gian**: Đặt thời gian chờ (timeout) để các thao tác không kéo dài vô thời hạn.
   - **Hủy bỏ tác vụ khi không cần**: Nếu client hủy một yêu cầu HTTP, bạn cũng có thể hủy các tác vụ liên quan trong server.
   - **Truyền thông tin**: `Context` có thể chứa các giá trị như mã xác thực (authentication token), ID người dùng, hoặc các thông tin khác.

## 4. **Ví dụ cách dùng context trong các giao dịch database:**

   ```go
   func TransferTx(ctx context.Context, fromAccountID, toAccountID, amount int64) error {
       // bắt đầu transaction với context
       tx, err := db.BeginTx(ctx, nil)
       if err != nil {
           return err
       }
       
       // thực hiện các bước trong transaction, ví dụ tạo transfer và cập nhật balance
       err = CreateTransfer(ctx, tx, fromAccountID, toAccountID, amount)
       if err != nil {
           // hủy transaction nếu có lỗi
           tx.Rollback()
           return err
       }
       
       // commit transaction nếu thành công
       return tx.Commit()
   }
   ```

## 5. **Lợi ích của việc dùng `context.Context`:**
   - **Dễ dàng kiểm soát vòng đời của các tác vụ**: Hủy một tác vụ khi nó không còn cần thiết.
   - **Hạn chế lãng phí tài nguyên**: Ngăn chặn việc một tác vụ chạy quá lâu khi nó không còn hữu ích (ví dụ, khi client đã hủy request).
   - **Đồng bộ giữa các tác vụ**: Mọi thao tác đều có thể bị hủy hoặc timeout đồng bộ nhờ vào context.

`Context` giúp cho ứng dụng Go của bạn trở nên linh hoạt và mạnh mẽ hơn trong việc quản lý các tác vụ và sử dụng tài nguyên hiệu quả.
