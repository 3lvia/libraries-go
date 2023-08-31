package elvid

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/3lvia/elvid-go/internal/tlsclient"
	"github.com/MicahParks/keyfunc/v2"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrGetJWKSFailed = errors.New("failed to create JWKS from resource at the given URL")
	ErrInvalidJWT    = errors.New("the JWT is invalid")
	ErrInvalidToken  = errors.New("the token is invalid")
)

// Client is the interface to interact with the ElvID IDP
type Client struct {
	httpClient *http.Client

	oidConfig *oidcConfig
	jwks      *keyfunc.JWKS
}

// New creates a new instance of Client, used for eg checking token authenticity.
func New(ctx context.Context, opts ...Option) (*Client, error) {
	config := newConfig(opts...)

	client := config.httpClient
	if client == nil {
		var err error
		client, err = tlsclient.New(config.cert)
		if err != nil {
			return nil, err
		}
	}

	oidConfig, err := fetchOIDConfig(ctx, client, config.address+config.discovery)
	if err != nil {
		return nil, err
	}

	options := keyfunc.Options{
		Ctx:    ctx,
		Client: client,
		RefreshErrorHandler: func(err error) {
			if config.errorHandler != nil {
				config.errorHandler(err)
			}
		},
		RefreshInterval:   time.Hour * 24,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	jwks, err := keyfunc.Get(oidConfig.JsonWebKeySetUri, options)
	if err != nil {
		return nil, errors.Join(ErrGetJWKSFailed, err)
	}

	return &Client{
		httpClient: client,
		oidConfig:  oidConfig,
		jwks:       jwks,
	}, nil
}

// Shutdown cleans up background resources.
func (elvid *Client) Shutdown() {
	elvid.jwks.EndBackground()
}

// AuthorizeRequest takes an incoming request on behalf of the service and extracts the token from the "Authorization" header.
// The token is then checked for authenticity, and then the claims of that token is verified against StandardClaims.
func (elvid *Client) AuthorizeRequest(r *http.Request) error {
	return elvid.AuthorizeRequestWithClaims(r, &StandardClaims{})
}

// AuthorizeRequestWithClaims takes an incoming request on behalf of the service and extracts the token from the "Authorization" header.
// The token is then checked for authenticity, and then the claims of that token is verified against the passed Claims.
func (elvid *Client) AuthorizeRequestWithClaims(r *http.Request, claims Claims) error {
	jwtB64 := r.Header.Get("Authorization")
	jwtB64 = strings.Replace(jwtB64, "Bearer ", "", 1)

	return elvid.AuthorizeWithClaims(jwtB64, claims)
}

// Authorize parses and checks the authenticity of the passed JSON Web Token (JWT).
// The claims of that token is verified against StandardClaims.
func (elvid *Client) Authorize(jwtB64 string) error {
	return elvid.AuthorizeWithClaims(jwtB64, &StandardClaims{})
}

// AuthorizeWithClaims parses and checks the authenticity of the passed JSON Web Token (JWT).
// The claims of that token is verified against the passed Claims.
func (elvid *Client) AuthorizeWithClaims(jwtB64 string, claims Claims) error {

	opts := []jwt.ParserOption{
		jwt.WithLeeway(time.Minute * 5),
		jwt.WithIssuedAt(),
	}

	token, err := jwt.ParseWithClaims(jwtB64, claims, elvid.jwks.Keyfunc, opts...)
	if err != nil {
		return errors.Join(ErrInvalidJWT, err)
	}

	if !token.Valid {
		return ErrInvalidToken
	}

	return nil
}