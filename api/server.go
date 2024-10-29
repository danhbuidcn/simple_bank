package api

import (
	"fmt"
	"net/http"
	db "simple_bank/db/sqlc"
	"simple_bank/token"
	"simple_bank/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server is the main struct for the API server
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	// Create a new token maker
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	// Create a new server
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// Register custom validator for currency
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// Setup routing
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	// Create a new gin router
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Define routing
	router.GET("/health", server.healthCheck)

	// Define routing for users
	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.getUser)

	// Define routing for login
	router.POST("/users/login", server.loginUser)

	// Create a new group for authenticated routes
	authRouters := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// Define routing for accounts
	authRouters.POST("/accounts", server.createAccount)
	authRouters.GET("/accounts/:id", server.getAccount)
	authRouters.GET("/accounts", server.listAccounts)

	// Define routing for transfers
	authRouters.POST("/transfers", server.createTransfer)

	server.router = router
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
