package mschema

// New creates and returns a new schema registry instance.
func New(opts ...Option) (Registry, error) {
	collector := &optionsCollector{}
	for _, opt := range opts {
		opt(collector)
	}

	r := newRegistry(collector)
	return r, nil
}
