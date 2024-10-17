package app

import (
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/health", HealthCheckHandler) // Sử dụng hàm HealthCheckHandler đã được sửa đổi
    return r
}
