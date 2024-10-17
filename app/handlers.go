package app

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Hàm HealthCheckHandler sửa đổi để tương thích với gin.HandlerFunc
func HealthCheckHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
