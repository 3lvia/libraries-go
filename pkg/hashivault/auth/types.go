package auth

import "time"

// Method is the authentication method to use when authenticating with Vault.
type Method int

const (
	_ Method = iota // skip 0
	// MethodGitHub is the authentication method where GitHub tokens are used to autenticate the user.
	MethodGitHub Method = iota

	// MethodK8s is the authentication method where Kubernetes service accounts are used to authenticate the user.
	MethodK8s

	// MethodOICD is the authentication method where OpenID Connect tokens are used to authenticate the user.
	MethodOICD

	// MethodToken is the authentication method where a Vault token has been obtained elsewhere and is used directly.
	MethodToken
)

func methodToString(m Method) string {
	switch m {
	case MethodGitHub:
		return "GitHub"
	case MethodK8s:
		return "Kubernetes"
	case MethodOICD:
		return "OIDC"
	case MethodToken:
		return "Token"
	default:
		return "Unknown"
	}
}

// gitToken holds github authentication information to be formatted to a bytes buffer
type gitToken struct {
	Token string `json:"token"`
}

// k8sToken holds kubernetes authentication information to be formatted to a bytes buffer
type k8sToken struct {
	JWT  string `json:"jwt"`
	Role string `json:"role"`
}

// AuthenticationResponse is the response from the Vault server after authentication.
type AuthenticationResponse interface {
	// ClientToken is the token to use when authenticating with Vault when fetching secrets.
	ClientToken() string

	// LeaseDurationSeconds is the duration of the token in seconds. This is the verbatim value returned by Vault.
	LeaseDurationSeconds() int

	// Renewable is true if the token is renewable.
	Renewable() bool

	// After returns a channel that fire when the token is about to expire. The channel will fire when the token has
	// 30 seconds left to live.
	After() <-chan time.Time
}

type authenticationResponse struct {
	RequestID     string             `json:"request_id"`
	LeaseID       string             `json:"lease_id"`
	Renew         bool               `json:"renewable"`
	LeaseDuration int                `json:"lease_duration"`
	Data          interface{}        `json:"data"`
	WrapInfo      interface{}        `json:"wrap_info"`
	Warnings      interface{}        `json:"warnings"`
	Auth          authenticationData `json:"auth"`
}

func (a authenticationResponse) ClientToken() string {
	return a.Auth.ClientToken
}

func (a authenticationResponse) LeaseDurationSeconds() int {
	return a.Auth.LeaseDuration
}

func (a authenticationResponse) Renewable() bool {
	return a.Auth.Renewable
}

func (a authenticationResponse) After() <-chan time.Time {
	secs := a.Auth.LeaseDuration
	if secs == 0 {
		secs = 3600 * 24 * 365 // 1 year
	}
	if secs >= 60 {
		secs -= 30 // 30 seconds leeway
	} else {
		secs = 1 // 1 second leeway
	}
	return time.After(time.Duration(secs) * time.Second)
}

type authenticationData struct {
	ClientToken    string                 `json:"client_token"`
	Accessor       string                 `json:"accessor"`
	Policies       []string               `json:"policies"`
	TokenPolicies  []string               `json:"token_policies"`
	Metadata       map[string]interface{} `json:"metadata"`
	LeaseDuration  int                    `json:"lease_duration"`
	Renewable      bool                   `json:"renewable"`
	EntityID       string                 `json:"entity_id"`
	TokenType      string                 `json:"token_type"`
	Orphan         bool                   `json:"orphan"`
	MFARequirement interface{}            `json:"mfa_requirement"`
	NumUses        int                    `json:"num_uses"`
}
