package api

import (
	"database/sql"
	"fmt"
	"net/http"
	db "simple_bank/db/sqlc"
	"simple_bank/token"

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

	// Check if the source account is valid
	fromAccount, valid := server.validAccount(ctx, req.FromAccountID, req.Currency, req.Amount, true)
	if !valid {
		return
	}

	// Get authorization payload
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// Check if the source account belongs to the authenticated user
	if fromAccount.Owner != authPayload.Username {
		err := fmt.Errorf("account [%d] does not belong to the authenticated user", fromAccount.ID)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// Check if the destination account is valid
	_, valid = server.validAccount(ctx, req.ToAccountID, req.Currency, req.Amount, false)
	if !valid {
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

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string, amount int64, checkBalance bool) (db.Account, bool) {
	// Check if the account exists
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	// Check if the account currency is correct
	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}

	// Check if the account has sufficient balance if needed
	if checkBalance && account.Balance < amount {
		err := fmt.Errorf("account [%d] has insufficient balance: %d < %d", account.ID, account.Balance, amount)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}

	return account, true
}
