vault login -method=oidc

go run main.go list --system=edna --vaultToken=TOKEN

# schema-cli
This is a command line interface (CLI) for interacting with the Confluence Schema API. The main use case is to generate
model classes that match the messages on a topic as specified by the corresponding schema.

## Usage
The CLI is invoked with the `schema-cli` command. It is designed to be self-documenting, so you can run `schema-cli help`
to see the available commands and options.

## Current features
Currently, only a simple list command and a command to generate model classes for schemas that are of type AVRO are
implemented. 

Future features include:
- Support for other schema types

## Authentication
The CLI uses the Hashicorp Vault to authenticate with the Confluence Schema API. You can authenticate with the CLI
either by letting it use the OICD flow to authenticate with Vault, or by providing a Vault token directly. Such a
token can be obtained by running `vault login -method=oidc` and copying the token from the output. This token is used 
with this cli either by passing it as the `--vaultToken` option, or by setting the `VAULT_TOKEN` environment variable.