# Pokemon SDK

## Overview

This is a Go SDK for the [PokeAPI](https://pokeapi.co/). It is generated using openapi Generator. To add additional functionality, please update the openapi spec, and regenerate the SDK using the Make task

# SDK Installation

```sh
go get github.com/dblooman/speakeasy
```

# SDK Usage

```go
package main

import (
	"fmt"
	"os"

	client "github.com/dblooman/speakeasy"
)

func main() {
	apiClient := client.New(nil)
	response, err := apiClient.GetStat(1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `read stat``: %v\n", err)
		return
	}

	// Handle response data
    //....
}
```

# View Stats

Stats determine certain aspects of battles. Each Pokémon has a value for each stat which grows as they gain levels and can be altered momentarily by effects in battles

```go
package main

import (
	"fmt"
	"os"

	client "github.com/dblooman/speakeasy"
)

func main() {
	apiClient := client.New(nil)
	response, err := apiClient.GetStat(1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `read stat``: %v\n", err)
		return
	}

	fmt.Println(response.Name)
}
```

# View Natures

Natures influence how a Pokémon's stats grow. See [Bulbapedia](http://bulbapedia.bulbagarden.net/wiki/Nature) for greater detail.

```go
package main

import (
	"fmt"
	"os"

	client "github.com/dblooman/speakeasy"
)

func main() {
	apiClient := client.New(nil)
	response, err := apiClient.GetNature(1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `read nature``: %v\n", err)
		return
	}

	fmt.Println(response.HatesFlavor)
}
```

# View Pokemon

Pokémon are the creatures that inhabit the world of the Pokémon games. They can be caught using Pokéballs and trained by battling with other Pokémon. Each Pokémon belongs to a specific species but may take on a variant which makes it differ from other Pokémon of the same species, such as base stats, available abilities and typings.

```go
import (
	"fmt"
	"os"

	openapiclient "github.com/dblooman/speakeasy/openapi"
)

func main() {
	apiClient := openapiclient.New(nil)
	response, err := apiClient.GetPokemon(1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when getting pokemon``: %v\n", err)
		return
	}

	fmt.Println(response.Species.Name)
}
```

# Client configuration

The client can be configured with authentication and connection details when created. See the documentation for more information.

```go
import (
    "fmt"
    "os"

    openapiclient "github.com/dblooman/speakeasy/openapi"
)

func main() {
    configuration := openapiclient.NewConfiguration()
	configuration.DefaultHeader = map[string]string{"User-Agent": "Speakeasy"}

    apiClient := openapiclient.New(configuration)
    response, err := apiClient.GetPokemon(1)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when getting pokemon``: %v\n", err)
		return
    }

    fmt.Println(response.Species.Name)
}
```

## With context

If using the client with a context, you can supply it to the NewWithContext function:

```go
import (
    "fmt"
    "os"
    "context"

    openapiclient "github.com/dblooman/speakeasy/openapi"
)

func main() {
    ctx := context.Background()
    apiClient := openapiclient.NewWithContext(ctx,configuration)
    response, err := apiClient.GetPokemon(1)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when getting pokemon``: %v\n", err)
		return
    }

    fmt.Println(response.Species.Name)
}
```

## Error handling

You can extract a custom error when the response from the API is unexepected:

```go
package main

import (
	"fmt"
	"os"

	openapiclient "github.com/dblooman/speakeasy/openapi"
)

func main() {
	apiClient := openapiclient.New(nil)
	response, err := apiClient.GetPokemon(1)
	if err != nil {
		if err == err.(*openapiclient.APIError) {
			fmt.Fprintf(os.Stderr, "Error when calling API, status code: %d, message: %s\n", err.(*openapiclient.APIError).Code, err.(*openapiclient.APIError).Message)
			return
		}
	}

	fmt.Println(response.Species.Name)
}
```

# Server Selection

You can set the server endpoing, localhost if you are running the API locally with the configuration

```go
package main

import (
	"fmt"

	openapiclient "github.com/dblooman/speakeasy/openapi"
)

func main() {
	configuration := openapiclient.NewConfiguration()
	configuration.Servers = openapiclient.ServerConfigurations{
		{
			URL: "https://pokeapi.com",
		},
	}

	apiClient := openapiclient.New(configuration)
	response, err := apiClient.GetPokemon(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Species.Name)
}
```

# Custom HTTP Client

You can set a custom HTTP client with the configuration

```go
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
```

# Testing using the SDK

We can use the httptest package to mock a server and test the SDK. The configuration can be set to use the mock server and the response can be set to return a specific response.

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	openapiclient "github.com/dblooman/speakeasy/openapi"
)

func TestGetPokemon(t *testing.T) {
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
```

# Adding a new endpoint to the SDK

To add a new endpoint to the SDK, update the openapi spec in api/openapi.yaml and run the Make task

A new route can be added to the openapi spec as follows:

```yaml
/api/v2/stat/{id}/:
    get:
      operationId: stat_read
      parameters:
	  ....
```

Once the spec has been updated, run the Make task to generate the SDK

```sh
make generate
```

Once generated, add a new function to the client to handle the new endpoint

```go
func (c *Client) GetStat(id int32) (*StatModel, error) {
	r, err := c.apiClient.DefaultAPI.StatRead(c.Ctx, id).Execute()
	if err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, APIError{Code: int32(r.StatusCode), Message: "status code was not 200"}
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var statModel StatModel
	err = json.Unmarshal(body, &statModel)
	if err != nil {
		return nil, err
	}

	return &statModel, nil
}
```

# Testing this SDK

To run the tests, run:

```
go test ./...
```
