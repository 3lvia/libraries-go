package elvid

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrEmptyClientID = errors.New("the clientid is empty")
)

// Claims is an interface that represents the claims fields in a JWT.
// There is a method Validate that is called during the authentication process.
type Claims interface {
	jwt.Claims
	Validate() error
}

// StandardClaims contains required claims in a JWT that is used for validation.
// It also contains claims that is standard for Client JWTs, eg Scope.
type StandardClaims struct {
	jwt.RegisteredClaims
	Scope      []string `json:"scope,omitempty"`
	ClientID   string   `json:"client_id,omitempty"`
	ClientName string   `json:"client_name,omitempty"`
}

// Validate validates the parsed claims in a JWT.
func (c StandardClaims) Validate() error {
	if c.ClientID == "" {
		return ErrEmptyClientID
	}

	return nil
}