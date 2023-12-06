package openapi

import "context"

type Client struct {
	apiClient *APIClient
	Ctx       context.Context
}

func NewWithContext(ctx context.Context, configuration *Configuration) *Client {
	if configuration == nil {
		configuration = NewConfiguration()
	}
	apiClient := NewAPIClient(configuration)
	return &Client{apiClient: apiClient, Ctx: ctx}
}

func New(configuration *Configuration) *Client {
	if configuration == nil {
		configuration = NewConfiguration()
	}
	apiClient := NewAPIClient(configuration)
	return &Client{apiClient: apiClient, Ctx: context.Background()}
}
