package auth

import (
	"crypto/rsa"
)

// Authenticator is the module in charge of authentication
type Authenticator struct {
	db     DB
	signer *rsa.PrivateKey
	iss    string
	aud    string
}

// NewAuthenticator is the Authenticator constructor
func NewAuthenticator(db DB, key *rsa.PrivateKey, aud, iss string) *Authenticator {
	a := &Authenticator{
		db:     db,
		signer: key,
		aud:    aud,
		iss:    iss,
	}
	return a
}
