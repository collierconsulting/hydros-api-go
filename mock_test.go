package hydros

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockServiceMethod(t *testing.T) {

	client, err := NewClient(SetHost("https://api.somewhere.com"))
	assert.Nil(t, err, "Error should be nil.")
	assert.NotNil(t, client, "Client should not be nil")

	err = MockServiceMethod(
		client,
		"Driller.Get",
		func(ID uint) (*DrillerModel, error) {
			return &DrillerModel{DefaultModelBase: &DefaultModelBase{ID: ID}}, nil
		})
	assert.Nil(t, err, "Error should be nil.")

	driller, err := client.Driller.Get(100)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(100), driller.ID, "Should have returned driller with same ID")

}

func TestMockServiceMethod_WithErrors(t *testing.T) {

	client, err := NewClient(SetHost("https://api.somewhere.com"))
	assert.Nil(t, err, "Error should be nil.")
	assert.NotNil(t, client, "Client should not be nil")

	err = MockServiceMethod(
		client,
		"SomeNonexistantService.Get",
		func(ID uint) (*DrillerModel, error) {
			return nil, nil
		})
	assert.NotNil(t, err, "Error should not be nil.")
	assert.Equal(t, err.Error(), "could not find service 'SomeNonexistantService'")

	err = MockServiceMethod(
		client,
		"Driller.MissingMethod",
		func(ID uint) (*DrillerModel, error) {
			return nil, nil
		})
	assert.NotNil(t, err, "Error should not be nil.")
	assert.Equal(t, err.Error(), "could not find method 'MissingMethod' in service 'Driller'")

	err = MockServiceMethod(
		client,
		"Driller.Get",
		func(ID int) (*DrillerModel, error) {
			return nil, nil
		})
	assert.NotNil(t, err, "Error should not be nil.")
	assert.Equal(t, err.Error(),
		"mock function is wrong type: expected: func(uint) (*hydros.DrillerModel, error) "+
			"but got: func(int) (*hydros.DrillerModel, error)")

}

func TestMockServiceMethod_WithClientAsAnonymousField(t *testing.T) {

	type ComposedClient struct {
		*Client
	}

	client, err := NewClient(SetHost("https://api.somewhere.com"))
	assert.Nil(t, err, "Error should be nil.")
	assert.NotNil(t, client, "Client should not be nil")

	composedClient := &ComposedClient{client}

	err = MockServiceMethod(
		composedClient,
		"Driller.Get",
		func(ID uint) (*DrillerModel, error) {
			return &DrillerModel{DefaultModelBase: &DefaultModelBase{ID: ID}}, nil
		})
	assert.Nil(t, err, "Error should be nil.")

	driller, err := client.Driller.Get(100)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(100), driller.ID, "Should have returned driller with same ID")
}
