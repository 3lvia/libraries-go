// Package mschema provides support for working with message schemas (the 'm' in mschema stands for 'message'). This
// package is a wrapper around the Confluent Schema Registry API, but future versions may support other use cases.
//
// The package is configured using the Options pattern. The following options are available:
//  1. WithUser. Is used when communicating with the schema registry. The user and password are used to authenticate
//     with the schema registry.
//  2. WithClient. Is used when communicating with the schema registry. If not set, the default HTTP client will be
//     used. This option is useful for testing.
//  3. WithOtelTracerName. This option can be used to set the name of the OpenTelemetry tracer to use when creating
//     spans. If no name is set the tracer name "go.opentelemetry.io/otel" is used.
//
// Example usage:
// ```
//
//	 import (
//			"context"
//			"github.com/3lvia/libraries-go/pkg/mschema"
//		)
//
// ctx := context.Background()
// url := "https://schema-registry.example.com"
// topic := "my-topic"
//
// schemaRegistry, err := mschema.NewSchemaRegistry(url)
// descriptor, err := schemaRegistry.GetBySubject(ctx, topic)
//
//	if err != nil {
//		// handle error
//	}
//
// _ = descriptor
// ```
package mschema
