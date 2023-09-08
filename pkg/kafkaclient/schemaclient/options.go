package schemaclient

//import (
//	"github.com/3lvia/libraries-go/pkg/hashivault"
//	"net/http"
//)
//
//type optionsCollector struct {
//	secrets hashivault.SecretsManager
//	client  *http.Client
//}
//
//// Option is a function that can be used to configure this package.
//type Option func(*optionsCollector)
//
//// WithSecretsManager sets the secrets manager to use when fetching secrets.
//func WithSecretsManager(secrets hashivault.SecretsManager) Option {
//	return func(o *optionsCollector) {
//		o.secrets = secrets
//	}
//}
//
//// WithHTTPClient sets the HTTP client to use when communicating with the schema registry.
//func WithHTTPClient(client *http.Client) Option {
//	return func(o *optionsCollector) {
//		o.client = client
//	}
//}
