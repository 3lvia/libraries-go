/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strings"
)

// generateCmd represents the generate command
var (
	generateCmd = &cobra.Command{
		Use:   "generate",
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
			if topic == "" {
				log.Fatal("topic must be set")
			}
			runGenerate(cmd.Context())
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runGenerate(ctx context.Context) {
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

	subject := topic + "-value"

	var d mschema.Descriptor
	ok := false
	for _, v := range all {
		fmt.Printf("Subject: %s\n", v.Subject())
		if v.Subject() == subject {
			d = v
			ok = true
			break
		}
	}

	if !ok {
		log.Fatal("Subject not found")
	}

	gf, err := getGenerationFolder(d)
	if err != nil {
		log.Fatal(err)
	}

	if err := writeFolderGenerations(gf, os.Stdout); err != nil {
		log.Fatal(err)
	}

	dfn := descriptorFileName(d)
	schemaFolder := "./json"
	if d.Type() == mschema.AVRO {
		schemaFolder = "./avsc"
	}
	fmt.Printf("//go:generate $GOPATH/bin/gogen-avro -containers %s %s/%s\n", gf, schemaFolder, dfn)

	// //go:generate $GOPATH/bin/gogen-avro -containers ./avro ./avsc/100044_dp_examples-value.avsc
}

func writeFolderGenerations(folder string, w io.Writer) error {
	arr := strings.Split(folder, "/")
	accum := "."
	for _, a := range arr {
		if a == "" {
			continue
		}
		accum = accum + "/" + a
		line := fmt.Sprintf("//go:generate mkdir -p %s\n", accum)
		_, err := w.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	return nil
}

func getGenerationFolder(d mschema.Descriptor) (string, error) {
	arr := strings.Split(d.Subject(), ".")
	if len(arr) != 4 {
		return "", fmt.Errorf("invalid subject: %s", d.Subject())
	}
	typ := arr[len(arr)-1]
	if strings.Contains(typ, "-value") {
		typ = strings.Replace(typ, "-value", "", 1)
	}
	return fmt.Sprintf("%s/%s/%s/%s/", arr[1], arr[2], typ, arr[0]), nil
}
