package auth

import (
	"context"
	"net/http"
	"strings"
)

// we need a type for context key
type ctxKey string

var (
	// tokenClaimsKey is the key in the request
	// context object for auth token claims
	tokenClaimsKey = ctxKey("claims")
)

// Wrap wraps an HTTP handler function, and checks both the
//
// Populates the access token claims object in the req ctx. Accessible through
// the GetClaims() function
func (a *Authenticator) Wrap(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := ExtractToken(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("no auth token provided"))
			return
		}

		verifiedClaims, err := a.ValidateJWT(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid auth token"))
			return
		}

		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), tokenClaimsKey, verifiedClaims)))
	})
}

// ExtractToken checks both the 'Authorization' HTTP header and
// the 'token' URL query parameter for an authentication token
func ExtractToken(r *http.Request) (string, bool) {
	token, ok := tokenFromAuthorizationHeader(r)
	if ok {
		return token, ok
	}
	token, ok = tokenFromURLQueryParameter(r)
	if ok {
		return token, ok
	}
	return "", false
}

func tokenFromAuthorizationHeader(r *http.Request) (string, bool) {
	jwt := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if jwt != "" {
		return jwt, true
	}
	return "", false
}

func tokenFromURLQueryParameter(r *http.Request) (string, bool) {
	token, ok := r.URL.Query()["token"]
	if !ok || len(token[0]) < 1 {
		return "", false
	}
	return token[0], true
}

// GetClaims returns the claims in a context object
func GetClaims(r *http.Request) *CustomClaims {
	return r.Context().Value(tokenClaimsKey).(*CustomClaims)
}
