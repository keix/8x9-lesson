package entity

import "errors"

var (
	ErrClientNotFound      = errors.New("client not found")
	ErrAuthCodeNotFound    = errors.New("auth code not found")
	ErrAccessTokenNotFound = errors.New("access token not found")
	ErrAccessTokenExpired  = errors.New("access token expired")
	ErrNonceAlreadySeen    = errors.New("nonce already seen")
)
