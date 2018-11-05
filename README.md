# HYDROS API Go Client

While older versions of go should work, we only support builds using go versions >= 11.  

[![Build Status](https://travis-ci.com/collierconsulting/hydros-api-go.svg?token=HfjrsxGu5QnCecfDNiK9&branch=master)](https://travis-ci.com/collierconsulting/hydros-api-go)

## Example
To use client, import:

```go
import "github.com/collierconsulting/hydros-api-go"
```

Initializing the client
```go
client, err := hydros.NewClient(
	hydros.SetHost("https://the.apihost.com"), 
	hydros.SetAccessToken("[your access token]"))
```

## Test Mocking

This library contains helper functions to assist in mocking of service methods for testing.  Currently you can only mock service 
methods.  I plan on additing the ability to mock model methods such as `Update()`, `Delete()` in the near future.  

```go
err = MockServiceMethod(
	client,
	"Driller.Get",
	func(ID uint) (*DrillerModel, error) {
		return &DrillerModel{DefaultModelBase: &DefaultModelBase{ID: ID}}, nil
	})
```
