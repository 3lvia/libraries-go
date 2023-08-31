package elvid

import (
	"testing"

	"github.com/3lvia/libraries-go/pkg/elvid/testing/assert"
	"github.com/golang-jwt/jwt/v5"
)

func TestStandardClaims_Validate(t *testing.T) {
	tt := []struct {
		test    string
		wantErr error
		claims  StandardClaims
	}{
		{
			test:    "invalid httpClient id",
			wantErr: ErrEmptyClientID,
			claims: StandardClaims{
				ClientID: "",
			},
		},
		{
			test:    "valid claims",
			wantErr: nil,
			claims: StandardClaims{
				RegisteredClaims: jwt.RegisteredClaims{},
				Scope: []string{
					"openid",
				},
				ClientID:   "123456",
				ClientName: "test",
			},
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()

			gotErr := tc.claims.Validate()

			assert.Err(t, gotErr, tc.wantErr)
		})
	}
}