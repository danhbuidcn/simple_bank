package api

import (
	"net/http"
	db "simple_bank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server is the main struct for the API server
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.SetTrustedProxies(nil)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// Define routing
	router.GET("/health", server.healthCheck)

	// Define routing for accounts
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse creates a JSON error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// healthCheck handles the health check endpoint
func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}
