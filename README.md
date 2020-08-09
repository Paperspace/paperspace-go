# go-paperspace

## Usage
```go
package main

import (
    paperspace "github.com/Paperspace/paperspace-go"
)

func getClient() *paperspace.Client {
    client := paperspace.NewClient()
    client.APIKey = p.APIKey
    return client
}
```