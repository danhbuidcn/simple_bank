package api

import (
	"database/sql"
	"errors"
	"net/http"
	db "simple_bank/db/sqlc"
	"simple_bank/token"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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

	// Get authorization payload
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	// Get accounts
	accounts, err := server.store.ListAccounts(ctx, db.ListAccountsParams{
		Owner:  authPayload.Username,
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

	// Get authorization payload
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account does not belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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

	// Get authorization payload
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	// Create a new account
	account, err := server.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    authPayload.Username,
		Balance:  0,
		Currency: req.Currency,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// Check if the error is a foreign key violation or unique violation
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	ctx.JSON(http.StatusOK, account)
}
