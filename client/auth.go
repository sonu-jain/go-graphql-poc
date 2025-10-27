package client

import (
	"fmt"
)

// LoginInput represents the input for login
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents the response from login
type LoginResponse struct {
	Token    string      `json:"token"`
	Customer interface{} `json:"customer"`
}

// Login performs a login request and returns the token and customer info
func (c *GraphQLClient) Login(email, password string) (*LoginResponse, error) {
	query := `
		query Login($input: LoginInput!) {
			login(input: $input) {
				token
				customer {
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
		}
	`

	variables := map[string]interface{}{
		"input": LoginInput{
			Email:    email,
			Password: password,
		},
	}

	var result struct {
		Login LoginResponse `json:"login"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	// Set the token for future requests
	c.SetToken(result.Login.Token)

	// Save token for persistence across commands
	if err := SaveToken(result.Login.Token); err != nil {
		// Don't fail the login if we can't save the token
		fmt.Printf("Warning: Could not save token: %v\n", err)
	}

	return &result.Login, nil
}

// LoginAndPrint performs login and prints the result
func (c *GraphQLClient) LoginAndPrint(email, password string) error {
	fmt.Println("🔐 Attempting to login...")

	loginResp, err := c.Login(email, password)
	if err != nil {
		fmt.Printf("❌ Login failed: %v\n", err)
		return err
	}

	fmt.Printf("✅ Login successful!\n")
	fmt.Printf("🔑 Token: %s\n", loginResp.Token)
	fmt.Printf("👤 Customer: %+v\n", loginResp.Customer)

	return nil
}
