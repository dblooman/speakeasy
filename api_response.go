package openapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type APIError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

// GetNature returns a NatureModel for the given id
func (c *Client) GetNature(id int32) (*NatureModel, error) {
	r, err := c.apiClient.DefaultAPI.NatureRead(c.Ctx, id).Execute()
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

	var natureModel NatureModel
	err = json.Unmarshal(body, &natureModel)
	if err != nil {
		return nil, err
	}

	return &natureModel, nil
}

// GetPokemon returns a PokemonModel for the given id
func (c *Client) GetPokemon(id int32) (*PokemonModel, error) {
	r, err := c.apiClient.DefaultAPI.PokemonRead(c.Ctx, id).Execute()
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

	var pokemonModel PokemonModel
	err = json.Unmarshal(body, &pokemonModel)
	if err != nil {
		return nil, err
	}

	return &pokemonModel, nil
}

// GetStat returns a StatModel for the given id
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
