# Webhook vs WebSocket

## Webhook

Webhook là một phương thức cho phép một ứng dụng cung cấp thông tin thời gian thực cho các ứng dụng khác.
Thay vì yêu cầu ứng dụng khác phải liên tục kiểm tra (polling) để xem có dữ liệu mới hay không, webhook sẽ tự động gửi dữ liệu đến ứng dụng khác khi có sự kiện xảy ra.

### Cách hoạt động của Webhook

1. **Thiết lập Webhook**: Ứng dụng A (ứng dụng gửi thông báo) sẽ cung cấp một URL webhook mà ứng dụng B (ứng dụng nhận thông báo) phải cung cấp. URL này là nơi ứng dụng A sẽ gửi dữ liệu khi có sự kiện xảy ra.
   
2. **Sự kiện xảy ra**: Khi một sự kiện cụ thể xảy ra trong ứng dụng A (ví dụ: một đơn hàng mới được tạo), ứng dụng A sẽ gửi một yêu cầu HTTP POST đến URL webhook của ứng dụng B với dữ liệu liên quan đến sự kiện đó.

3. **Xử lý dữ liệu**: Ứng dụng B nhận được dữ liệu từ yêu cầu HTTP POST và xử lý nó theo cách mà nó đã được lập trình.

### Lợi ích của Webhook

- **Thời gian thực**: Webhook cung cấp dữ liệu ngay lập tức khi sự kiện xảy ra, không cần phải chờ đợi.
- **Hiệu quả**: Giảm tải cho hệ thống vì không cần phải liên tục kiểm tra dữ liệu mới.
- **Đơn giản**: Dễ dàng thiết lập và sử dụng với các yêu cầu HTTP đơn giản.

## WebSocket

WebSocket là một giao thức truyền thông máy tính, cung cấp một kênh giao tiếp hai chiều, toàn thời gian giữa máy khách và máy chủ qua một kết nối TCP duy nhất. 
WebSocket được thiết kế để hoạt động qua các cổng HTTP tiêu chuẩn (80 và 443) và tương thích với các proxy và tường lửa hiện có.

### Cách hoạt động của WebSocket

1. **Thiết lập kết nối**: Kết nối WebSocket bắt đầu với một yêu cầu HTTP từ máy khách đến máy chủ để thiết lập kết nối. Nếu máy chủ chấp nhận yêu cầu, nó sẽ trả về một phản hồi HTTP 101 Switching Protocols, và kết nối WebSocket được thiết lập.

2. **Giao tiếp hai chiều**: Sau khi kết nối được thiết lập, cả máy khách và máy chủ có thể gửi tin nhắn bất kỳ lúc nào mà không cần phải thiết lập lại kết nối. Điều này khác với HTTP, nơi mỗi yêu cầu từ máy khách phải được máy chủ phản hồi.

3. **Đóng kết nối**: Kết nối WebSocket có thể được đóng bởi bất kỳ bên nào (máy khách hoặc máy chủ) bằng cách gửi một thông báo đóng kết nối.

### Lợi ích của WebSocket

- **Thời gian thực**: WebSocket cung cấp giao tiếp thời gian thực, lý tưởng cho các ứng dụng như chat, thông báo, và các trò chơi trực tuyến.
- **Hiệu quả**: WebSocket sử dụng một kết nối duy nhất cho giao tiếp hai chiều, giảm tải cho hệ thống so với việc phải thiết lập nhiều kết nối HTTP.
- **Đơn giản**: Dễ dàng thiết lập và sử dụng với các thư viện hỗ trợ WebSocket có sẵn.

## So sánh Webhook và WebSocket

### Webhook
- **Cách hoạt động**: Webhook là một cơ chế thông báo một chiều. Khi một sự kiện xảy ra, máy chủ sẽ gửi một yêu cầu HTTP POST đến một URL được chỉ định trước.
- **Thời gian thực**: Gần như thời gian thực, nhưng có thể có độ trễ nhỏ do phụ thuộc vào HTTP.
- **Kết nối**: Không duy trì kết nối liên tục. Mỗi sự kiện sẽ tạo ra một yêu cầu HTTP riêng biệt.
- **Sử dụng**: Thường được sử dụng cho các sự kiện không yêu cầu phản hồi ngay lập tức, như thông báo thanh toán, cập nhật trạng thái đơn hàng, v.v.
- **Đơn giản**: Dễ dàng thiết lập và sử dụng với các yêu cầu HTTP đơn giản.

### WebSocket
- **Cách hoạt động**: WebSocket cung cấp một kênh giao tiếp hai chiều, liên tục giữa máy khách và máy chủ qua một kết nối TCP duy nhất.
- **Thời gian thực**: Thực sự thời gian thực, không có độ trễ do duy trì kết nối liên tục.
- **Kết nối**: Duy trì kết nối liên tục, cho phép giao tiếp hai chiều mà không cần thiết lập lại kết nối.
- **Sử dụng**: Thường được sử dụng cho các ứng dụng yêu cầu giao tiếp thời gian thực, như chat trực tuyến, thông báo thời gian thực, trò chơi trực tuyến, v.v.
- **Phức tạp hơn**: Cần thiết lập và quản lý kết nối liên tục, có thể phức tạp hơn so với webhook.

## Các phương thức tương tự

Ngoài Webhook và WebSocket, còn có một số phương thức khác để giao tiếp giữa các ứng dụng:

### Server-Sent Events (SSE)
- **Cách hoạt động**: SSE cho phép máy chủ gửi các cập nhật tự động đến máy khách qua một kết nối HTTP duy nhất.
- **Thời gian thực**: Gần như thời gian thực, nhưng chỉ hỗ trợ giao tiếp một chiều từ máy chủ đến máy khách.
- **Kết nối**: Duy trì kết nối liên tục, nhưng chỉ cho phép máy chủ gửi dữ liệu đến máy khách.
- **Sử dụng**: Thường được sử dụng cho các ứng dụng cần cập nhật liên tục từ máy chủ, như thông báo, cập nhật dữ liệu thời gian thực, v.v.

### Long Polling
- **Cách hoạt động**: Máy khách gửi một yêu cầu HTTP đến máy chủ và giữ kết nối mở cho đến khi máy chủ có dữ liệu để gửi lại. Sau khi nhận được dữ liệu, máy khách sẽ gửi một yêu cầu mới.
- **Thời gian thực**: Gần như thời gian thực, nhưng có thể có độ trễ do phải thiết lập lại kết nối sau mỗi lần nhận dữ liệu.
- **Kết nối**: Không duy trì kết nối liên tục, nhưng giữ kết nối mở trong thời gian dài.
- **Sử dụng**: Thường được sử dụng khi không thể sử dụng WebSocket hoặc SSE, nhưng vẫn cần cập nhật thời gian thực.

## Kết luận

- **Webhook**: Tốt cho các sự kiện không yêu cầu phản hồi ngay lập tức và không cần duy trì kết nối liên tục.
- **WebSocket**: Tốt cho các ứng dụng yêu cầu giao tiếp thời gian thực và hai chiều.
- **SSE**: Tốt cho các ứng dụng cần cập nhật liên tục từ máy chủ đến máy khách.
- **Long Polling**: Tốt cho các ứng dụng cần cập nhật thời gian thực nhưng không thể sử dụng WebSocket hoặc SSE.
