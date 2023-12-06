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

# Testing

To run the tests, run:

```
go test ./...
```