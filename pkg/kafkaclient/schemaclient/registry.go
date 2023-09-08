package schemaclient

//
//const (
//	secretSchemaRegistryURL       = "edna/kv/data/cloudevents/info"
//	secretKeySchemaRegistru       = "schema-registry-url"
//	secretPathPatternRegistry     = "edna/kv/data/cloudevents/creds/%s"
//	secretKeySchemaRegistry       = "schema_registry_key"
//	secretKeySchemaRegistrySecret = "schema_registry_secret"
//)
//
//func newRegistry(ctx context.Context, system string, secrets hashivault.SecretsManager, client *http.Client) (SchemaRegistry, error) {
//	secret, err := secrets.GetSecret(ctx, fmt.Sprintf(secretPathPatternRegistry, system))
//	if err != nil {
//		return nil, err
//	}
//	m := secret()
//	registryKey := m[secretKeySchemaRegistry].(string)
//	registrySecret := m[secretKeySchemaRegistrySecret].(string)
//
//	secret, err = secrets.GetSecret(ctx, secretSchemaRegistryURL)
//	if err != nil {
//		return nil, err
//	}
//	url := secret()[secretKeySchemaRegistru].(string)
//
//	r := &registry{
//		key:     registryKey,
//		secret:  registrySecret,
//		url:     url,
//		client:  client,
//		mux:     &sync.RWMutex{},
//		schemas: map[string]*Schema{},
//	}
//
//	return r, nil
//}
//
//type registry struct {
//	key    string
//	secret string
//	url    string
//	client *http.Client
//
//	schemas map[string]*Schema
//	mux     *sync.RWMutex
//}
//
//func (r *registry) Get(topic string) (*Schema, error) {
//	r.mux.RLock()
//	s, ok := r.schemas[topic]
//	r.mux.RUnlock()
//	if ok {
//		return s, nil
//	}
//
//	r.mux.Lock()
//	defer r.mux.Unlock()
//
//	s, err := r.getCore(topic)
//	if err != nil {
//		return nil, err
//	}
//
//	r.schemas[topic] = s
//
//	return s, nil
//}
//
//func (r *registry) getCore(topic string) (*Schema, error) {
//	schemaRegistryURL := fmt.Sprintf("%s/subjects/%s-value/versions/latest", r.url, topic)
//
//	// Create an HTTP request
//	req, err := http.NewRequest("GET", schemaRegistryURL, nil)
//	if err != nil {
//		//fmt.Println("Error creating request:", err)
//		return nil, err
//	}
//
//	// Add authentication headers
//	req.SetBasicAuth(r.key, r.secret)
//
//	// Send the request
//	resp, err := r.client.Do(req)
//	if err != nil {
//		//fmt.Println("Error sending request:", err)
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	// Check the response status code
//	if resp.StatusCode != http.StatusOK {
//		err = fmt.Errorf("Schema Registry returned a non-OK status: %s\n", resp.Status)
//		return nil, err
//	}
//
//	// Read the response body
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		//fmt.Println("Error reading response body:", err)
//		return nil, err
//	}
//
//	var s Schema
//	if err = json.Unmarshal(body, &s); err != nil {
//		return nil, err
//	}
//
//	return &s, nil
//}
//
//func (r *registry) subjects() ([]string, error) {
//
//	// Replace these with your Confluent Cloud Schema Registry details
//
//	//schemaRegistryURL := "https://<your-schema-registry>.confluent.cloud/subjects/<your-subject>/versions/latest"
//	//schemaRegistryURL := fmt.Sprintf("%s/subjects/%s/versions/latest", r.url, "private.dp.edna.examples")
//
//	schemaRegistryURL := fmt.Sprintf("%s/subjects", r.url)
//
//	// Create an HTTP request
//	req, err := http.NewRequest("GET", schemaRegistryURL, nil)
//	if err != nil {
//		//fmt.Println("Error creating request:", err)
//		return nil, err
//	}
//
//	// Add authentication headers
//	req.SetBasicAuth(r.key, r.secret)
//
//	// Send the request
//	resp, err := r.client.Do(req)
//	if err != nil {
//		//fmt.Println("Error sending request:", err)
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	// Check the response status code
//	if resp.StatusCode != http.StatusOK {
//		err = fmt.Errorf("Schema Registry returned a non-OK status: %s\n", resp.Status)
//		return nil, err
//	}
//
//	// Read the response body
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		//fmt.Println("Error reading response body:", err)
//		return nil, err
//	}
//
//	var subs []string
//	if err = json.Unmarshal(body, &subs); err != nil {
//		return nil, err
//	}
//
//	return subs, nil
//
//}
