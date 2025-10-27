package client

import (
	"fmt"
)

// UpdateCustomerInput represents the input for updating a customer
type UpdateCustomerInput struct {
	Name         *string            `json:"name,omitempty"`
	Email        *string            `json:"email,omitempty"`
	CompanyName  *string            `json:"companyName,omitempty"`
	PremiumTier  *string            `json:"premiumTier,omitempty"`
	PersonalInfo *PersonalInfoInput `json:"personalInfo,omitempty"`
	BusinessInfo *BusinessInfoInput `json:"businessInfo,omitempty"`
}

// UpdateCustomer updates an existing customer
func (c *GraphQLClient) UpdateCustomer(id string, input UpdateCustomerInput) (interface{}, error) {
	query := `
		mutation UpdateCustomer($id: ID!, $input: UpdateCustomerInput!) {
			updateCustomer(id: $id, input: $input) {
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
		"id":    id,
		"input": input,
	}

	var result struct {
		UpdateCustomer interface{} `json:"updateCustomer"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return nil, fmt.Errorf("failed to update customer: %w", err)
	}

	return result.UpdateCustomer, nil
}

// DeleteCustomer deletes a customer by ID
func (c *GraphQLClient) DeleteCustomer(id string) (bool, error) {
	query := `
		mutation DeleteCustomer($id: ID!) {
			deleteCustomer(id: $id)
		}
	`

	variables := map[string]interface{}{
		"id": id,
	}

	var result struct {
		DeleteCustomer bool `json:"deleteCustomer"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return false, fmt.Errorf("failed to delete customer: %w", err)
	}

	return result.DeleteCustomer, nil
}

// CreateCustomerWithErrorHandling creates a customer with error handling
func (c *GraphQLClient) CreateCustomerWithErrorHandling(input CreateIndividualCustomerInput) (interface{}, error) {
	query := `
		mutation CreateCustomerWithErrorHandling($input: CreateIndividualCustomerInput!) {
			createCustomerWithErrorHandling(input: $input) {
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
		"input": input,
	}

	var result struct {
		CreateCustomerWithErrorHandling interface{} `json:"createCustomerWithErrorHandling"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return nil, fmt.Errorf("failed to create customer with error handling: %w", err)
	}

	return result.CreateCustomerWithErrorHandling, nil
}

// Print functions for easy testing

// UpdateCustomerAndPrint updates a customer and prints the result
func (c *GraphQLClient) UpdateCustomerAndPrint(id string, input UpdateCustomerInput) {
	fmt.Printf("✏️ Updating customer with ID: %s\n", id)

	customer, err := c.UpdateCustomer(id, input)
	if err != nil {
		fmt.Printf("❌ Failed to update customer: %v\n", err)
		return
	}

	fmt.Printf("✅ Customer updated successfully: %+v\n", customer)
}

// DeleteCustomerAndPrint deletes a customer and prints the result
func (c *GraphQLClient) DeleteCustomerAndPrint(id string) {
	fmt.Printf("🗑️ Deleting customer with ID: %s\n", id)

	success, err := c.DeleteCustomer(id)
	if err != nil {
		fmt.Printf("❌ Failed to delete customer: %v\n", err)
		return
	}

	if success {
		fmt.Printf("✅ Customer deleted successfully\n")
	} else {
		fmt.Printf("❌ Customer deletion failed\n")
	}
}

// CreateCustomerWithErrorHandlingAndPrint creates a customer with error handling and prints the result
func (c *GraphQLClient) CreateCustomerWithErrorHandlingAndPrint(input CreateIndividualCustomerInput) {
	fmt.Println("⚠️ Creating customer with error handling...")

	result, err := c.CreateCustomerWithErrorHandling(input)
	if err != nil {
		fmt.Printf("❌ Failed to create customer: %v\n", err)
		return
	}

	fmt.Printf("✅ Customer creation result: %+v\n", result)
}

// Convenience functions for common update operations

// UpdateCustomerName updates only the customer's name
func (c *GraphQLClient) UpdateCustomerName(id, name string) (interface{}, error) {
	input := UpdateCustomerInput{
		Name: stringPtr(name),
	}
	return c.UpdateCustomer(id, input)
}

// UpdateCustomerEmail updates only the customer's email
func (c *GraphQLClient) UpdateCustomerEmail(id, email string) (interface{}, error) {
	input := UpdateCustomerInput{
		Email: stringPtr(email),
	}
	return c.UpdateCustomer(id, input)
}

// UpdateCustomerPersonalInfo updates only the customer's personal info
func (c *GraphQLClient) UpdateCustomerPersonalInfo(id string, personalInfo PersonalInfoInput) (interface{}, error) {
	input := UpdateCustomerInput{
		PersonalInfo: &personalInfo,
	}
	return c.UpdateCustomer(id, input)
}

// UpdateCustomerBusinessInfo updates only the customer's business info
func (c *GraphQLClient) UpdateCustomerBusinessInfo(id string, businessInfo BusinessInfoInput) (interface{}, error) {
	input := UpdateCustomerInput{
		BusinessInfo: &businessInfo,
	}
	return c.UpdateCustomer(id, input)
}

// UpdateCustomerPremiumTier updates only the customer's premium tier
func (c *GraphQLClient) UpdateCustomerPremiumTier(id, premiumTier string) (interface{}, error) {
	input := UpdateCustomerInput{
		PremiumTier: stringPtr(premiumTier),
	}
	return c.UpdateCustomer(id, input)
}
