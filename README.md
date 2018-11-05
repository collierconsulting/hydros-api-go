# HYDROS API Go Client

While older versions of go should work, we only support builds using go versions >= 11.  

[![Build Status](https://travis-ci.com/collierconsulting/hydros-api-go.svg?token=HfjrsxGu5QnCecfDNiK9&branch=master)](https://travis-ci.com/collierconsulting/hydros-api-go)

## Example
To use client, import:

```go
import "github.com/collierconsulting/hydros-api-go"
```

Intializing the client
```go
client, err := hydros.NewHydrosClient(hydros.SetHost("https://the.apihost.com"), hydros.SetAccessToken("[your access token]"))
```