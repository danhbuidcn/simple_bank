package api

import (
	"database/sql"
	"net/http"
	db "simple_bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

// ListAccountsRequest defines the request payload for listAccounts handler
type ListAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// listAccounts returns a paginated list of accounts
func (server *Server) listAccounts(ctx *gin.Context) {
	// Parse request
	var req ListAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get accounts
	accounts, err := server.store.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Ensure accounts is not nil
	if accounts == nil {
		accounts = []db.Account{}
	}

	// Return response
	ctx.JSON(http.StatusOK, accounts)
}

// GetAccountRequest defines the request payload for getAccount handler
type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// getAccount returns an account by ID
func (server *Server) getAccount(ctx *gin.Context) {
	// Parse request
	var req GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get account by ID
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	ctx.JSON(http.StatusOK, account)
}

// CreateAccountRequest defines the request payload for createAccount handler
type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

// createAccount creates a new account
func (server *Server) createAccount(ctx *gin.Context) {
	// Parse request
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Create a new account
	account, err := server.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	ctx.JSON(http.StatusOK, account)
}
