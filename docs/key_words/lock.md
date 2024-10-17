## 1. Optimistic Lock và Pessimistic Lock

### Pessimistic Lock (Khóa bi quan)

- **Khái niệm**: Là cơ chế kiểm soát đồng thời dựa trên việc khóa dữ liệu trước khi cập nhật, nhằm tránh xung đột giữa các transaction.
- **Cách hoạt động**:
  - Khi một transaction (T1) bắt đầu và cần thay đổi dữ liệu, nó sẽ khóa hàng, trang, hoặc bảng tùy theo điều kiện.
  - Các transaction khác (T2, T3, ...) không thể thực hiện thay đổi cho đến khi T1 hoàn thành.
- **Ưu điểm**: Đảm bảo không có xung đột dữ liệu giữa các transaction.
- **Nhược điểm**: Gây ra tình trạng chờ đợi cho các transaction sau và có thể dẫn đến deadlock.
- **Ví dụ**: Giả sử bạn có một bảng `accounts` và hai transaction muốn cập nhật số dư tài khoản của cùng một người.
- **Cú pháp**
    ```sql
    -- Transaction T1
    BEGIN; -- Bắt đầu transaction
    SELECT * FROM accounts WHERE account_id = 1 FOR UPDATE; -- Khóa hàng để cập nhật
    UPDATE accounts SET balance = balance + 100 WHERE account_id = 1;
    COMMIT; -- Kết thúc transaction

    -- Transaction T2
    BEGIN;
    SELECT * FROM accounts WHERE account_id = 1 FOR UPDATE; -- T2 sẽ phải chờ T1 hoàn thành
    UPDATE accounts SET balance = balance - 50 WHERE account_id = 1;
    COMMIT;
    ```

### Optimistic Lock (Khóa lạc quan)

- **Khái niệm**: Cơ chế này cho phép nhiều transaction đọc dữ liệu mà không bị chặn và chỉ kiểm tra xung đột khi cố gắng cập nhật.
- **Cách hoạt động**:
  - Tất cả transaction đều có thể truy cập và cập nhật dữ liệu, nhưng khi commit, nó sẽ kiểm tra version của dữ liệu.
  - Nếu version cũ khớp với version hiện tại, transaction sẽ được thực hiện. Nếu không, transaction sẽ bị từ chối.
- **Ưu điểm**: Giảm thiểu tình trạng chờ đợi, tối ưu hiệu suất.
- **Nhược điểm**: Có thể dẫn đến việc cập nhật không thành công do xung đột version và yêu cầu retry.
- **Ví dụ**: Giả sử bạn có một bảng `products` với cột `version` để theo dõi phiên bản.
- **Cú pháp**
    ```sql
    BEGIN; -- Bắt đầu transaction

    -- Transaction T1
    SELECT id, version FROM products WHERE id = 1; -- Lấy dữ liệu và phiên bản hiện tại
    -- Giả sử version = 3
    UPDATE products SET price = price + 10, version = version + 1 WHERE id = 1 AND version = 3; -- Cập nhật chỉ khi phiên bản khớp

    IF ROW_COUNT() = 0 THEN
        -- Xử lý xung đột (ví dụ: throw an error or retry)
    END IF;

    COMMIT; -- Kết thúc transaction
    ```

## 2. Exclusive Lock và Shared Lock

### Shared Lock  (Khóa chia sẻ)

- **Khái niệm**: Còn gọi là read lock, được sử dụng khi một transaction muốn đọc dữ liệu.
- **Cách hoạt động**:
  - Cho phép nhiều transaction cùng đọc dữ liệu mà không gây xung đột.
  - Tuy nhiên, không cho phép bất kỳ transaction nào khác thực hiện thay đổi dữ liệu (write) trong khoảng thời gian đó.
- **Mục tiêu**: Đảm bảo tính toàn vẹn dữ liệu trong quá trình đọc.
- **Ví dụ**: Giả sử bạn có bảng `employees` và muốn cho phép nhiều transaction cùng đọc thông tin nhân viên.
- **Cú pháp**
    ```sql
    BEGIN; -- Bắt đầu transaction

    -- Transaction T1
    SELECT * FROM employees WHERE department = 'Sales' WITH SHARE LOCK; -- Đặt shared lock để đọc
    -- T1 có thể thực hiện các truy vấn đọc khác

    COMMIT; -- Kết thúc transaction

    -- Transaction T2
    BEGIN;
    SELECT * FROM employees WHERE department = 'Sales' WITH SHARE LOCK; -- T2 cũng có thể đọc cùng lúc
    COMMIT;
    ```

### Exclusive Lock (Khóa độc quyền)

- **Khái niệm**: Còn gọi là read-write lock, được sử dụng khi một transaction muốn thay đổi dữ liệu.
- **Cách hoạt động**:
  - Chỉ cho phép một transaction chiếm giữ exclusive lock tại một thời điểm, ngăn chặn mọi transaction khác đọc hoặc ghi dữ liệu.
- **Mục tiêu**: Đảm bảo không có sự can thiệp nào trong quá trình cập nhật dữ liệu.
- **Ví dụ**: Giả sử bạn muốn cập nhật thông tin cho một nhân viên mà không cho phép bất kỳ transaction nào khác can thiệp vào quá trình này.
- **Cú pháp**
    ```sql
    BEGIN; -- Bắt đầu transaction

    -- Transaction T1
    SELECT * FROM employees WHERE id = 1 FOR UPDATE; -- Đặt exclusive lock
    UPDATE employees SET salary = salary + 1000 WHERE id = 1;

    COMMIT; -- Kết thúc transaction

    -- Transaction T2
    BEGIN;
    SELECT * FROM employees WHERE id = 1 FOR UPDATE; -- T2 sẽ phải chờ cho đến khi T1 hoàn thành
    UPDATE employees SET salary = salary - 500 WHERE id = 1;

    COMMIT;
    ```

## Tóm tắt

- **Pessimistic Lock** phù hợp cho các trường hợp có khả năng xung đột cao, trong khi **Optimistic Lock** thích hợp cho các trường hợp xung đột thấp.
- **Shared Lock** cho phép nhiều transaction đọc dữ liệu cùng lúc nhưng không cho phép thay đổi, trong khi **Exclusive Lock** ngăn cản mọi thao tác khác khi một transaction đang thực hiện thay đổi dữ liệu. 

---------------------

## So sánh giữa `FOR UPDATE` và `FOR NO KEY UPDATE`

#### **1. `FOR UPDATE`**

- **Khóa toàn bộ hàng**: Khi sử dụng **`FOR UPDATE`**, hàng được khóa hoàn toàn, không cho phép các giao dịch khác đọc hoặc ghi vào hàng đó cho đến khi giao dịch hiện tại hoàn thành.
- **Ngăn cản giao dịch khác đọc và ghi**: Không chỉ ngăn các giao dịch khác cập nhật hàng bị khóa, mà còn ngăn họ đọc hàng này nếu giao dịch đang thực hiện một cập nhật.
- **Thường được sử dụng khi cần cập nhật dữ liệu**: Sử dụng trong các trường hợp bạn muốn chắc chắn rằng không có giao dịch nào khác có thể sửa đổi dữ liệu trong khi giao dịch của bạn chưa hoàn thành, như khi cập nhật dữ liệu trong bảng có liên quan đến khóa chính.

#### **Ví dụ `FOR UPDATE`**

```sql
BEGIN;

-- Lấy dòng và khóa toàn bộ hàng để cập nhật
SELECT balance FROM accounts WHERE account_id = 1 FOR UPDATE;

-- Tiến hành các thao tác cập nhật khác mà không lo ngại về sự can thiệp từ giao dịch khác
UPDATE accounts SET balance = balance - 100 WHERE account_id = 1;

COMMIT;
```
Trong ví dụ này, không có giao dịch nào khác có thể đọc hoặc ghi vào hàng `account_id = 1` cho đến khi giao dịch này được commit.

#### **2. `FOR NO KEY UPDATE`**

- **Khóa hàng, nhưng không khóa khóa chính (primary key)**: `FOR NO KEY UPDATE` chỉ khóa hàng để ngăn chặn việc cập nhật dữ liệu trên hàng đó, nhưng không ngăn cản các cập nhật liên quan đến khóa chính. Cập nhật liên quan đến khóa chính là những thay đổi trực tiếp đến giá trị của khóa chính (như thay đổi id trong bảng) hoặc các cập nhật cascade ảnh hưởng đến các khóa ngoại trong bảng khác liên quan đến khóa chính đó.
- **Cho phép các giao dịch khác đọc dữ liệu**: Các giao dịch khác vẫn có thể đọc hàng đó hoặc thậm chí khóa các hàng khác trong bảng mà không bị ảnh hưởng.
- **Giảm khả năng xung đột và deadlock**: `FOR NO KEY UPDATE` ít chặt chẽ hơn, do đó sẽ ít khả năng gây ra deadlock hơn và tăng hiệu suất khi có nhiều giao dịch cùng truy cập vào bảng.
    + Cho phép các giao dịch khác đọc hoặc khóa các khóa chính (primary key) của các hàng mà nó đang thao tác,
    + Nhưng vẫn khóa các khóa không phải khóa chính để ngăn chặn các giao dịch khác cập nhật hoặc xóa các hàng này.

#### **Ví dụ `FOR NO KEY UPDATE`**

```sql
BEGIN;

-- Lấy dòng và khóa hàng nhưng không ngăn cản các giao dịch khác đọc hoặc khóa khóa chính
SELECT balance FROM accounts WHERE account_id = 1 FOR NO KEY UPDATE;

-- Tiến hành các thao tác khác, chẳng hạn như cập nhật balance
UPDATE accounts SET balance = balance - 100 WHERE account_id = 1;

COMMIT;
```
Trong ví dụ này, giao dịch khác vẫn có thể thực hiện các thao tác đọc dữ liệu từ hàng này hoặc cập nhật trên các cột khác không liên quan đến khóa chính.

### **Tóm tắt sự khác biệt**

| Đặc điểm                     | `FOR UPDATE`                                 | `FOR NO KEY UPDATE`                            |
|------------------------------|----------------------------------------------|------------------------------------------------|
| **Loại khóa**                | Khóa toàn bộ hàng                            | Khóa hàng, nhưng không khóa khóa chính         |
| **Tác động đến đọc ghi**      | Ngăn các giao dịch khác đọc và ghi vào hàng  | Cho phép các giao dịch khác đọc, nhưng không cho phép cập nhật |
| **Xung đột và deadlock**      | Dễ gây deadlock hơn vì khóa toàn bộ hàng     | Ít gây deadlock, giảm khả năng xung đột        |
| **Tình huống sử dụng**        | Khi cần đảm bảo rằng hàng không bị thay đổi trong quá trình giao dịch | Khi cần cập nhật hàng mà không thay đổi khóa chính và giảm thiểu khóa chặn|

### **Khi nào nên dùng loại nào?**
- **`FOR UPDATE`**: Sử dụng khi bạn cần đảm bảo rằng không có giao dịch khác có thể thay đổi hoặc đọc dữ liệu trong hàng bạn đang cập nhật, đặc biệt khi liên quan đến thay đổi khóa chính.
- **`FOR NO KEY UPDATE`**: Sử dụng khi bạn chỉ cần khóa hàng để cập nhật nhưng không muốn ảnh hưởng đến việc đọc hoặc các thay đổi liên quan đến khóa chính.

Cách sử dụng `FOR NO KEY UPDATE` có thể giúp tăng cường khả năng đồng thời (concurrency) cho hệ thống của bạn trong các tình huống yêu cầu ít khóa chặn hơn.
