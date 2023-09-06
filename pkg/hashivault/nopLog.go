package hashivault

// nullWriter implements the io.Write interface but doesn't do anything.
type nullWriter int

// Write implements the io.Write interface but is a noop.
func (nullWriter) Write([]byte) (int, error) { return 0, nil }
