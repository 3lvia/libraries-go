package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/3lvia/elvid-go/elvid"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	elvidBaseUrl       = "https://elvid.test-elvia.io"
	elvidTokenEndpoint = elvidBaseUrl + "/connect/token"

	clientId     = "client-id"
	clientSecret = "client-secret"
	clientScope  = "client.api"
)

func main() {
	ctx := context.Background()

	// Acquire an access token from ElvID
	// The app must be configured with machine client access to have a client id and client secret
	// Note: the client secret is sensitive information and should be loaded from a secrets vault
	config := &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     elvidTokenEndpoint,
	}

	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("fetching token failed: %v", err)
	}

	fmt.Println(fmt.Sprintf("access token: %s", token.AccessToken))

	// Create an ElvID client
	opts := []elvid.Option{
		elvid.WithAddress(elvidBaseUrl),
		elvid.WithErrorHandler(func(err error) {
			log.Printf("there was an error: %s", err.Error())
		}),
	}

	elvID, err := elvid.New(ctx, opts...)
	if err != nil {
		log.Fatalf("failed creating elvid: %v", err)
	}
	defer elvID.Shutdown()

	// Check the access token for authenticity
	// Use the MyAppClaims struct to add your own validation rules, eg ad groups
	claims := &MyAppClaims{}
	err = elvID.AuthorizeWithClaims(token.AccessToken, claims)
	if err != nil {
		log.Fatalf("authorization failed %v", err)
	}
	log.Println("authorization succeeded")

	// Use the claims loaded from the token to change behaviour
	isAdmin := hasAdGroup(claims.AdGroups, "ad-group-id")
	log.Println("user is admin:", isAdmin)
}

func hasAdGroup(groups []string, id string) bool {
	for _, v := range groups {
		if v == id {
			return true
		}
	}
	return false
}

type MyAppClaims struct {
	elvid.StandardClaims

	// Add further claims depending on your app requirements, for example:
	Name     string   `json:"name,omitempty"`
	Email    string   `json:"email,omitempty"`
	AdGroups []string `json:"ad_groups,omitempty"`
}

func (c MyAppClaims) Validate() error {

	if err := c.StandardClaims.Validate(); err != nil {
		return err
	}

	if c.ClientID != clientId {
		return errors.New("the client id is invalid")
	}

	hasScope := func(scopes []string, requiredScope string) bool {
		for _, s := range scopes {
			if s == requiredScope {
				return true
			}
		}
		return false
	}(c.Scope, clientScope)

	if !hasScope {
		return errors.New("token did not contain the required scope")
	}

	return nil
}