## Mô hình ACID của cơ sở dữ liệu, đảm bảo rằng các giao dịch (transactions) được thực thi một cách tin cậy và bảo mật

### 1. **Atomicity** (Tính nguyên tử)
- Một giao dịch phải thực hiện **toàn bộ** hoặc **không thực hiện gì** cả.
- Nếu có bất kỳ lỗi nào xảy ra trong giao dịch, hệ thống sẽ rollback (hoàn tác) về trạng thái ban đầu.
- **Ví dụ**: Chuyển tiền giữa hai tài khoản, nếu trừ tiền thành công nhưng cộng tiền thất bại, giao dịch sẽ bị hoàn tác hoàn toàn.

### 2. **Consistency** (Tính nhất quán)
- Giao dịch phải đảm bảo dữ liệu chuyển từ **trạng thái nhất quán này** sang **trạng thái nhất quán khác**.
- Các quy tắc ràng buộc và tính toàn vẹn của dữ liệu luôn được duy trì trước và sau giao dịch.
- **Ví dụ**: Trong một cơ sở dữ liệu ngân hàng, tổng số tiền của tất cả tài khoản trước và sau khi giao dịch phải không thay đổi.

### 3. **Isolation** (Tính cô lập)
- Mỗi giao dịch phải được thực thi trong **môi trường cô lập**, không bị ảnh hưởng bởi các giao dịch khác đang chạy.
- Kết quả của một giao dịch chưa hoàn thành sẽ không được thấy bởi các giao dịch khác.
- **Ví dụ**: Nếu hai người dùng đang đồng thời chỉnh sửa cùng một tài khoản, thì chỉ có một người được phép hoàn tất giao dịch trước khi giao dịch của người còn lại được xử lý.

### 4. **Durability** (Tính bền vững)
- Sau khi giao dịch hoàn thành, mọi thay đổi sẽ được **lưu trữ vĩnh viễn**, ngay cả khi hệ thống gặp sự cố (như mất điện).
- Dữ liệu đã cam kết phải có khả năng phục hồi sau khi gặp sự cố.
- **Ví dụ**: Sau khi hoàn thành giao dịch chuyển tiền, số dư mới của các tài khoản sẽ được lưu lại ngay cả khi hệ thống bị tắt đột ngột.

Mô hình ACID giúp đảm bảo tính toàn vẹn và độ tin cậy cho dữ liệu trong hệ thống cơ sở dữ liệu.