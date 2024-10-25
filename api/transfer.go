package api

import (
	"net/http"
	db "simple_bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

// transferRequest defines the request payload for createTransfer handler
type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,min=1"`
	Currency      string `json:"currency" binding:"required,currency"`
}

// createTransfer creates a new transfer
func (server *Server) createTransfer(ctx *gin.Context) {
	// Parse request
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Create a new transfer
	transfer, err := server.store.TransferTx(ctx, db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Currency:      req.Currency,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	ctx.JSON(http.StatusOK, transfer)
}
