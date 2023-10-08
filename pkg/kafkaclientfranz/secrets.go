package kafkaclientfranz

import (
	"context"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/hashivault"
)

const (
	secretSchemaRegistryURL   = "edna/kv/data/cloudevents/info"
	secretPathPatternRegistry = "edna/kv/data/cloudevents/creds/%s"
)

func getSecrets(ctx context.Context, system string, secrets hashivault.SecretsManager) (*secretConfigValues, error) {
	secret, err := secrets.GetSecret(ctx, fmt.Sprintf(secretPathPatternRegistry, system))
	if err != nil {
		return nil, err
	}
	m := secret()

	infoSecret, err := secrets.GetSecret(ctx, secretSchemaRegistryURL)
	if err != nil {
		return nil, err
	}
	mInfo := infoSecret()

	cv := &secretConfigValues{
		registryURL:    mInfo["schema-registry-url"].(string),
		registryKey:    m["schema_registry_key"].(string),
		registrySecret: m["schema_registry_secret"].(string),

		blobSasURI:                                     m["blobstore-sas-uri"].(string),
		exceptionInvocationConnectionString:            m["exceptioninvocations-connection-string"].(string),
		exceptionInvocationLowPriorityConnectionString: m["exceptioninvocations-low-priority-connection-string"].(string),
		exceptionInvocationFailOverConnectionString:    m["exceptioninvocations-low-priority-failed-over-connection-string"].(string),
		exceptionInvocationLowPriorityTopic:            m["exceptioninvocations-low-priority-topic"].(string),
		exceptionInvocationsTopic:                      m["exceptioninvocations-topic"].(string),

		key:    m["key"].(string),
		secret: m["secret"].(string),

		successfulInvocationsConnectionString:            m["successfulinvocations-connection-string"].(string),
		successfulInvocationsLowPriorityConnectionString: m["successfulinvocations-low-priority-connection-string"].(string),
		successfulInvocationsFailOverConnectionString:    m["successfulinvocations-low-priority-failed-over-connection-string"].(string),
		successfulInvocationsLowPriorityTopic:            m["successfulinvocations-low-priority-topic"].(string),
		successfulInvocationsTopic:                       m["successfulinvocations-topic"].(string),

		bootstrapServer:    mInfo["bootstrap-server"].(string),
		environment:        mInfo["environment"].(string),
		kafkaEnvID:         mInfo["kafa-env-id"].(string),
		kafkaMainClusterID: mInfo["kafa-main-cluster-id"].(string),
	}

	return cv, nil
}

type secretConfigValues struct {
	registryURL, registryKey, registrySecret string

	key, secret string

	blobSasURI string

	exceptionInvocationConnectionString            string
	exceptionInvocationLowPriorityConnectionString string
	exceptionInvocationFailOverConnectionString    string
	exceptionInvocationLowPriorityTopic            string
	exceptionInvocationsTopic                      string

	successfulInvocationsConnectionString            string
	successfulInvocationsLowPriorityConnectionString string
	successfulInvocationsFailOverConnectionString    string
	successfulInvocationsLowPriorityTopic            string
	successfulInvocationsTopic                       string

	bootstrapServer    string
	environment        string
	kafkaEnvID         string
	kafkaMainClusterID string
}
