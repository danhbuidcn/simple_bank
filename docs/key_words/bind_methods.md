# Gin Binding Methods

Gin là một framework web mạnh mẽ cho Go, cung cấp nhiều phương thức để bind dữ liệu từ các request vào các struct. Dưới đây là một số phương thức phổ biến và cách sử dụng chúng.

## ShouldBindJSON

Phương thức `ShouldBindJSON` được sử dụng để bind dữ liệu JSON từ body của request vào một struct.

### Cách sử dụng

```go
type Account struct {
    Owner    string `json:"owner" binding:"required"`
    Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func createAccount(ctx *gin.Context) {
    var req Account
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Xử lý logic tạo tài khoản
}
```

### Giải thích

- `ctx.ShouldBindJSON(&req)`: Bind dữ liệu JSON từ body của request vào struct `req`.
- Nếu có lỗi trong quá trình bind, trả về mã lỗi `400 Bad Request`.

## ShouldBindUri

Phương thức `ShouldBindUri` được sử dụng để bind dữ liệu từ URI vào một struct.

### Cách sử dụng

```go
type AccountUri struct {
    ID uint `uri:"id" binding:"required"`
}

func getAccount(ctx *gin.Context) {
    var uri AccountUri
    if err := ctx.ShouldBindUri(&uri); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Xử lý logic lấy thông tin tài khoản
}
```

### Giải thích

- `ctx.ShouldBindUri(&uri)`: Bind dữ liệu từ URI vào struct `uri`.
- Nếu có lỗi trong quá trình bind, trả về mã lỗi `400 Bad Request`.

## ShouldBindQuery

Phương thức `ShouldBindQuery` được sử dụng để bind dữ liệu từ query string vào một struct.

### Cách sử dụng

```go
type AccountQuery struct {
    Page  int `form:"page" binding:"required"`
    Limit int `form:"limit" binding:"required"`
}

func listAccounts(ctx *gin.Context) {
    var query AccountQuery
    if err := ctx.ShouldBindQuery(&query); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Xử lý logic liệt kê tài khoản
}
```

### Giải thích

- `ctx.ShouldBindQuery(&query)`: Bind dữ liệu từ query string vào struct `query`.
- Nếu có lỗi trong quá trình bind, trả về mã lỗi `400 Bad Request`.

## ShouldBindHeader

Phương thức `ShouldBindHeader` được sử dụng để bind dữ liệu từ header của request vào một struct.

### Cách sử dụng

```go
type Header struct {
    Authorization string `header:"Authorization" binding:"required"`
}

func someHandler(ctx *gin.Context) {
    var header Header
    if err := ctx.ShouldBindHeader(&header); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Xử lý logic với dữ liệu từ header
}
```

### Giải thích

- `ctx.ShouldBindHeader(&header)`: Bind dữ liệu từ header của request vào struct `header`.
- Nếu có lỗi trong quá trình bind, trả về mã lỗi `400 Bad Request`.

## ShouldBind

Phương thức `ShouldBind` tự động xác định nguồn dữ liệu (JSON, query string, form, v.v.) và bind vào struct.

### Cách sử dụng

```go
type Account struct {
    Owner    string `json:"owner" form:"owner" binding:"required"`
    Currency string `json:"currency" form:"currency" binding:"required,oneof=USD EUR"`
}

func createAccount(ctx *gin.Context) {
    var req Account
    if err := ctx.ShouldBind(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Xử lý logic tạo tài khoản
}
```

### Giải thích

- `ctx.ShouldBind(&req)`: Tự động xác định nguồn dữ liệu và bind vào struct `req`.
- Nếu có lỗi trong quá trình bind, trả về mã lỗi `400 Bad Request`.

## Kết luận

- Gin cung cấp nhiều phương thức bind dữ liệu từ các nguồn khác nhau vào struct:
    - **ShouldBindJSON**:
        - **Công dụng**: Bind dữ liệu JSON từ body của request vào một struct.
        - **Ví dụ**: `ctx.ShouldBindJSON(&req)`
    - **ShouldBindUri**:
        - **Công dụng**: Bind dữ liệu từ URI vào một struct.
        - **Ví dụ**: `ctx.ShouldBindUri(&uri)`
    - **ShouldBindQuery**:
        - **Công dụng**: Bind dữ liệu từ query string vào một struct.
        - **Ví dụ**: `ctx.ShouldBindQuery(&query)`
    - **ShouldBindHeader**:
        - **Công dụng**: Bind dữ liệu từ header của request vào một struct.
        - **Ví dụ**: `ctx.ShouldBindHeader(&header)`
    - **ShouldBind**:
        - **Công dụng**: Tự động xác định nguồn dữ liệu (JSON, query string, form, v.v.) và bind vào struct.
        - **Ví dụ**: `ctx.ShouldBind(&req)`
