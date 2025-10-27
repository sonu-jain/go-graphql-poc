//go:build integration

package client

import (
	"testing"
	"time"
)

// IntegrationTest tests the client against a running server
// This test should be run with: go test -tags=integration

func TestIntegration(t *testing.T) {
	// Skip if not running integration tests
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := NewGraphQLClient("http://localhost:8080/query")

	// Test creating an individual customer
	t.Run("CreateIndividualCustomer", func(t *testing.T) {
		customer, err := client.CreateIndividualCustomer(CreateIndividualCustomerInput{
			Name:     "Integration Test User",
			Email:    "integration@test.com",
			Password: "password123",
			PersonalInfo: &PersonalInfoInput{
				Phone: stringPtr("+1-555-INTEGRATION"),
			},
		})

		if err != nil {
			t.Fatalf("Failed to create individual customer: %v", err)
		}

		if customer.Name != "Integration Test User" {
			t.Errorf("Expected name to be 'Integration Test User', got %s", customer.Name)
		}

		if customer.Email != "integration@test.com" {
			t.Errorf("Expected email to be 'integration@test.com', got %s", customer.Email)
		}
	})

	// Test login
	t.Run("Login", func(t *testing.T) {
		loginResp, err := client.Login("integration@test.com", "password123")
		if err != nil {
			t.Fatalf("Failed to login: %v", err)
		}

		if loginResp.Token == "" {
			t.Error("Expected token to be non-empty")
		}

		// Verify token is set on client
		if client.Token != loginResp.Token {
			t.Error("Expected client token to be set after login")
		}
	})

	// Test creating business customer
	t.Run("CreateBusinessCustomer", func(t *testing.T) {
		customer, err := client.CreateBusinessCustomer(CreateBusinessCustomerInput{
			Name:        "Integration Business User",
			Email:       "business@integration.com",
			Password:    "password123",
			CompanyName: "Integration Corp",
			BusinessInfo: &BusinessInfoInput{
				Industry:      stringPtr("Technology"),
				EmployeeCount: int32Ptr(10),
			},
		})

		if err != nil {
			t.Fatalf("Failed to create business customer: %v", err)
		}

		if customer.CompanyName != "Integration Corp" {
			t.Errorf("Expected company name to be 'Integration Corp', got %s", customer.CompanyName)
		}
	})

	// Test creating premium customer
	t.Run("CreatePremiumCustomer", func(t *testing.T) {
		customer, err := client.CreatePremiumCustomer(CreatePremiumCustomerInput{
			Name:        "Integration Premium User",
			Email:       "premium@integration.com",
			Password:    "password123",
			PremiumTier: "GOLD",
		})

		if err != nil {
			t.Fatalf("Failed to create premium customer: %v", err)
		}

		if customer.PremiumTier != "GOLD" {
			t.Errorf("Expected premium tier to be 'GOLD', got %s", customer.PremiumTier)
		}

		if len(customer.Benefits) == 0 {
			t.Error("Expected benefits to be populated")
		}
	})

	// Test error handling
	t.Run("ErrorHandling", func(t *testing.T) {
		// Try to create customer with invalid data
		_, err := client.CreateIndividualCustomer(CreateIndividualCustomerInput{
			Name:     "", // Empty name should cause error
			Email:    "invalid-email",
			Password: "123", // Short password
		})

		if err == nil {
			t.Error("Expected error when creating customer with invalid data")
		}
	})

	// Test custom GraphQL query
	t.Run("CustomQuery", func(t *testing.T) {
		query := `
			query {
				customers(page: 1, offset: 0) {
					... on IndividualCustomer {
						id
						name
						email
					}
				}
			}
		`

		resp, err := client.Execute(query, nil)
		if err != nil {
			t.Fatalf("Failed to execute custom query: %v", err)
		}

		if resp.Data == nil {
			t.Error("Expected response data to be non-nil")
		}
	})
}

// TestServerConnection tests if the server is reachable
func TestServerConnection(t *testing.T) {
	client := NewGraphQLClient("http://localhost:8080/query")

	// Simple introspection query to test connection
	query := `
		query {
			__schema {
				queryType {
					name
				}
			}
		}
	`

	resp, err := client.Execute(query, nil)
	if err != nil {
		t.Skipf("Server not available: %v", err)
	}

	if resp.Data == nil {
		t.Error("Expected response data to be non-nil")
	}
}

// BenchmarkCustomerCreation benchmarks customer creation
func BenchmarkCustomerCreation(b *testing.B) {
	client := NewGraphQLClient("http://localhost:8080/query")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.CreateIndividualCustomer(CreateIndividualCustomerInput{
			Name:     "Benchmark User",
			Email:    "benchmark@test.com",
			Password: "password123",
		})
		if err != nil {
			b.Fatalf("Failed to create customer: %v", err)
		}
		time.Sleep(10 * time.Millisecond) // Small delay to avoid overwhelming the server
	}
}
