package hydros

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHydrosClient(t *testing.T) {

	client, err := NewHydrosClient()
	assert.Nil(t, err, "Error should be nil.")
	assert.NotNil(t, client, "Client should not be nil")
	assert.Equal(t, DefaultURL, client.URL.String(), "Host URL should be default")
}

func TestSetHost(t *testing.T) {

	client, err := NewHydrosClient(SetHost("https://api.somewhere.com"))
	assert.Nil(t, err, "Error should be nil.")
	assert.NotNil(t, client, "Client should not be nil")
	assert.Equal(t, "https://api.somewhere.com", client.URL.String(), "Host URL should have been updated")

	client, err = NewHydrosClient(SetHost("https://api/somePath"))
	assert.Nil(t, err, "Error should be nil.")
	assert.NotNil(t, client, "Client should not be nil")
	assert.Equal(t, "https", client.URL.Scheme)
	assert.Equal(t, "api", client.URL.Host)
	assert.Equal(t, "/somePath", client.URL.Path)

	client, err = NewHydrosClient(SetHost("api"))
	assert.NotNil(t, err, "Error should not be nil.")
	assert.Nil(t, client, "Client should be nil")
	assert.Equal(t, "missing scheme from host parameter: please prefix with http:// or https://", err.Error())

	client, err = NewHydrosClient(SetHost(""))
	assert.NotNil(t, err, "Error should not be nil.")
	assert.Nil(t, client, "Client should be nil")
	assert.Equal(t, "missing scheme from host parameter: please prefix with http:// or https://", err.Error())
}

func TestSetAccessToken(t *testing.T) {
	client, err := NewHydrosClient(SetAccessToken("2718281828259"))
	assert.Nil(t, err, "Error should be nil.")
	assert.NotNil(t, client, "Client should not be nil")
	assert.Equal(t, AuthTypeOpenID, client.AuthType)
	assert.Equal(t, "2718281828259", client.AccessToken)
}
