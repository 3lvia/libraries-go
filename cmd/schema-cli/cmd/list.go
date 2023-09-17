/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

const (
	defaultVaultURL          = "https://vault.dev-elvia.io/"
	pathSecretSchemaRegistry = "edna/kv/data/cloudevents/info"
	pathPatternRegistry      = "edna/kv/data/cloudevents/creds/%s"
)

// listCmd represents the list command
var (
	url           string
	system        string
	vaultToken    string
	storageFolder string

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			runList(cmd.Context())
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "url", defaultVaultURL, "Vault URL")
	rootCmd.PersistentFlags().StringVar(&system, "system", "", "System")
	rootCmd.PersistentFlags().StringVar(&vaultToken, "vaultToken", "", "Vault token")

	rootCmd.PersistentFlags().StringVar(&storageFolder, "storageFolder", "", "Storage folder")

	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runList(ctx context.Context) {
	if system == "" {
		log.Fatal("system must be set")
	}

	var secrets hashivault.SecretsManager
	var err error
	var errChan <-chan error

	token := vaultToken
	if token != "" {
		log.Printf("Using token from flag")
		if err := os.Setenv("VAULT_TOKEN", token); err != nil {
			log.Fatal(err)
		}
	}
	if token == "" {
		token = os.Getenv("VAULT_TOKEN")
	}

	vaultURL := url
	if vaultURL == "" {
		vaultURL = os.Getenv("VAULT_ADDR")
	}

	if token != "" {
		secrets, errChan, err = hashivault.New(ctx, hashivault.WithVaultAddress(vaultURL), hashivault.WithVaultToken(token))
	} else {
		secrets, errChan, err = hashivault.New(ctx, hashivault.WithVaultAddress(vaultURL), hashivault.WithOIDC())
	}

	if err != nil {
		log.Fatal(err)
	}
	go func(ch <-chan error) {
		for e := range ch {
			log.Fatal(e)
		}
	}(errChan)

	infoSecret, err := secrets.GetSecret(ctx, pathSecretSchemaRegistry)
	if err != nil {
		log.Fatal(err)
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
	fmt.Printf("M: %v\n", m)
	registryKey := m["schema_registry_key"].(string)
	registrySecret := m["schema_registry_secret"].(string)

	registry, err := mschema.New(rURL, mschema.WithUser(registryKey, registrySecret))
	descriptors, err := registry.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	store := storageFolder != ""
	for _, descriptor := range descriptors {
		fmt.Printf("%d\t%s\t%s\n", descriptor.ID(), descriptor.Subject(), mschema.TypeName(descriptor.Type()))
		if store {
			sf := storageFolder + "/avsc"
			if descriptor.Type() != mschema.AVRO {
				sf = storageFolder + "/json"
			}
			fn := path.Join(sf, descriptorFileName(descriptor))
			f, err := os.Create(fn)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			if _, err := f.WriteString(descriptor.Schema()); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func descriptorFileName(d mschema.Descriptor) string {
	sub := strings.Replace(d.Subject(), ".", "_", -1)
	suffix := "json"
	if d.Type() == mschema.AVRO {
		suffix = "avsc"
	}

	return fmt.Sprintf("%d_%s.%s", d.ID(), sub, suffix)
}
