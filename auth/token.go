package auth

import (
	"fmt"
	"time"
)

// Token represents an Inter token
type Token struct {
	AccessToken string `json:"access_token"`
	Type        string `json:"token_type"`
	Scope       string `json:"scope"`
	ExpiresIn   int64  `json:"expires_in"`

	expiresAt time.Time `json:"-"`
}

// SetExpiresAt sets the expiration time
func (t *Token) SetExpiresAt(expiresAt time.Time) *Token {
	t.expiresAt = expiresAt
	return t
}

// Valid returns true if the token is still valid
func (t Token) Valid() bool {
	return t.expiresAt.After(time.Now())
}

// GetAccessToken returns the token
func (t Token) GetAccessToken() string {
	return t.AccessToken
}

// getAuthorization returns the authorization header
func (t Token) GetAuthorization() string {
	return fmt.Sprintf("%s %s", t.Type, t.AccessToken)
}
