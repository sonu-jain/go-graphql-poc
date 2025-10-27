package client

import (
	"os"
	"path/filepath"
)

var tokenFile = "graphql_token.txt"

// SaveToken saves the authentication token to a file
func SaveToken(token string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	tokenPath := filepath.Join(homeDir, tokenFile)
	return os.WriteFile(tokenPath, []byte(token), 0600)
}

// LoadToken loads the authentication token from a file
func LoadToken() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	tokenPath := filepath.Join(homeDir, tokenFile)
	tokenBytes, err := os.ReadFile(tokenPath)
	if err != nil {
		return "", err
	}

	return string(tokenBytes), nil
}

// ClearToken removes the saved token
func ClearToken() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	tokenPath := filepath.Join(homeDir, tokenFile)
	return os.Remove(tokenPath)
}

// NewGraphQLClientWithToken creates a new GraphQL client and loads any saved token
func NewGraphQLClientWithToken(url string) *GraphQLClient {
	client := NewGraphQLClient(url)

	// Try to load saved token
	if token, err := LoadToken(); err == nil && token != "" {
		client.SetToken(token)
	}

	return client
}
