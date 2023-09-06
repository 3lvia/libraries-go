package hashivault

import (
	"context"
	"net/http"
	"testing"
)

func Test_manager_SetDefaultGoogleCredentials(t *testing.T) {
	type fields struct {
		vaultAddress string
		client       *http.Client
		tokenGetter  tokenGetterFunc
		errChan      chan<- error
	}
	type args struct {
		path string
		key  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &manager{
				vaultAddress: tt.fields.vaultAddress,
				client:       tt.fields.client,
				tokenGetter:  tt.fields.tokenGetter,
				errChan:      tt.fields.errChan,
			}
			if err := m.SetDefaultGoogleCredentials(context.Background(), tt.args.path, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("SetDefaultGoogleCredentials() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
