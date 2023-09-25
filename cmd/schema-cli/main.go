/*
schema-cli is a command line tool for generating Go code from the confluent schema registry. It is a wrapper around the
libraries-go/schema package.

Currently, the only functioning command is List:

```
list --system=edna --storageFolder=../schemas
```
*/
package main

import "github.com/3lvia/libraries-go/cmd/schema-cli/cmd"

func main() {
	cmd.Execute()
}
