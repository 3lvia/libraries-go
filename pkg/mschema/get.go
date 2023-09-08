package mschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// GetById returns the schema descriptor for the given schema ID.
func getById(id int, url, user, password string, client *http.Client) (Descriptor, error) {
	fullURL := fmt.Sprintf("%s/schemas/ids/%d", url, id)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(user, password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var d descriptor
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}

	if d.ErrorCode != 0 {
		return nil, errors.New(d.Message)
	}

	return d, nil
}

// Get returns the schema descriptor for the latest version of the given topic.
func get(topic, url, user, password string, client *http.Client) (Descriptor, error) {
	fullURL := fmt.Sprintf("%s/subjects/%s-value/versions/latest", url, topic)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(user, password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var d descriptor
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}

	if d.ErrorCode != 0 {
		return nil, errors.New(d.Message)
	}

	return d, nil
}

//// List returns all the schemas that are registered in the schema registry.
//func List() ([]Descriptor, error) {
//	url := fmt.Sprintf("%s/schemas?deleted=false&latestOnly=true", schemaRegistryURL)
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	req.SetBasicAuth(user, password)
//
//	resp, err := currentHTTPClient.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	var descriptors []descriptor
//	err = json.Unmarshal(body, &descriptors)
//	if err != nil {
//		return nil, err
//	}
//
//	var res []Descriptor
//	for _, d := range descriptors {
//		res = append(res, d)
//	}
//	return res, nil
//}
