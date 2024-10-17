### Định nghĩa Deadlock

**Deadlock** (hay **bế tắc**) là một trạng thái trong hệ thống khi hai hay nhiều tiến trình (trong trường hợp này là các giao dịch) bị chặn chờ lẫn nhau để giải phóng các tài nguyên mà mỗi tiến trình đang giữ. Điều này khiến tất cả các tiến trình bị kẹt và không thể tiếp tục thực hiện.

### Deadlock trong trường hợp chuyển tiền

Trong ngữ cảnh của hàm `TransferTx`, deadlock xảy ra khi có **hai giao dịch đồng thời** đang cố gắng **chuyển tiền giữa cùng hai tài khoản**, nhưng theo **thứ tự ngược nhau**. 

#### Ví dụ cụ thể:
- Giả sử chúng ta có hai tài khoản `account1` và `account2`.
- Giao dịch 1: Chuyển từ `account1` sang `account2`.
- Giao dịch 2: Chuyển từ `account2` sang `account1`.

Trong trường hợp này:
1. **Giao dịch 1** sẽ bắt đầu bằng cách **khóa `account1`** trước rồi đến **khóa `account2`** để thực hiện chuyển tiền.
2. **Giao dịch 2** lại bắt đầu bằng cách **khóa `account2`** trước rồi đến **khóa `account1`**.
3. Nếu cả hai giao dịch này diễn ra đồng thời, chúng sẽ rơi vào tình trạng deadlock:
   - Giao dịch 1 chờ Giao dịch 2 giải phóng `account2`.
   - Giao dịch 2 chờ Giao dịch 1 giải phóng `account1`.
   - Cả hai giao dịch đều không thể tiếp tục vì mỗi giao dịch đang chờ tài nguyên (tài khoản) bị khóa bởi giao dịch kia.

### Giải pháp tránh Deadlock

Deadlock có thể tránh được nếu **các giao dịch khóa tài nguyên theo thứ tự nhất quán**. Trong trường hợp này, khi thực hiện giao dịch chuyển tiền giữa hai tài khoản, ta nên:
1. **Khóa tài khoản theo thứ tự ID**, luôn khóa tài khoản có ID nhỏ trước rồi đến tài khoản có ID lớn, bất kể hướng của giao dịch (từ tài khoản nào sang tài khoản nào).
2. Bằng cách này, tất cả các giao dịch sẽ khóa tài khoản theo cùng thứ tự, đảm bảo không có trường hợp khóa tài khoản ngược chiều, từ đó tránh được deadlock.

### Tóm tắt:
- **Deadlock** xảy ra khi hai giao dịch đồng thời cố gắng khóa hai tài khoản theo thứ tự ngược nhau.
- Để **tránh deadlock**, cần đảm bảo rằng tất cả các giao dịch luôn khóa tài khoản theo thứ tự cố định (ví dụ, ID nhỏ trước, ID lớn sau).
