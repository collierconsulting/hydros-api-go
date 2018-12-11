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

This library contains helper functions to assist in mocking of service methods for testing.  

#### Service Method Mocking

For example, you could mock out the driller service's `Get()` routine to return a driller with the same ID passed in:
```go
err = MockServiceMethod(
	client,
	"Driller.Get",
	func(ID uint) (*DrillerModel, error) {
		return &DrillerModel{DefaultModelBase: &DefaultModelBase{ID: ID}}, nil
	})
```

#### Model Service Mocking
To mock service methods on the payload model such as `Delete()` and `Save()`:
```go 
err = MockModelServiceMethod(
	client.Driller,
	"Save",
	func(model *DrillerModel) (*DrillerModel, error) {
		return model, nil
	})
```

**Note:** There is one exception to the above.  If you have mocked a service method, the model returned by that service method 
will not contain service method implementations or mocks.  If you need to mock a service method that returns a model with 
its own mocked service methods you can define them both at the same time by mocking defining a `ServiceSpec` 
with `ModelServiceCallMocks` and initializing it on the returned model.

The following mocks the `Well.Search` method that returned a Well model with a mocked `Delete` method.

```go
err := hydros.MockServiceMethod(
	app.APIClient,
	"Well.Search",
	func(query string, filters []string, from int, size int, sort []hydros.Sort) (*hydros.WellSearchResults, error) {
		wells := make([]*hydros.WellModel, 1)
		wells[0] = (&hydros.WellModel{
			DefaultModelBase: &hydros.DefaultModelBase{ID: 1}}).Init(
			&hydros.ServiceSpec{
				ServiceName:      "wells",
				Client:           app.APIClient.Client,
				PayloadModelType: reflect.TypeOf(hydros.WellModel{}),
				ModelServiceCallMocks: map[string]*hydros.ModelServiceCallMock{
					"Delete": {
						MockFunc: func(model *hydros.WellModel) error {
							return nil
						}}},
			})
		return &hydros.WellSearchResults{Total: 1, Results: wells}, nil
	})
```