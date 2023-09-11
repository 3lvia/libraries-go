package main

import (
	"fmt"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"log"
	"os"
)

const (
	secretSchemaRegistryURL   = "edna/kv/data/cloudevents/info"
	secretPathPatternRegistry = "edna/kv/data/cloudevents/creds/%s"
)

func main() {

}

func list() {
	ctx := context.Background()

	vaultAddr := "https://vault.dev-elvia.io"
	if err := os.Setenv("VAULT_ADDR", vaultAddr); err != nil {
		log.Fatal(err)
	}

	v, errChan, err := hashivault.New(
		ctx,
		hashivault.WithOIDC(),
		hashivault.WithVaultAddress(vaultAddr),
	)
	if err != nil {
		log.Fatal(err)
	}

	go func(ec <-chan error) {
		for err := range ec {
			log.Println(err)
		}
	}(errChan)

	secret, err := v.GetSecret(ctx, "edna/kv/data/cloudevents/info")
	if err != nil {
		log.Fatal(err)
	}
	mInfo := secret()
	registryURL := mInfo["schema-registry-url"].(string)

	system := "edna"

	secret, err = v.GetSecret(ctx, fmt.Sprintf(secretPathPatternRegistry, system))
	if err != nil {
		log.Fatal(err)
	}
	m := secret()
	registryKey := m["schema_registry_key"].(string)
	registrySecret := m["schema_registry_secret"].(string)

	registry, err := mschema.New(registryURL, mschema.WithUser(registryKey, registrySecret))

	l, err := registry.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range l {
		fmt.Println(d.Subject())
	}
}
