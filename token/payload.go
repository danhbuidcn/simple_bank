package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of errors returned by the VerifyToken function
var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload is the payload of the token
type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IssueAt  time.Time `json:"issue_at"`
	ExpireAt time.Time `json:"expire_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       tokenID,
		Username: username,
		IssueAt:  time.Now(),
		ExpireAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token is valid or not
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpireAt) {
		return ErrExpiredToken
	}
	return nil
}
