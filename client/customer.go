package client

import (
	"fmt"
)

// PersonalInfoInput represents personal information input
type PersonalInfoInput struct {
	Phone       *string `json:"phone,omitempty"`
	Address     *string `json:"address,omitempty"`
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
}

// BusinessInfoInput represents business information input
type BusinessInfoInput struct {
	TaxID         *string `json:"taxId,omitempty"`
	Industry      *string `json:"industry,omitempty"`
	EmployeeCount *int32  `json:"employeeCount,omitempty"`
	Website       *string `json:"website,omitempty"`
}

// CreateIndividualCustomerInput represents input for creating individual customer
type CreateIndividualCustomerInput struct {
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	Password     string             `json:"password"`
	PersonalInfo *PersonalInfoInput `json:"personalInfo,omitempty"`
}

// CreateBusinessCustomerInput represents input for creating business customer
type CreateBusinessCustomerInput struct {
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	Password     string             `json:"password"`
	CompanyName  string             `json:"companyName"`
	BusinessInfo *BusinessInfoInput `json:"businessInfo,omitempty"`
}

// CreatePremiumCustomerInput represents input for creating premium customer
type CreatePremiumCustomerInput struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PremiumTier string `json:"premiumTier"`
}

// IndividualCustomer represents an individual customer
type IndividualCustomer struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	CreatedAt    string        `json:"createdAt"`
	UpdatedAt    string        `json:"updatedAt"`
	PersonalInfo *PersonalInfo `json:"personalInfo,omitempty"`
}

// BusinessCustomer represents a business customer
type BusinessCustomer struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	CreatedAt    string        `json:"createdAt"`
	UpdatedAt    string        `json:"updatedAt"`
	CompanyName  string        `json:"companyName"`
	BusinessInfo *BusinessInfo `json:"businessInfo,omitempty"`
}

// PremiumCustomer represents a premium customer
type PremiumCustomer struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
	PremiumTier string   `json:"premiumTier"`
	Benefits    []string `json:"benefits"`
}

// PersonalInfo represents personal information
type PersonalInfo struct {
	Phone       *string `json:"phone,omitempty"`
	Address     *string `json:"address,omitempty"`
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
}

// BusinessInfo represents business information
type BusinessInfo struct {
	TaxID         *string `json:"taxId,omitempty"`
	Industry      *string `json:"industry,omitempty"`
	EmployeeCount *int32  `json:"employeeCount,omitempty"`
	Website       *string `json:"website,omitempty"`
}

// CreateIndividualCustomer creates a new individual customer
func (c *GraphQLClient) CreateIndividualCustomer(input CreateIndividualCustomerInput) (*IndividualCustomer, error) {
	query := `
		mutation CreateIndividualCustomer($input: CreateIndividualCustomerInput!) {
			createIndividualCustomer(input: $input) {
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
		}
	`

	variables := map[string]interface{}{
		"input": input,
	}

	var result struct {
		CreateIndividualCustomer IndividualCustomer `json:"createIndividualCustomer"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return nil, fmt.Errorf("failed to create individual customer: %w", err)
	}

	return &result.CreateIndividualCustomer, nil
}

// CreateBusinessCustomer creates a new business customer
func (c *GraphQLClient) CreateBusinessCustomer(input CreateBusinessCustomerInput) (*BusinessCustomer, error) {
	query := `
		mutation CreateBusinessCustomer($input: CreateBusinessCustomerInput!) {
			createBusinessCustomer(input: $input) {
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
		}
	`

	variables := map[string]interface{}{
		"input": input,
	}

	var result struct {
		CreateBusinessCustomer BusinessCustomer `json:"createBusinessCustomer"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return nil, fmt.Errorf("failed to create business customer: %w", err)
	}

	return &result.CreateBusinessCustomer, nil
}

// CreatePremiumCustomer creates a new premium customer
func (c *GraphQLClient) CreatePremiumCustomer(input CreatePremiumCustomerInput) (*PremiumCustomer, error) {
	query := `
		mutation CreatePremiumCustomer($input: CreatePremiumCustomerInput!) {
			createPremiumCustomer(input: $input) {
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
		"input": input,
	}

	var result struct {
		CreatePremiumCustomer PremiumCustomer `json:"createPremiumCustomer"`
	}

	if err := c.ExecuteWithResult(query, variables, &result); err != nil {
		return nil, fmt.Errorf("failed to create premium customer: %w", err)
	}

	return &result.CreatePremiumCustomer, nil
}

// CreateIndividualCustomerAndPrint creates an individual customer and prints the result
func (c *GraphQLClient) CreateIndividualCustomerAndPrint(name, email, password string, personalInfo *PersonalInfoInput) error {
	fmt.Println("👤 Creating individual customer...")

	input := CreateIndividualCustomerInput{
		Name:         name,
		Email:        email,
		Password:     password,
		PersonalInfo: personalInfo,
	}

	customer, err := c.CreateIndividualCustomer(input)
	if err != nil {
		fmt.Printf("❌ Failed to create individual customer: %v\n", err)
		return err
	}

	fmt.Printf("✅ Individual customer created successfully!\n")
	fmt.Printf("🆔 ID: %s\n", customer.ID)
	fmt.Printf("👤 Name: %s\n", customer.Name)
	fmt.Printf("📧 Email: %s\n", customer.Email)
	fmt.Printf("📅 Created: %s\n", customer.CreatedAt)
	if customer.PersonalInfo != nil {
		fmt.Printf("📞 Phone: %v\n", customer.PersonalInfo.Phone)
		fmt.Printf("🏠 Address: %v\n", customer.PersonalInfo.Address)
		fmt.Printf("🎂 Date of Birth: %v\n", customer.PersonalInfo.DateOfBirth)
	}

	return nil
}

// CreateBusinessCustomerAndPrint creates a business customer and prints the result
func (c *GraphQLClient) CreateBusinessCustomerAndPrint(name, email, password, companyName string, businessInfo *BusinessInfoInput) error {
	fmt.Println("🏢 Creating business customer...")

	input := CreateBusinessCustomerInput{
		Name:         name,
		Email:        email,
		Password:     password,
		CompanyName:  companyName,
		BusinessInfo: businessInfo,
	}

	customer, err := c.CreateBusinessCustomer(input)
	if err != nil {
		fmt.Printf("❌ Failed to create business customer: %v\n", err)
		return err
	}

	fmt.Printf("✅ Business customer created successfully!\n")
	fmt.Printf("🆔 ID: %s\n", customer.ID)
	fmt.Printf("👤 Name: %s\n", customer.Name)
	fmt.Printf("📧 Email: %s\n", customer.Email)
	fmt.Printf("🏢 Company: %s\n", customer.CompanyName)
	fmt.Printf("📅 Created: %s\n", customer.CreatedAt)
	if customer.BusinessInfo != nil {
		fmt.Printf("🏷️ Tax ID: %v\n", customer.BusinessInfo.TaxID)
		fmt.Printf("🏭 Industry: %v\n", customer.BusinessInfo.Industry)
		fmt.Printf("👥 Employee Count: %v\n", customer.BusinessInfo.EmployeeCount)
		fmt.Printf("🌐 Website: %v\n", customer.BusinessInfo.Website)
	}

	return nil
}

// CreatePremiumCustomerAndPrint creates a premium customer and prints the result
func (c *GraphQLClient) CreatePremiumCustomerAndPrint(name, email, password, premiumTier string) error {
	fmt.Println("⭐ Creating premium customer...")

	input := CreatePremiumCustomerInput{
		Name:        name,
		Email:       email,
		Password:    password,
		PremiumTier: premiumTier,
	}

	customer, err := c.CreatePremiumCustomer(input)
	if err != nil {
		fmt.Printf("❌ Failed to create premium customer: %v\n", err)
		return err
	}

	fmt.Printf("✅ Premium customer created successfully!\n")
	fmt.Printf("🆔 ID: %s\n", customer.ID)
	fmt.Printf("👤 Name: %s\n", customer.Name)
	fmt.Printf("📧 Email: %s\n", customer.Email)
	fmt.Printf("⭐ Premium Tier: %s\n", customer.PremiumTier)
	fmt.Printf("📅 Created: %s\n", customer.CreatedAt)
	fmt.Printf("🎁 Benefits: %v\n", customer.Benefits)

	return nil
}
