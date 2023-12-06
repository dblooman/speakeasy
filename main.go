package main

import (
	"fmt"
	"net/http"
	"time"

	openapiclient "github.com/dblooman/speakeasy/openapi"
)

func main() {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	configuration := openapiclient.NewConfiguration()
	configuration.HTTPClient = &httpClient

	apiClient := openapiclient.New(configuration)
	response, err := apiClient.GetPokemon(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Species.Name)
}
