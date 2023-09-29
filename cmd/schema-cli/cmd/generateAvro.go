/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// generateAvroCmd represents the generateAvro command
var generateAvroCmd = &cobra.Command{
	Use:   "generateAvro",
	Short: "Generates go source code for one specified schema from the Confluent registry.",
	Long: `Generates go source code for one specified schema from the Confluent registry. The schema can be specified by 
either topic or schemaID. SchemaID takes precedence.

The generated code is written to the directory that is given through the option 'storageFolder', if no such option is
given, the current directory is used. 

Rather than generating the code directly, a go:generate comment is written to a file in the storage folder. This file
can be used to generate the code by running 'go generate' in the storage folder. Along with the go:generate comment, an
'avsc' file is written to the storage folder. This file contains the schema in avro format. This file is referenced in
the go:generate comment. The file name is <schemaID>_<version>.avsc.
`,

	Run: func(cmd *cobra.Command, args []string) {
		if system == "" {
			log.Fatal("system must be set")
		}
		if topic == "" && schemaID == 0 {
			log.Fatal("topic or schemaID must be set")
		}
		if storageFolder == "" {
			storageFolder = "./"
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

	subject := ""
	if topic != "" {
		subject = topic + "-value"
	}
	if schemaID != 0 {
		subject = ""
	}

	var d mschema.Descriptor
	ok := false
	for _, desc := range all {
		if schemaID != 0 && desc.ID() == schemaID {
			d = desc
			ok = true
			break
		}
		if desc.Subject() == subject {
			d = desc
			ok = true
			break
		}
	}
	if !ok {
		log.Fatalf("subject %s not found", subject)
	}

	schemaFile, err := storeSchema(d)

	if err != nil {
		log.Fatal(err)
	}

	fn := path.Join(storageFolder, "generateAvro.go")
	f, err := os.Create(fn)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = f.WriteString("package main\n\n"); err != nil {
		log.Fatal(err)
	}
	if _, err = f.WriteString(fmt.Sprintf("//go:generate $GOPATH/bin/gogen-avro -package main %s %s", storageFolder, schemaFile)); err != nil {
		log.Fatal(err)
	}

	//go:generate $GOPATH/bin/gogen-avro -package model model ./100112_1.avsc
	//
	//paths := pc.sortedPaths()
	//for _, path := range paths {
	//	if path == "" {
	//		continue
	//	}
	//	fmt.Fprintf(w, "//go:generate mkdir -p %s\n", path)
	//	//go:generate mkdir -p ./dp
	//}
	//for _, d := range avroSchemas {
	//	subjPath := subjectToPath(d.Subject())
	//	fmt.Fprintf(w, "//go:generate $GOPATH/bin/gogen-avro -containers %s ./avsc/%d_%d.avsc\n", subjPath, d.ID(), d.Version())
	//}
}

func storeSchema(d mschema.Descriptor) (string, error) {
	fn := path.Join(storageFolder, fmt.Sprintf("%d_%d.avsc", d.ID(), d.Version()))
	f, err := os.Create(fn)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := f.WriteString(d.Schema()); err != nil {
		return "", err
	}
	return fn, nil
}
