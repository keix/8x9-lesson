package entity

import "time"

// AuthCode represents an OAuth 2.0 authorization code
type AuthCode struct {
	Code                string    `json:"code"`
	ClientID            string    `json:"client_id"`
	UserID              string    `json:"user_id"`
	RedirectURI         string    `json:"redirect_uri"`
	CodeChallenge       string    `json:"code_challenge,omitempty"`
	CodeChallengeMethod string    `json:"code_challenge_method,omitempty"`
	Scope               string    `json:"scope"`
	State               string    `json:"state"`
	Nonce               string    `json:"nonce,omitempty"`
	ExpiresAt           time.Time `json:"expires_at"`
}
