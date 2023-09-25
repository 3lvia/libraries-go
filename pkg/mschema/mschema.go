package mschema

// New creates and returns a new schema registry instance.
func New(url string, opts ...Option) (Registry, error) {
	collector := &optionsCollector{}
	for _, opt := range opts {
		opt(collector)
	}

	r, err := newRegistry(url, collector)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func Work() {}
