package hashivault

import (
	"os"
	"testing"
)

func Test_optionsCollector_validate_nothing(t *testing.T) {
	clearEnvVars(t)
	c := &optionsCollector{}
	err := c.build()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "VAULT_ADDR not set" {
		t.Fatal("unexpected error message")
	}
}

func Test_optionsCollector_validate_nothingButVaultAddress(t *testing.T) {
	clearEnvVars(t)

	opt := WithVaultAddress("http://localhost:8200")
	c := &optionsCollector{}
	opt(c)

	err := c.build()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "GITHUB_TOKEN or MOUNT_PATH not set" {
		t.Fatalf("unexpected error message, got: %s", err.Error())
	}
}

func Test_optionsCollector_validate_options(t *testing.T) {
	clearEnvVars(t)

	c := &optionsCollector{}
	opt := WithVaultAddress("http://localhost:8200")
	opt(c)
	opt = WithGitHubToken("my-token")
	opt(c)

	if err := c.build(); err != nil {
		t.Fatal(err)
	}

	if c.vaultAddress != "http://localhost:8200" {
		t.Errorf("unexpected vault address, got: %s", c.vaultAddress)
	}
	if c.gitHubToken != "my-token" {
		t.Errorf("unexpected github token, got: %s", c.gitHubToken)
	}
}

func Test_optionsCollector_validate_overrideWithEnvVars(t *testing.T) {
	clearEnvVars(t)

	c := &optionsCollector{}
	opt := WithVaultAddress("http://localhost:8200")
	opt(c)
	opt = WithGitHubToken("my-token")
	opt(c)

	if err := os.Setenv("VAULT_ADDR", "http://localhost:8201"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("GITHUB_TOKEN", "my-other-token"); err != nil {
		t.Fatal(err)
	}

	if err := c.build(); err != nil {
		t.Fatal(err)
	}

	if c.vaultAddress != "http://localhost:8201" {
		t.Errorf("unexpected vault address, got: %s", c.vaultAddress)
	}
	if c.gitHubToken != "my-other-token" {
		t.Errorf("unexpected github token, got: %s", c.gitHubToken)
	}
}

func Test_optionsCollector_validate_FromEnvVars(t *testing.T) {
	clearEnvVars(t)
	if err := os.Setenv("VAULT_ADDR", "http://localhost:8200"); err != nil {
		t.Fatal(err)
	}
	if err := os.Setenv("GITHUB_TOKEN", "my-token"); err != nil {
		t.Fatal(err)
	}

	c := &optionsCollector{}
	//opt(c)

	if err := c.build(); err != nil {
		t.Fatal(err)
	}

	if c.vaultAddress != "http://localhost:8200" {
		t.Errorf("unexpected vault address, got: %s", c.vaultAddress)
	}
	if c.gitHubToken != "my-token" {
		t.Errorf("unexpected github token, got: %s", c.gitHubToken)
	}
}

func clearEnvVars(t *testing.T) {
	if err := os.Unsetenv("VAULT_ADDR"); err != nil {
		t.Fatal(err)
	}
	if err := os.Unsetenv("GITHUB_TOKEN"); err != nil {
		t.Fatal(err)
	}
	if err := os.Unsetenv("MOUNT_PATH"); err != nil {
		t.Fatal(err)
	}
	if err := os.Unsetenv("ROLE"); err != nil {
		t.Fatal(err)
	}
}
