package mschema

import (
	"fmt"
	"net/http"
	"sync"
)

type registry struct {
	url      string
	user     string
	password string

	client *http.Client

	descriptors map[string]Descriptor
	mux         *sync.RWMutex
}

func (r *registry) GetByID(id int) (Descriptor, error) {
	r.mux.RLock()
	k := fmt.Sprintf("ID-%d", id)
	if d, ok := r.descriptors[k]; ok {
		r.mux.RUnlock()
		return d, nil
	}
	r.mux.RUnlock()

	d, err := getById(id, r.url, r.user, r.password, r.client)
	if err != nil {
		return nil, err
	}

	r.storeDescriptor(d)
	return d, nil
}

func (r *registry) GetBySubject(subject string) (Descriptor, error) {
	r.mux.RLock()
	if d, ok := r.descriptors[subject]; ok {
		r.mux.RUnlock()
		return d, nil
	}
	r.mux.RUnlock()

	d, err := get(subject, r.url, r.user, r.password, r.client)
	if err != nil {
		return nil, err
	}

	r.storeDescriptor(d)
	return d, nil
}

func (r *registry) storeDescriptor(d Descriptor) {
	r.mux.Lock()
	defer r.mux.Unlock()

	k := fmt.Sprintf("ID-%d", d.ID())

	r.descriptors[k] = d
	r.descriptors[d.Subject()] = d
}
