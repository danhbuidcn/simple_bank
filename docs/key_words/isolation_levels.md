Bốn mức độ cô lập (isolation levels) theo chuẩn ANSI SQL là các mức độ kiểm soát tương tác giữa các transaction để đảm bảo tính nhất quán của dữ liệu. Các mức độ này là:

### 1. **Read Uncommitted**

- **Mô tả**: Đây là mức cô lập thấp nhất. Một transaction có thể đọc dữ liệu từ một transaction khác ngay cả khi transaction kia chưa hoàn thành (committed).
- **Vấn đề có thể xảy ra**:
  - **Dirty Read**: Transaction này đọc dữ liệu tạm thời từ một transaction khác chưa hoàn thành, dẫn đến nguy cơ đọc dữ liệu sai hoặc không chính xác.
  - **Non-repeatable Read**: Giá trị của một hàng có thể thay đổi khi cùng đọc nó ở hai thời điểm khác nhau trong cùng một transaction do transaction khác thay đổi và commit.
  - **Phantom Read**: Transaction có thể thấy một số hàng mới thêm vào hoặc thay đổi do transaction khác thêm/ sửa.
- **Cú pháp**:
  ```sql
  SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
  ```
- **Ví dụ**:
  ```sql
  -- Transaction 1 (Đang chỉnh sửa dữ liệu nhưng chưa commit)
  BEGIN;
  UPDATE accounts SET balance = balance - 100 WHERE account_id = 1;

  -- Transaction 2 (Có thể đọc dữ liệu chưa commit từ transaction 1)
  SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
  SELECT balance FROM accounts WHERE account_id = 1;  -- Có thể thấy thay đổi của Transaction 1 ngay lập tức
  
  -- Transaction 1 commit
  COMMIT;
  ```

  **Kết quả**: Transaction 2 có thể thấy thay đổi của Transaction 1 trước khi nó được commit, dẫn đến hiện tượng **Dirty Read**.

### 2. **Read Committed** (Mặc định trong nhiều hệ quản trị CSDL)

- **Mô tả**: Mức này ngăn chặn **Dirty Read** bằng cách đảm bảo rằng chỉ có dữ liệu đã được commit mới có thể được đọc.
- **Vấn đề có thể xảy ra**:
  - **Non-repeatable Read**: Giá trị của cùng một hàng có thể khác nhau nếu đọc lại sau khi một transaction khác đã cập nhật và commit hàng đó.
  - **Phantom Read**: Một transaction có thể thấy thêm hoặc bớt các hàng do một transaction khác commit sau khi transaction đầu tiên đã bắt đầu.
- **Cú pháp**:
  ```sql
  SET TRANSACTION ISOLATION LEVEL READ COMMITTED;
  ```
- **Ví dụ**:
  ```sql
  -- Transaction 1 (Đang chỉnh sửa dữ liệu nhưng chưa commit)
  BEGIN;
  UPDATE accounts SET balance = balance - 100 WHERE account_id = 1;

  -- Transaction 2 (Chỉ có thể đọc dữ liệu đã được commit)
  SET TRANSACTION ISOLATION LEVEL READ COMMITTED;
  SELECT balance FROM accounts WHERE account_id = 1;  -- Không thấy thay đổi của Transaction 1

  -- Transaction 1 commit
  COMMIT;

  -- Transaction 2 (Sau khi commit, có thể thấy thay đổi)
  SELECT balance FROM accounts WHERE account_id = 1;
  ```

  **Kết quả**: Transaction 2 chỉ có thể đọc dữ liệu đã được commit. Tuy nhiên, **Non-repeatable Read** có thể xảy ra nếu Transaction 2 đọc lại sau khi Transaction 1 commit.

### 3. **Repeatable Read**

- **Mô tả**: Ở mức này, **Non-repeatable Read** được ngăn chặn. Các hàng đã đọc trong transaction sẽ giữ nguyên, không bị ảnh hưởng bởi các transaction khác cho đến khi transaction hiện tại kết thúc.
- **Vấn đề có thể xảy ra**:
  - **Phantom Read**: Dữ liệu có thể thay đổi về số lượng (thêm hoặc bớt hàng) do các transaction khác thêm/xóa sau khi transaction đã bắt đầu.
- **Cú pháp**:
  ```sql
  SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;
  ```
- **Ví dụ**:
  ```sql
  -- Transaction 1 (Lấy giá trị của balance)
  BEGIN;
  SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;
  SELECT balance FROM accounts WHERE account_id = 1;

  -- Transaction 2 (Cập nhật balance và commit)
  BEGIN;
  UPDATE accounts SET balance = balance - 100 WHERE account_id = 1;
  COMMIT;

  -- Transaction 1 (Đọc lại cùng dòng, sẽ thấy giá trị không thay đổi)
  SELECT balance FROM accounts WHERE account_id = 1;  -- Không thấy thay đổi của Transaction 2

  -- Transaction 1 commit
  COMMIT;
  ```

  **Kết quả**: Trong **Repeatable Read**, Transaction 1 sẽ không thấy thay đổi của Transaction 2 cho đến khi nó kết thúc, tránh hiện tượng **Non-repeatable Read**. Tuy nhiên, **Phantom Read** vẫn có thể xảy ra nếu Transaction 2 thêm/xóa hàng mới.

### 4. **Serializable** (Mức cao nhất)

- **Mô tả**: Đây là mức độ cô lập chặt chẽ nhất, đảm bảo rằng các transaction hoàn toàn độc lập và tuần tự hóa. Tất cả các transaction chạy như thể chúng được thực hiện tuần tự, không có tác động qua lại nào.
- **Vấn đề có thể xảy ra**:
  - Không có vấn đề nào xảy ra (ngăn chặn tất cả các hiện tượng: **Dirty Read**, **Non-repeatable Read**, và **Phantom Read**), nhưng sẽ có hiệu suất thấp hơn vì cần phải khóa nhiều tài nguyên hơn.
- **Cú pháp**:
  ```sql
  SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
  ```
- **Ví dụ**:
  ```sql
  -- Transaction 1 (Đọc dữ liệu từ bảng accounts)
  BEGIN;
  SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
  SELECT balance FROM accounts WHERE account_id = 1;

  -- Transaction 2 (Cố gắng cập nhật balance)
  BEGIN;
  SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
  UPDATE accounts SET balance = balance - 100 WHERE account_id = 1;  -- Bị khóa chờ Transaction 1 hoàn thành

  -- Transaction 1 commit
  COMMIT;

  -- Transaction 2 sau khi Transaction 1 kết thúc
  UPDATE accounts SET balance = balance - 100 WHERE account_id = 1;
  COMMIT;
  ```

  **Kết quả**: Trong mức **Serializable**, các transaction sẽ thực hiện tuần tự. Nếu Transaction 2 cố gắng thực hiện thay đổi trong khi Transaction 1 chưa hoàn thành, nó sẽ phải chờ. Điều này tránh tất cả các hiện tượng **Dirty Read**, **Non-repeatable Read**, và **Phantom Read**.

### Tóm tắt các hiện tượng lỗi trong transaction:

- **Dirty Read**: Đọc dữ liệu chưa được commit từ transaction khác.
- **Non-repeatable Read**: Đọc cùng một hàng hai lần, và dữ liệu bị thay đổi do transaction khác giữa hai lần đọc.
- **Phantom Read**: Kết quả của câu truy vấn có thể khác nhau (thêm hoặc mất hàng) khi chạy lại do transaction khác thêm/xóa hàng.
- **serialization anomaly**: Kết quả không tương đương với thứ tự tuần tự của các transaction, gây ra sự không nhất quán trong dữ liệu.