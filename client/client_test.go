package client

import (
	"testing"
)

func TestNewGraphQLClient(t *testing.T) {
	// Test with default URL
	client := NewGraphQLClient("")
	if client.client == nil {
		t.Error("Expected client to be initialized")
	}

	// Test with custom URL
	customURL := "http://localhost:3000/query"
	client = NewGraphQLClient(customURL)
	if client.client == nil {
		t.Error("Expected client to be initialized")
	}
}

func TestSetToken(t *testing.T) {
	client := NewGraphQLClient("")
	token := "test-token-123"

	client.SetToken(token)
	if client.Token != token {
		t.Errorf("Expected token to be %s, got %s", token, client.Token)
	}
}

func TestHelperFunctions(t *testing.T) {
	// Test stringPtr
	str := "test"
	ptr := stringPtr(str)
	if *ptr != str {
		t.Errorf("Expected stringPtr to return pointer to %s, got %s", str, *ptr)
	}

	// Test int32Ptr
	num := int32(42)
	ptr32 := int32Ptr(num)
	if *ptr32 != num {
		t.Errorf("Expected int32Ptr to return pointer to %d, got %d", num, *ptr32)
	}
}
