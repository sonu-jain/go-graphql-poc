package client

import (
	"context"
	"fmt"
	"time"

	"github.com/machinebox/graphql"
)

// GraphQLClient represents a client for making GraphQL requests
type GraphQLClient struct {
	client *graphql.Client
	Token  string
}

// NewGraphQLClient creates a new GraphQL client
func NewGraphQLClient(url string) *GraphQLClient {
	if url == "" {
		url = "http://localhost:8080/query"
	}

	client := graphql.NewClient(url)

	return &GraphQLClient{
		client: client,
	}
}

// SetToken sets the authentication token for the client
func (c *GraphQLClient) SetToken(token string) {
	c.Token = token
}

// Execute executes a GraphQL request
func (c *GraphQLClient) Execute(query string, variables map[string]interface{}) (interface{}, error) {
	req := graphql.NewRequest(query)

	// Set variables
	for key, value := range variables {
		req.Var(key, value)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result interface{}

	err := c.client.Run(ctx, req, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GraphQL request: %w", err)
	}

	return result, nil
}

// ExecuteWithResult executes a GraphQL request and unmarshals the result into the provided interface
func (c *GraphQLClient) ExecuteWithResult(query string, variables map[string]interface{}, result interface{}) error {
	req := graphql.NewRequest(query)

	// Set variables
	for key, value := range variables {
		req.Var(key, value)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := c.client.Run(ctx, req, result)
	if err != nil {
		return fmt.Errorf("failed to execute GraphQL request: %w", err)
	}

	return nil
}
