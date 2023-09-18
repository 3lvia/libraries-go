package cmd

import (
	"context"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"log"
	"os"
)

func getRegistry(ctx context.Context) (mschema.Registry, <-chan error, error) {
	var secrets hashivault.SecretsManager
	var err error
	var errChan <-chan error

	vaultURL := url
	if vaultURL == "" {
		vaultURL = os.Getenv("VAULT_ADDR")
	}
	fmt.Printf("Vault URL: %s\n", vaultURL)

	token := vaultToken

	if token != "" {
		secrets, errChan, err = hashivault.New(ctx, hashivault.WithVaultAddress(vaultURL), hashivault.WithVaultToken(token))
	} else {
		secrets, errChan, err = hashivault.New(ctx, hashivault.WithVaultAddress(vaultURL), hashivault.WithOIDC())
	}
	if err != nil {
		return nil, nil, err
	}
	infoSecret, err := secrets.GetSecret(ctx, pathSecretSchemaRegistry)
	if err != nil {
		return nil, nil, err
	}
	m := infoSecret()
	rURL := infoSecret()["schema-registry-url"].(string)
	fmt.Printf("Schema registry URL: %s\n", rURL)

	p := fmt.Sprintf(pathPatternRegistry, system)
	fmt.Printf("Path: %s\n", p)
	systemSecret, err := secrets.GetSecret(ctx, p)
	if err != nil {
		log.Fatal(err)
	}
	m = systemSecret()
	registryKey := m["schema_registry_key"].(string)
	registrySecret := m["schema_registry_secret"].(string)

	registry, err := mschema.New(rURL, mschema.WithUser(registryKey, registrySecret))

	if err != nil {
		return nil, nil, err
	}

	return registry, errChan, nil
}
