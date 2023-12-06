package openapi

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	openapiclient "github.com/dblooman/speakeasy"
	"github.com/stretchr/testify/assert"
)

func TestPokemonAPI(t *testing.T) {
	apiClient := openapiclient.New(nil)

	response, err := apiClient.GetPokemon(1)
	assert.NoError(t, err)

	assert.Equal(t, "bulbasaur", response.Species.Name)
}

func TestNatureAPI(t *testing.T) {
	apiClient := openapiclient.New(nil)

	response, err := apiClient.GetNature(1)
	assert.NoError(t, err)

	assert.Equal(t, "hardy", response.Name)
}

func TestPokemonAPIError(t *testing.T) {
	apiClient := openapiclient.New(nil)

	_, err := apiClient.GetPokemon(0)
	assert.Error(t, err)
	assert.Equal(t, "404 Not Found", err.Error())
}

func TestClientConfiguration(t *testing.T) {
	ctx := context.Background()

	configuration := openapiclient.NewConfiguration()
	configuration.HTTPClient = &http.Client{
		Timeout: 5 * time.Second,
	}
	apiClient := openapiclient.NewWithContext(ctx, configuration)

	response, err := apiClient.GetPokemon(1)
	assert.NoError(t, err)

	assert.Equal(t, "bulbasaur", response.Species.Name)
}

func TestGetPokemonWithMockServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"id":1,"name":"bulbasaur","base_experience":64,"height":7,"is_default":true,"order":1,"weight":69}`)
	}))
	defer ts.Close()

	configuration := openapiclient.NewConfiguration()
	configuration.Servers = openapiclient.ServerConfigurations{
		{
			URL: ts.URL,
		},
	}

	apiClient := openapiclient.New(configuration)
	response, err := apiClient.GetPokemon(1)
	if err != nil {
		t.Errorf("Error when calling `read pokemon``: %v\n", err)
		return
	}

	if response.Name != "bulbasaur" {
		t.Errorf("Expected response.Name to be bulbasaur, got %s", response.Name)
	}
}
