package schemaclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"io/ioutil"
	"net/http"
)

const (
	secretSchemaRegistryURL       = "edna/kv/data/cloudevents/info"
	secretKeySchemaRegistru       = "schema-registry-url"
	secretPathPatternRegistry     = "edna/kv/data/cloudevents/creds/%s"
	secretKeySchemaRegistry       = "schema_registry_key"
	secretKeySchemaRegistrySecret = "schema_registry_secret"
)

func newRegistry(ctx context.Context, system, topic string, secrets hashivault.SecretsManager, client *http.Client) (SchemaRegistry, error) {
	secret, err := secrets.GetSecret(ctx, fmt.Sprintf(secretPathPatternRegistry, system))
	if err != nil {
		return nil, err
	}
	m := secret()
	registryKey := m[secretKeySchemaRegistry].(string)
	registrySecret := m[secretKeySchemaRegistrySecret].(string)

	secret, err = secrets.GetSecret(ctx, secretSchemaRegistryURL)
	if err != nil {
		return nil, err
	}
	url := secret()[secretKeySchemaRegistru].(string)

	r := &registry{
		topic:  topic,
		key:    registryKey,
		secret: registrySecret,
		url:    url,
		client: client,
	}

	err = r.Get()
	if err != nil {
		return nil, err
	}

	return r, nil
}

type registry struct {
	topic  string
	key    string
	secret string
	url    string
	client *http.Client
}

func (r *registry) Get() error {
	schemaRegistryURL := fmt.Sprintf("%s/subjects/%s-value/versions/latest", r.url, r.topic)

	// Create an HTTP request
	req, err := http.NewRequest("GET", schemaRegistryURL, nil)
	if err != nil {
		//fmt.Println("Error creating request:", err)
		return err
	}

	// Add authentication headers
	req.SetBasicAuth(r.key, r.secret)

	// Send the request
	resp, err := r.client.Do(req)
	if err != nil {
		//fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Schema Registry returned a non-OK status: %s\n", resp.Status)
		return err
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("Error reading response body:", err)
		return err
	}
	_ = body

	return nil
}

func (r *registry) subjects() ([]string, error) {

	// Replace these with your Confluent Cloud Schema Registry details

	//schemaRegistryURL := "https://<your-schema-registry>.confluent.cloud/subjects/<your-subject>/versions/latest"
	//schemaRegistryURL := fmt.Sprintf("%s/subjects/%s/versions/latest", r.url, "private.dp.edna.examples")

	schemaRegistryURL := fmt.Sprintf("%s/subjects", r.url)

	// Create an HTTP request
	req, err := http.NewRequest("GET", schemaRegistryURL, nil)
	if err != nil {
		//fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Add authentication headers
	req.SetBasicAuth(r.key, r.secret)

	// Send the request
	resp, err := r.client.Do(req)
	if err != nil {
		//fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Schema Registry returned a non-OK status: %s\n", resp.Status)
		return nil, err
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var subs []string
	if err = json.Unmarshal(body, &subs); err != nil {
		return nil, err
	}

	return subs, nil

}
