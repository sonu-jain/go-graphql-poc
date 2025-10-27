package client

import (
	"fmt"
)

// CustomerType represents the customer type enum
type CustomerType string

const (
	CustomerTypeIndividual CustomerType = "INDIVIDUAL"
	CustomerTypeBusiness   CustomerType = "BUSINESS"
	CustomerTypePremium    CustomerType = "PREMIUM"
)

// CustomerStatus represents the customer status enum
type CustomerStatus string

const (
	CustomerStatusActive    CustomerStatus = "ACTIVE"
	CustomerStatusInactive  CustomerStatus = "INACTIVE"
	CustomerStatusSuspended CustomerStatus = "SUSPENDED"
	CustomerStatusPending   CustomerStatus = "PENDING"
)

// GetCustomers retrieves all customers with pagination
func (c *GraphQLClient) GetCustomers(page, offset int) ([]interface{}, error) {
	query := `
		query GetCustomers($page: Int, $offset: Int) {
			customers(page: $page, offset: $offset) {
				... on IndividualCustomer {
					id
					name
					email
					createdAt
					updatedAt
					personalInfo {
						phone
						address
						dateOfBirth
					}
				}
				... on BusinessCustomer {
					id
					name
					email
					createdAt
					updatedAt
					companyName
					businessInfo {
						taxId
						industry
						employeeCount
						website
					}
				}
				... on PremiumCustomer {
					id
					name
					email
					createdAt
					updatedAt
					premiumTier
					benefits
				}
			}
		}
	`

	variables := map[string]interface{}{
		"page":   page,
		"offset": offset,
	}

	var result struct {
		Customers []interface{} `json:"customers"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return nil, fmt.Errorf("failed to get customers: %w", err)
	}

	return result.Customers, nil
}

// GetCustomer retrieves a single customer by ID
func (c *GraphQLClient) GetCustomer(id string) (interface{}, error) {
	query := `
		query GetCustomer($id: ID!) {
			customer(id: $id) {
				... on IndividualCustomer {
					id
					name
					email
					createdAt
					updatedAt
					personalInfo {
						phone
						address
						dateOfBirth
					}
				}
				... on BusinessCustomer {
					id
					name
					email
					createdAt
					updatedAt
					companyName
					businessInfo {
						taxId
						industry
						employeeCount
						website
					}
				}
				... on PremiumCustomer {
					id
					name
					email
					createdAt
					updatedAt
					premiumTier
					benefits
				}
			}
		}
	`

	variables := map[string]interface{}{
		"id": id,
	}

	var result struct {
		Customer interface{} `json:"customer"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return nil, fmt.Errorf("failed to get customer: %w", err)
	}

	return result.Customer, nil
}

// GetCustomersByType retrieves customers filtered by type
func (c *GraphQLClient) GetCustomersByType(customerType CustomerType, page, offset int) ([]interface{}, error) {
	query := `
		query GetCustomersByType($type: CustomerType!, $page: Int, $offset: Int) {
			customersByType(type: $type, page: $page, offset: $offset) {
				... on IndividualCustomer {
					id
					name
					email
					createdAt
					updatedAt
					personalInfo {
						phone
						address
						dateOfBirth
					}
				}
				... on BusinessCustomer {
					id
					name
					email
					createdAt
					updatedAt
					companyName
					businessInfo {
						taxId
						industry
						employeeCount
						website
					}
				}
				... on PremiumCustomer {
					id
					name
					email
					createdAt
					updatedAt
					premiumTier
					benefits
				}
			}
		}
	`

	variables := map[string]interface{}{
		"type":   customerType,
		"page":   page,
		"offset": offset,
	}

	var result struct {
		CustomersByType []interface{} `json:"customersByType"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return nil, fmt.Errorf("failed to get customers by type: %w", err)
	}

	return result.CustomersByType, nil
}

// SearchCustomers searches customers by query string
func (c *GraphQLClient) SearchCustomers(query string) ([]interface{}, error) {
	searchQuery := `
		query SearchCustomers($query: String!) {
			searchCustomers(query: $query) {
				... on IndividualCustomer {
					id
					name
					email
					createdAt
					updatedAt
					personalInfo {
						phone
						address
						dateOfBirth
					}
				}
				... on BusinessCustomer {
					id
					name
					email
					createdAt
					updatedAt
					companyName
					businessInfo {
						taxId
						industry
						employeeCount
						website
					}
				}
				... on PremiumCustomer {
					id
					name
					email
					createdAt
					updatedAt
					premiumTier
					benefits
				}
			}
		}
	`

	variables := map[string]interface{}{
		"query": query,
	}

	var result struct {
		SearchCustomers []interface{} `json:"searchCustomers"`
	}

	if err := c.ExecuteWithResult(searchQuery, variables, &result); err != nil {
		return nil, fmt.Errorf("failed to search customers: %w", err)
	}

	return result.SearchCustomers, nil
}

// GetCustomerWithErrorHandling retrieves a customer with error handling
func (c *GraphQLClient) GetCustomerWithErrorHandling(id string) (interface{}, error) {
	query := `
		query GetCustomerWithErrorHandling($id: ID!) {
			getCustomerWithErrorHandling(id: $id) {
				... on IndividualCustomer {
					id
					name
					email
					createdAt
					updatedAt
					personalInfo {
						phone
						address
						dateOfBirth
					}
				}
				... on BusinessCustomer {
					id
					name
					email
					createdAt
					updatedAt
					companyName
					businessInfo {
						taxId
						industry
						employeeCount
						website
					}
				}
				... on PremiumCustomer {
					id
					name
					email
					createdAt
					updatedAt
					premiumTier
					benefits
				}
				... on OperationError {
					code
					message
					field
				}
			}
		}
	`

	variables := map[string]interface{}{
		"id": id,
	}

	var result struct {
		GetCustomerWithErrorHandling interface{} `json:"getCustomerWithErrorHandling"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return nil, fmt.Errorf("failed to get customer with error handling: %w", err)
	}

	return result.GetCustomerWithErrorHandling, nil
}

// GetCustomersByStatus retrieves customers filtered by status
func (c *GraphQLClient) GetCustomersByStatus(status CustomerStatus, page, offset int) ([]interface{}, error) {
	query := `
		query GetCustomersByStatus($status: CustomerStatus!, $page: Int, $offset: Int) {
			customersByStatus(status: $status, page: $page, offset: $offset) {
				... on IndividualCustomer {
					id
					name
					email
					createdAt
					updatedAt
					personalInfo {
						phone
						address
						dateOfBirth
					}
				}
				... on BusinessCustomer {
					id
					name
					email
					createdAt
					updatedAt
					companyName
					businessInfo {
						taxId
						industry
						employeeCount
						website
					}
				}
				... on PremiumCustomer {
					id
					name
					email
					createdAt
					updatedAt
					premiumTier
					benefits
				}
			}
		}
	`

	variables := map[string]interface{}{
		"status": status,
		"page":   page,
		"offset": offset,
	}

	var result struct {
		CustomersByStatus []interface{} `json:"customersByStatus"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return nil, fmt.Errorf("failed to get customers by status: %w", err)
	}

	return result.CustomersByStatus, nil
}

// GetPremiumCustomersByTier retrieves premium customers by tier
func (c *GraphQLClient) GetPremiumCustomersByTier(tier string, page, offset int) ([]PremiumCustomer, error) {
	query := `
		query GetPremiumCustomersByTier($tier: String!, $page: Int, $offset: Int) {
			premiumCustomersByTier(tier: $tier, page: $page, offset: $offset) {
				id
				name
				email
				createdAt
				updatedAt
				premiumTier
				benefits
			}
		}
	`

	variables := map[string]interface{}{
		"tier":   tier,
		"page":   page,
		"offset": offset,
	}

	var result struct {
		PremiumCustomersByTier []PremiumCustomer `json:"premiumCustomersByTier"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return nil, fmt.Errorf("failed to get premium customers by tier: %w", err)
	}

	return result.PremiumCustomersByTier, nil
}

// Print functions for easy testing

// GetCustomersAndPrint retrieves customers and prints the result
func (c *GraphQLClient) GetCustomersAndPrint(page, offset int) {
	fmt.Println("📋 Getting customers...")

	customers, err := c.GetCustomers(page, offset)
	if err != nil {
		fmt.Printf("❌ Failed to get customers: %v\n", err)
		return
	}

	fmt.Printf("✅ Found %d customers:\n", len(customers))
	for i, customer := range customers {
		fmt.Printf("  %d. %+v\n", i+1, customer)
	}
}

// GetCustomerAndPrint retrieves a single customer and prints the result
func (c *GraphQLClient) GetCustomerAndPrint(id string) {
	fmt.Printf("👤 Getting customer with ID: %s\n", id)

	customer, err := c.GetCustomer(id)
	if err != nil {
		fmt.Printf("❌ Failed to get customer: %v\n", err)
		return
	}

	fmt.Printf("✅ Customer found: %+v\n", customer)
}

// SearchCustomersAndPrint searches customers and prints the result
func (c *GraphQLClient) SearchCustomersAndPrint(query string) {
	fmt.Printf("🔍 Searching customers for: %s\n", query)

	customers, err := c.SearchCustomers(query)
	if err != nil {
		fmt.Printf("❌ Failed to search customers: %v\n", err)
		return
	}

	fmt.Printf("✅ Found %d customers matching '%s':\n", len(customers), query)
	for i, customer := range customers {
		fmt.Printf("  %d. %+v\n", i+1, customer)
	}
}
