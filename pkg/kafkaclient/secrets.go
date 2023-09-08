package kafkaclient

import (
	"context"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/hashivault"
)

const (
	secretSchemaRegistryURL       = "edna/kv/data/cloudevents/info"
	secretKeySchemaRegistru       = "schema-registry-url"
	secretPathPatternRegistry     = "edna/kv/data/cloudevents/creds/%s"
	secretKeySchemaRegistry       = "schema_registry_key"
	secretKeySchemaRegistrySecret = "schema_registry_secret"
)

func getSecrets(ctx context.Context, system string, secrets hashivault.SecretsManager) (registryKey, registrySecret, url string, err error) {
	secret, err := secrets.GetSecret(ctx, fmt.Sprintf(secretPathPatternRegistry, system))
	if err != nil {
		return
	}
	m := secret()
	registryKey = m[secretKeySchemaRegistry].(string)
	registrySecret = m[secretKeySchemaRegistrySecret].(string)

	secret, err = secrets.GetSecret(ctx, secretSchemaRegistryURL)
	if err != nil {
		return
	}
	url = secret()[secretKeySchemaRegistru].(string)
	return
}
