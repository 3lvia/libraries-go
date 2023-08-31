# elvid-go
Functionality for working with Elvid in Go.

# Usage

Create a new ElvID object by using `elvid.New`.

**NOTE** call the `Shutdown` method before the object goes out of scope.

```go
ctx := context.Background()

opts := []elvid.Option{
    elvid.WithAddress("https://elvid.elvia.io"),
    elvid.WithErrorHandler(func(err error) {
        log.Printf("There was an error: %s", err.Error())
    }),
}

elvID, err := elvid.New(ctx, opts...)
if err != nil {
    log.Fatalf("failed creating elvid: %v", err)
}
defer elvID.Shutdown()
```

## Authorizing a token

Use the `Authorize` method variants to check a token for authenticity against ElvID.

```go
token := "eyJh..."
err := elvID.Authorize(token)
if err != nil {
    log.Fatalf("authorization failed %v", err)
}
log.Println("authorization succeeded")
```

There are helper functions to authorize an HTTP request.
These functions check the `Authorization` header in the `http.Request` for a valid JWT.

```go
func httpHandler(w http.ResponseWriter, r *http.Request) {
    err := elvID.AuthorizeRequest(r)
    if err != nil {
        log.Fatalf("authorization failed %v", err)
    }
    log.Println("authorization succeeded")
}
```

## Extracting claims

To get the claims, use the `AuthorizeWithClaims` methods. You can use the `elvid.StandardClaims` object to get the most common, eg scopes.

```go
token := "eyJh..."

claims := &elvid.StandardClaims{}
err = elvID.AuthorizeWithClaims(token, claims)
if err != nil {
    log.Fatalf("authorization failed %v", err)
}
log.Println("authorization succeeded")

// use claims.Scope to validate required app scope(s)
hasScope := func(scopes []string, requiredScope string) bool {
    for _, s := range scopes {
        if s == requiredScope {
            return true
        }
    }
    return false
}(claims.Scope, "my-app.scope")

if !hasScope {
    log.Fatal("token did not contain the required scope")
}
```

If you need to extract and validate additional optional claims, you can implement the `elvid.Claims` interface.
Using the `Validate() error` method, you can do validation on the claims.


```go
type MyAppClaims struct {
	elvid.StandardClaims
	Name     string   `json:"name,omitempty"`
	Email    string   `json:"email,omitempty"`
	AdGroups []string `json:"ad_groups,omitempty"`
}

func (c MyAppClaims) Validate() error {

    if err := c.StandardClaims.Validate(); err != nil {
        return err
    }
    
    hasScope := func(scopes []string, requiredScope string) bool {
        for _, s := range scopes {
            if s == requiredScope {
                return true
            }
        }
        return false
    }(c.Scope, "my-app.scope")
    
    if !hasScope {
        return errors.New("token did not contain the required scope")
    }

    return nil
}
```

You can use this struct to pass into `AuthorizeWithClaims` methods

```go
token := "eyJh..."

claims := &MyAppClaims{}
err = elvID.AuthorizeWithClaims(token, claims)
if err != nil {
    log.Fatalf("authorization failed %v", err)
}
log.Println("authorization succeeded")

isAdmin := func(groups []string, id string) bool {
        for _, v := range groups {
            if v == id {
            return true
        }
    }
    return false
}(claims.AdGroups, "ad-group-id")

log.Println("user is admin:", isAdmin)

```



## Generated documentation
The source code in this repository is documented using godoc standards (https://tip.golang.org/doc/comment). The
documentation can be generated and viewed on the local development machine as follows:
0. Install godoc: `go install -v golang.org/x/tools/cmd/godoc@latest` (if not already installed)
1. Run `godoc -http=:6060` in the root of the project
2. Open a browser and navigate to `http://localhost:6060/pkg/github.com/3lvia/elvid-go/`

# Acquiring a token from ElvID

This library does not include a way to request an access token from ElvID.

To get one, use the [golang.org/x/oauth2](https://pkg.go.dev/golang.org/x/oauth2) package.

```go
ctx := context.Background()

config := &clientcredentials.Config{
    ClientID:     "client-id",
    ClientSecret: "client-secret",
    TokenURL:     "https://elvid.elvia.io/connect/token",
}

token, err := config.Token(ctx)
if err != nil {
    log.Fatalf("fetching token failed: %v", err)
}

fmt.Println(fmt.Sprintf("access token: %s", token.AccessToken))
```