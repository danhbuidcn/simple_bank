## Redis là gì ?

**Redis** (Remote Dictionary Server) là một cơ sở dữ liệu lưu trữ dữ liệu trên bộ nhớ (in-memory), dạng key-value, mã nguồn mở. Nó được dùng để:

- **Lưu trữ tạm thời (caching)**: Giúp truy xuất dữ liệu nhanh chóng.
- **Quản lý session**: Lưu trữ thông tin phiên (session) cho người dùng.
- **Hàng đợi công việc**: Xử lý các công việc cần chạy theo thứ tự.
- **Pub/Sub messaging**: Cung cấp cơ chế giao tiếp giữa các dịch vụ bằng cách xuất bản (publish) và đăng ký (subscribe) thông điệp.

=> Redis nổi tiếng nhờ tốc độ rất nhanh do lưu trữ dữ liệu trực tiếp trong RAM.

## Lý do sử dụng redis: (hiệu suất cao, đồng thời cao)

### 1. **Hiệu suất cao**
- **Ý nghĩa**: Redis có khả năng xử lý hàng triệu yêu cầu mỗi giây nhờ vào việc lưu trữ dữ liệu trong RAM. 
- **Cách hoạt động**: Dữ liệu được truy xuất và ghi vào bộ nhớ nhanh hơn nhiều so với việc truy cập đĩa, giúp giảm thời gian phản hồi cho người dùng.
- **Lợi ích**: Điều này đặc biệt quan trọng cho các ứng dụng yêu cầu truy cập dữ liệu nhanh, như các dịch vụ web thời gian thực, trò chơi trực tuyến, hoặc hệ thống phân tích dữ liệu.

### 2. **Đồng thời cao**
- **Ý nghĩa**: Redis có khả năng phục vụ nhiều kết nối và yêu cầu từ nhiều client cùng lúc mà không làm giảm hiệu suất.
- **Cách hoạt động**: Redis sử dụng một mô hình xử lý đơn (single-threaded) với event-driven architecture, giúp xử lý các yêu cầu một cách hiệu quả mà không bị khóa bởi các thao tác IO. 
- **Lợi ích**: Điều này giúp các ứng dụng có khả năng mở rộng tốt, cho phép nhiều người dùng truy cập và tương tác với hệ thống mà không gặp phải độ trễ cao hoặc tắc nghẽn.

### Tóm tắt
- **Hiệu suất cao**: Redis cho phép truy xuất dữ liệu nhanh chóng nhờ lưu trữ trong RAM.
- **Đồng thời cao**: Redis có khả năng xử lý nhiều yêu cầu cùng lúc mà không làm giảm tốc độ phản hồi, giúp nâng cao trải nghiệm người dùng. 

## So sánh giữa Redis và Memcached

| **Tiêu chí**              | **Redis**                                        | **Memcached**                                  |
|---------------------------|--------------------------------------------------|------------------------------------------------|
| **Kiểu dữ liệu**          | Hỗ trợ nhiều kiểu dữ liệu (string, list, set, hash, sorted set) | Chỉ hỗ trợ kiểu dữ liệu đơn giản (key-value, string) |
| **Persistence**           | Hỗ trợ lưu trữ dữ liệu xuống đĩa (RDB, AOF), dữ liệu có thể bền vững | Không hỗ trợ lưu trữ dữ liệu, dữ liệu sẽ mất khi server khởi động lại |
| **Replication**           | Hỗ trợ sao chép dữ liệu (master-slave replication) | Không hỗ trợ sao chép dữ liệu                  |
| **Tính năng**             | Hỗ trợ Pub/Sub, transactions, scripting, và nhiều tính năng nâng cao khác | Chức năng đơn giản, không có hỗ trợ cho Pub/Sub hay transactions |
| **Lưu trữ**               | Dữ liệu lưu trữ trong RAM và có thể lưu xuống đĩa, cho phép khôi phục sau sự cố | Dữ liệu hoàn toàn lưu trữ trong RAM, không có khả năng khôi phục sau sự cố |
| **Tốc độ**                | Rất nhanh nhờ vào lưu trữ trên RAM và kiến trúc event-driven | Cũng rất nhanh nhưng tối ưu cho caching đơn giản |
| **Mô hình xử lý**         | **Đơn luồng**: Sử dụng một luồng duy nhất để xử lý các yêu cầu với event-driven architecture, giúp tránh các vấn đề đồng bộ hóa | **Đa luồng**: Hỗ trợ nhiều luồng, cho phép xử lý nhiều yêu cầu đồng thời mà không bị khóa. |
| **Sử dụng điển hình**     | Ứng dụng phức tạp, phân tích thời gian thực, game online | Caching đơn giản, tăng tốc độ truy xuất dữ liệu |

- Tóm tắt:
    - **Redis**: Sử dụng mô hình đơn luồng, xử lý nhanh và hiệu quả với nhiều tính năng và khả năng lưu trữ bền vững.
    - **Memcached**: Sử dụng mô hình đa luồng, tối ưu cho caching đơn giản nhưng không có khả năng lưu trữ bền vững.

## Redis có bao nhiêu kiểu dữ liệu và kịch bản sử dụng ?

### 1. **String**
- **Mô tả**: Kiểu dữ liệu đơn giản nhất, có thể lưu trữ chuỗi, số hoặc nhị phân.
- **Kịch bản sử dụng**:
  - Lưu trữ thông tin cấu hình.
  - Lưu trữ giá trị tạm thời như token xác thực.
  - Caching cho dữ liệu có kích thước nhỏ.

### 2. **List**
- **Mô tả**: Danh sách các chuỗi được sắp xếp theo thứ tự, có thể thêm hoặc xóa ở đầu hoặc cuối.
- **Kịch bản sử dụng**:
  - Hàng đợi (queue) để xử lý công việc (job queue).
  - Lưu trữ lịch sử hoạt động hoặc nhật ký (log).
  - Chat hoặc thông điệp (message) trong ứng dụng thời gian thực.

### 3. **Set**
- **Mô tả**: Tập hợp các chuỗi độc nhất, không có thứ tự.
- **Kịch bản sử dụng**:
  - Lưu trữ các ID người dùng hoặc tag mà không cần trùng lặp.
  - Tìm kiếm các phần tử chung giữa nhiều tập hợp.
  - Xây dựng các hệ thống theo dõi, chẳng hạn như người theo dõi (follower) trong mạng xã hội.

### 4. **Sorted Set**
- **Mô tả**: Giống như Set, nhưng mỗi phần tử có một giá trị điểm (score), cho phép sắp xếp.
- **Kịch bản sử dụng**:
  - Xếp hạng (leaderboard) trong trò chơi.
  - Quản lý thứ tự của các nhiệm vụ hoặc sự kiện.
  - Thống kê và phân tích dữ liệu thời gian thực.

### 5. **Hash**
- **Mô tả**: Tập hợp các cặp key-value, tương tự như bảng băm.
- **Kịch bản sử dụng**:
  - Lưu trữ thông tin của một đối tượng, chẳng hạn như thông tin người dùng.
  - Cấu hình cho các dịch vụ hoặc tài nguyên phức tạp.
  - Quản lý trạng thái của phiên làm việc (session state).

### 6. **Bitmaps**
- **Mô tả**: Sử dụng bit để lưu trữ dữ liệu nhị phân.
- **Kịch bản sử dụng**:
  - Theo dõi sự hiện diện (presence tracking).
  - Lưu trữ thông tin kiểu "đã xem" hoặc "đã thực hiện".

### 7. **HyperLogLog**
- **Mô tả**: Cấu trúc dữ liệu thống kê để ước lượng số lượng phần tử độc nhất.
- **Kịch bản sử dụng**:
  - Theo dõi số lượng người dùng truy cập duy nhất trên trang web.
  - Thống kê lượng giao dịch hoặc sự kiện.

### 8. **Geospatial**
- **Mô tả**: Lưu trữ và truy vấn dữ liệu địa lý.
- **Kịch bản sử dụng**:
  - Tìm kiếm các điểm gần một vị trí cụ thể.
  - Quản lý thông tin địa lý cho ứng dụng bản đồ hoặc dịch vụ vị trí.


## Redis giải quyết cơ chế hết hạn dữ liệu thông qua việc:

### 1. **Thiết lập thời gian hết hạn**
- Bạn có thể thiết lập thời gian hết hạn cho một key khi lưu trữ dữ liệu bằng cách sử dụng các lệnh như `EXPIRE` hoặc `SETEX`.
- **Cách sử dụng**:
  - **`EXPIRE key seconds`**: Đặt thời gian hết hạn cho key trong số giây cụ thể.
  - **`SETEX key seconds value`**: Thiết lập key với giá trị và thời gian hết hạn cùng một lúc.

### 2. **Kiểm tra và xóa tự động**
- Redis thực hiện việc kiểm tra và xóa các key đã hết hạn bằng hai cơ chế chính:
  - **Lazy Expiration**: Khi một key được truy cập, Redis kiểm tra xem nó có hết hạn hay không. Nếu đã hết hạn, key sẽ bị xóa. Điều này giúp tiết kiệm tài nguyên vì không cần phải kiểm tra mọi key liên tục.
  - **Active Expiration**: Redis định kỳ (khoảng 100 ms) quét qua một số lượng key nhất định để xóa các key đã hết hạn. Điều này đảm bảo rằng các key hết hạn sẽ được dọn dẹp ngay cả khi không có yêu cầu truy cập.

### 3. **Xử lý khối lượng lớn dữ liệu hết hạn**
- Redis không sử dụng chỉ một chiến lược mà kết hợp cả hai để đảm bảo hiệu suất cao và quản lý bộ nhớ hiệu quả.
- Nếu có nhiều key đã hết hạn, Redis sẽ thực hiện cả lazy và active expiration để duy trì hiệu suất và giải phóng bộ nhớ.

### 4. **Nhận thông báo khi dữ liệu hết hạn**
- Redis cũng hỗ trợ thông báo hết hạn thông qua cơ chế **Pub/Sub**. Bạn có thể nhận thông báo khi một key hết hạn bằng cách sử dụng các kênh thông báo được định nghĩa trước.

### Tóm tắt

Redis giải quyết cơ chế hết hạn dữ liệu thông qua việc:
- Thiết lập thời gian hết hạn cho các key.
- Sử dụng cơ chế kiểm tra lười biếng (lazy) và quét chủ động (active) để xóa key đã hết hạn.
- Cung cấp thông báo khi key hết hạn thông qua Pub/Sub.

## khóa bi quan (pessimistic locking) và khóa lạc quan (optimistic locking) 

- Trong Redis, khái niệm khóa bi quan (pessimistic locking) và khóa lạc quan (optimistic locking) liên quan đến cách thức quản lý đồng bộ hóa và truy cập đồng thời vào dữ liệu.

### 1. **Khóa Bi Quan (Pessimistic Locking)**
- **Khái niệm**: Đây là phương pháp khóa tài nguyên trước khi thực hiện thao tác, ngăn chặn các giao dịch khác truy cập vào tài nguyên đó cho đến khi thao tác hoàn tất.
- **Cách hoạt động**:
  - Khi một client muốn thực hiện thao tác trên một key, nó sẽ sử dụng một lệnh để khóa key đó.
  - Các client khác sẽ bị chặn lại (hoặc phải chờ) cho đến khi khóa được giải phóng.
- **Ví dụ**: Nếu một client đang cập nhật thông tin người dùng, nó sẽ khóa tài khoản đó. Trong thời gian khóa, không có client nào khác có thể thực hiện thao tác trên tài khoản đó cho đến khi thao tác đầu tiên hoàn tất và khóa được giải phóng.
- **Ưu điểm**: Đảm bảo tính toàn vẹn dữ liệu và ngăn chặn các xung đột trong truy cập.
- **Nhược điểm**: Có thể dẫn đến tình trạng chờ đợi lâu, giảm hiệu suất và tăng độ trễ.

### 2. **Khóa Lạc Quan (Optimistic Locking)**
- **Khái niệm**: Đây là phương pháp giả định rằng các xung đột hiếm xảy ra, vì vậy không khóa tài nguyên ngay lập tức mà chỉ kiểm tra tính nhất quán khi cập nhật.
- **Cách hoạt động**:
  - Khi một client muốn cập nhật một key, nó thực hiện thao tác mà không khóa key đó.
  - Trước khi xác nhận thay đổi, client kiểm tra xem giá trị của key có còn nguyên vẹn không (thông thường sử dụng một phiên bản hoặc một giá trị thời gian).
  - Nếu giá trị không thay đổi, giao dịch sẽ được thực hiện; nếu đã thay đổi, client sẽ từ chối thay đổi và có thể thử lại.
- **Ví dụ**: Một client đọc giá trị của một key, thực hiện một số tính toán dựa trên giá trị đó, và khi muốn cập nhật, nó sẽ so sánh giá trị hiện tại của key với giá trị mà nó đã đọc ban đầu. Nếu giá trị hiện tại không thay đổi, nó sẽ cập nhật; nếu đã thay đổi, nó sẽ từ chối cập nhật.
- **Ưu điểm**: Tăng hiệu suất vì không có khóa; tốt cho các tình huống có ít xung đột.
- **Nhược điểm**: Có thể dẫn đến thất bại trong việc cập nhật nếu có nhiều client cố gắng thay đổi cùng một key.

### Tóm tắt
- **Khóa Bi Quan**: Khóa tài nguyên ngay lập tức, ngăn chặn các giao dịch khác truy cập; bảo vệ dữ liệu nhưng có thể giảm hiệu suất.
- **Khóa Lạc Quan**: Không khóa tài nguyên ngay lập tức, kiểm tra trước khi cập nhật; tốt cho hiệu suất nhưng có thể gặp xung đột nếu có nhiều client cùng truy cập.

Redis chủ yếu sử dụng phương pháp **khóa lạc quan** thông qua cơ chế **WATCH** để quản lý đồng bộ hóa và tránh xung đột trong môi trường có nhiều client.