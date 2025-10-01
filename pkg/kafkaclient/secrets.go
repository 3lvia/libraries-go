package kafkaclient

import (
	"context"
	"fmt"

	"github.com/3lvia/libraries-go/pkg/hashivault"
)

const (
	secretSchemaRegistryURL   = "edna/kv/data/cloudevents/info"
	secretPathPatternRegistry = "edna/kv/data/cloudevents/creds/%s"
)

type SecretsResolver interface {
	Get(ctx context.Context, system string) (*SecretConfigValues, error)
}

type K8sSecrets struct {
	secrets hashivault.SecretsManager
}

func (k K8sSecrets) Get(ctx context.Context, system string) (*SecretConfigValues, error) {
	secret, err := k.secrets.GetSecret(ctx, fmt.Sprintf(secretPathPatternRegistry, system))
	if err != nil {
		return nil, err
	}
	m := secret()

	infoSecret, err := k.secrets.GetSecret(ctx, secretSchemaRegistryURL)
	if err != nil {
		return nil, err
	}
	mInfo := infoSecret()

	cv := &SecretConfigValues{
		registryURL:    mInfo["schema-registry-url"].(string),
		registryKey:    m["schema_registry_key"].(string),
		registrySecret: m["schema_registry_secret"].(string),

		key:    m["key"].(string),
		secret: m["secret"].(string),

		bootstrapServer:    mInfo["bootstrap-server"].(string),
		environment:        mInfo["environment"].(string),
		kafkaEnvID:         mInfo["kafa-env-id"].(string),
		kafkaMainClusterID: mInfo["kafa-main-cluster-id"].(string),
	}

	return cv, nil
}

type SecretConfigValues struct {
	registryURL, registryKey, registrySecret string

	key, secret string

	bootstrapServer    string
	environment        string
	kafkaEnvID         string
	kafkaMainClusterID string
}
