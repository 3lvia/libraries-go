package auth

import (
	"encoding/json"
	"os"
	"path"
	"time"
)

const (
	tokenFileName = "token-311bbd6d-9ed6-40e7-9888-c49e51b9d12f.json"
)

var currentCacheFileSystem cacheFileSystem

func newCache() localCache {
	return &localCacheImpl{}
}

type localCache interface {
	get() (AuthenticationResponse, bool)
	save(AuthenticationResponse) (string, error)
}

type localCacheImpl struct {
}

func (l localCacheImpl) get() (AuthenticationResponse, bool) {
	fs := currentCacheFileSystem
	if fs == nil {
		fs = cacheFileSystemImpl{}
	}
	return l.getCore(fs)
}

func (l localCacheImpl) getCore(fs cacheFileSystem) (AuthenticationResponse, bool) {
	f := path.Join(os.TempDir(), tokenFileName)
	if !fs.exists(f) {
		return nil, false
	}

	b, err := fs.read(f)
	if err != nil {
		return nil, false
	}

	var cr cachedAuthenticationResponse
	if err := json.Unmarshal(b, &cr); err != nil {
		return nil, false
	}

	if time.Now().UTC().After(cr.ValidUntil) {
		return nil, false
	}

	return cr, true
}

func (l localCacheImpl) save(response AuthenticationResponse) (string, error) {
	fs := currentCacheFileSystem
	if fs == nil {
		fs = cacheFileSystemImpl{}
	}
	return l.saveCore(response, fs)
}

func (l localCacheImpl) saveCore(response AuthenticationResponse, fs cacheFileSystem) (string, error) {
	validUntil := time.Now().UTC().Add(time.Duration(response.LeaseDurationSeconds()) * time.Second)
	cr := cachedAuthenticationResponse{
		ValidUntil:    validUntil,
		Token:         response.ClientToken(),
		LeaseDuration: response.LeaseDurationSeconds(),
		IsRenewable:   response.Renewable(),
	}

	b, err := json.Marshal(cr)
	if err != nil {
		return "", err
	}

	p := path.Join(os.TempDir(), tokenFileName)

	if err = fs.store(p, b); err != nil {
		return "", err
	}

	return p, nil
}

type cachedAuthenticationResponse struct {
	ValidUntil    time.Time `json:"validUntil"`
	Token         string    `json:"token"`
	LeaseDuration int       `json:"leaseDuration"`
	IsRenewable   bool      `json:"renewable"`
}

func (c cachedAuthenticationResponse) ClientToken() string {
	return c.Token
}

func (c cachedAuthenticationResponse) LeaseDurationSeconds() int {
	return c.LeaseDuration
}

func (c cachedAuthenticationResponse) Renewable() bool {
	return c.IsRenewable
}

func (c cachedAuthenticationResponse) After() <-chan time.Time {
	secs := c.LeaseDuration
	if secs == 0 {
		secs = 3600 * 24 * 365 // 1 year
	}
	if secs >= 60 {
		secs -= 30 // 30 seconds leeway
	} else {
		secs = 1 // 1 second leeway
	}
	return time.After(time.Duration(secs) * time.Second)
}

type cacheFileSystem interface {
	exists(string) bool
	read(string) ([]byte, error)
	store(string, []byte) error
}

type cacheFileSystemImpl struct{}

func (c cacheFileSystemImpl) exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (c cacheFileSystemImpl) read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (c cacheFileSystemImpl) store(path string, b []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(b)
	if err != nil {
		return err
	}
	return nil
}
