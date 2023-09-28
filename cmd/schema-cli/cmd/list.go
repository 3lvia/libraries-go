/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/spf13/cobra"
	"log"
)

const (
	defaultVaultURL          = "https://vault.dev-elvia.io/"
	pathSecretSchemaRegistry = "edna/kv/data/cloudevents/info"
	pathPatternRegistry      = "edna/kv/data/cloudevents/creds/%s"
)

// listCmd represents the list command
var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			if system == "" {
				log.Fatal("system must be set")
			}
			runList(cmd.Context())
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList(ctx context.Context) {
	registry, errChan, err := getRegistry(ctx)

	go func(ch <-chan error) {
		for {
			err := <-ch
			if err != nil {
				log.Fatal(err)
			}
		}
	}(errChan)

	descriptors, err := registry.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//store := storageFolder != ""
	for _, descriptor := range descriptors {
		//fmt.Printf("%d\t%s\t%s\n", descriptor.ID(), descriptor.Subject(), mschema.TypeName(descriptor.Type()))
		fmt.Println(descriptor.Subject())
		//if store {
		//	sf := storageFolder + "/avsc"
		//	if descriptor.Type() != mschema.AVRO {
		//		sf = storageFolder + "/json"
		//	}
		//	fn := path.Join(sf, descriptorFileName(descriptor))
		//	f, err := os.Create(fn)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	defer f.Close()
		//	if _, err := f.WriteString(descriptor.Schema()); err != nil {
		//		log.Fatal(err)
		//	}
		//}
	}
}

func descriptorFileName(d mschema.Descriptor) string {
	suffix := "json"
	if d.Type() == mschema.AVRO {
		suffix = "avsc"
	}

	return fmt.Sprintf("%d_%d.%s", d.ID(), d.Version(), suffix)
}
