package auth

import (
	"encoding/json"
	"os"
	"path"
	"testing"
	"time"
)

func Test_localCacheImpl_get_noCachedFile(t *testing.T) {
	fs := &mockFileSystem{e: false}
	currentCacheFileSystem = fs

	cache := newCache(false)
	_, ok := cache.get()
	if ok {
		t.Error("expected no cached file")
	}

	expectedFile := path.Join(os.TempDir(), tokenFileName)
	if fs.file != expectedFile {
		t.Errorf("unexpected exists file, got %s", fs.file)
	}
}

func Test_localCacheImpl_get_cachedFile(t *testing.T) {
	cachedResponse := cachedAuthenticationResponse{
		ValidUntil: time.Now().UTC().Add(time.Hour),
		Token:      "token",
	}
	fs := &mockFileSystem{e: true, cached: cachedResponse}
	currentCacheFileSystem = fs

	cache := newCache(false)
	response, ok := cache.get()
	if !ok {
		t.Error("expected ok")
	}

	if response.ClientToken() != cachedResponse.Token {
		t.Errorf("unexpected token, got %s", response.ClientToken())
	}

	expectedFile := path.Join(os.TempDir(), tokenFileName)
	if fs.file != expectedFile {
		t.Errorf("unexpected exists file, got %s", fs.file)
	}
}

func Test_localCacheImpl_get_staleCachedFile(t *testing.T) {
	cachedResponse := cachedAuthenticationResponse{
		ValidUntil: time.Now().UTC().Add(time.Hour * -1),
		Token:      "token",
	}
	fs := &mockFileSystem{e: true, cached: cachedResponse}
	currentCacheFileSystem = fs

	cache := newCache(false)
	_, ok := cache.get()
	if ok {
		t.Error("expected not ok")
	}
}

func Test_localCacheImpl_get_save(t *testing.T) {
	fs := &mockFileSystem{}
	currentCacheFileSystem = fs

	cache := newCache(false)

	response := cachedAuthenticationResponse{
		ValidUntil: time.Now().UTC().Add(time.Hour * -1),
		Token:      "token",
	}

	file, err := cache.save(response)
	if err != nil {
		t.Errorf("unexpected error, got %s", err)
	}

	expectedFile := path.Join(os.TempDir(), tokenFileName)
	if file != expectedFile {
		t.Errorf("unexpected exists file, got %s", fs.file)
	}

	if fs.stored.Token != response.Token {
		t.Errorf("unexpected token, got %s", fs.stored.Token)
	}
}

func Test_noopCache(t *testing.T) {
	cache := newCache(true)
	_, ok := cache.get()
	if ok {
		t.Error("expected no cached file")
	}

	fn, err := cache.save(cachedAuthenticationResponse{})
	if err != nil {
		t.Errorf("unexpected error, got %s", err)
	}
	if fn != "noop_no_file" {
		t.Errorf("unexpected file name, got %s", fn)
	}
}

type mockFileSystem struct {
	e              bool
	file           string
	cached, stored cachedAuthenticationResponse
}

func (m *mockFileSystem) exists(s string) bool {
	m.file = s
	return m.e
}

func (m *mockFileSystem) read(s string) ([]byte, error) {
	return json.Marshal(m.cached)
}

func (m *mockFileSystem) store(s string, bytes []byte) error {
	m.file = s
	if err := json.Unmarshal(bytes, &m.stored); err != nil {
		return err
	}
	return nil
}
