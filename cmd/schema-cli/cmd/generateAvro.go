/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// generateAvroCmd represents the generateAvro command
var generateAvroCmd = &cobra.Command{
	Use:   "generateAvro",
	Short: "Generates go source code for all schemas in the Confluent registry that are of type AVRO.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if system == "" {
			log.Fatal("system must be set")
		}
		runGenerateAvro(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(generateAvroCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateAvroCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateAvroCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runGenerateAvro(ctx context.Context) {
	registry, errChan, err := getRegistry(ctx)
	if err != nil {
		log.Fatal(err)
	}
	go func(ch <-chan error) {
		for e := range ch {
			log.Fatal(e)
		}
	}(errChan)

	all, err := registry.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	pc := newPathComputer("")
	var avroSchemas []mschema.Descriptor
	for _, schema := range all {
		if schema.Type() == mschema.AVRO {
			avroSchemas = append(avroSchemas, schema)
			pc.addSubject(schema.Subject())
		}
	}

	w := bufio.NewWriter(os.Stdout)

	paths := pc.sortedPaths()
	for _, path := range paths {
		fmt.Fprintf(w, "//go:generate mkdir -p %s\n", path)
		//go:generate mkdir -p ./dp
	}
}
