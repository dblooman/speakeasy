package openapi

import (
	"context"
	"net/http"
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
