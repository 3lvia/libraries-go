package cert

import (
	"errors"
	"os"
)

var (
	ErrFailedParsePEM = errors.New("failed to parse PEM")
	ErrLoadCAFromDisk = errors.New("failed to read CA file disk")
	ErrLoadCA         = errors.New("failed to load CA file")
)

// AppendFromFiles loads the certs from a slice of filenames to the cert pool
func (pool Pool) AppendFromFiles(certFiles []string) error {
	for _, cert := range certFiles {
		if cert != "" {
			if err := pool.loadFromFile(cert); err != nil {
				return err
			}
		}
	}

	return nil
}

// loadFromFile loads the certs from a specified file to the provided pool
func (pool Pool) loadFromFile(cert string) error {
	pem, err := os.ReadFile(cert)
	if err != nil {
		return errors.Join(ErrLoadCAFromDisk, err)
	}

	if err := pool.loadFromPEM(pem); err != nil {
		return errors.Join(ErrLoadCA, err)
	}

	return nil
}

// loadFromPEM appends the PEM-formatted certificate to the provided pool
func (pool Pool) loadFromPEM(pem []byte) error {
	if ok := pool.Certs.AppendCertsFromPEM(pem); !ok {
		return ErrFailedParsePEM
	}

	return nil
}