package client

import (
	"fmt"
	"strings"
)

// CreateCustomerAndPrint creates a customer and prints the result
func CreateCustomerAndPrint(name, email, password string) {
	client := NewGraphQLClientWithToken("http://localhost:8080/query")

	fmt.Println("👤 Creating customer...")

	customer, err := client.CreateIndividualCustomer(CreateIndividualCustomerInput{
		Name:     name,
		Email:    email,
		Password: password,
	})

	if err != nil {
		fmt.Printf("❌ Failed to create customer: %v\n", err)
		return
	}

	fmt.Printf("✅ Customer created successfully!\n")
	fmt.Printf("🆔 ID: %s\n", customer.ID)
	fmt.Printf("👤 Name: %s\n", customer.Name)
	fmt.Printf("📧 Email: %s\n", customer.Email)
	fmt.Printf("📅 Created: %s\n", customer.CreatedAt)
}

// LoginAndPrint performs login and prints the result
func LoginAndPrint(email, password string) {
	client := NewGraphQLClientWithToken("http://localhost:8080/query")

	fmt.Println("🔐 Attempting to login...")

	loginResp, err := client.Login(email, password)
	if err != nil {
		fmt.Printf("❌ Login failed: %v\n", err)
		return
	}

	fmt.Printf("✅ Login successful!\n")
	fmt.Printf("🔑 Token: %s\n", loginResp.Token)
	fmt.Printf("👤 Customer: %+v\n", loginResp.Customer)
}

// DemoAllQueries demonstrates all available queries
func DemoAllQueries() {
	client := NewGraphQLClientWithToken("http://localhost:8080/query")

	fmt.Println("🔍 GraphQL Queries Demo")
	fmt.Println("======================")

	// Get all customers
	fmt.Println("\n1. Getting all customers...")
	client.GetCustomersAndPrint(10, 0)

	// Get customer by ID (assuming ID "1" exists)
	fmt.Println("\n2. Getting customer by ID...")
	client.GetCustomerAndPrint("1")

	// Search customers
	fmt.Println("\n3. Searching customers...")
	client.SearchCustomersAndPrint("test")

	// Get customers by type
	fmt.Println("\n4. Getting customers by type...")
	customers, err := client.GetCustomersByType(CustomerTypeIndividual, 10, 0)
	if err != nil {
		fmt.Printf("❌ Failed to get customers by type: %v\n", err)
	} else {
		fmt.Printf("✅ Found %d individual customers\n", len(customers))
	}

	// Get customers by status
	fmt.Println("\n5. Getting customers by status...")
	customers, err = client.GetCustomersByStatus(CustomerStatusActive, 10, 0)
	if err != nil {
		fmt.Printf("❌ Failed to get customers by status: %v\n", err)
	} else {
		fmt.Printf("✅ Found %d active customers\n", len(customers))
	}

	// Get premium customers by tier
	fmt.Println("\n6. Getting premium customers by tier...")
	premiumCustomers, err := client.GetPremiumCustomersByTier("GOLD", 10, 0)
	if err != nil {
		fmt.Printf("❌ Failed to get premium customers by tier: %v\n", err)
	} else {
		fmt.Printf("✅ Found %d GOLD premium customers\n", len(premiumCustomers))
	}

	fmt.Println("\n✅ Queries demo completed!")
}

// DemoAllMutations demonstrates all available mutations
func DemoAllMutations() {
	client := NewGraphQLClientWithToken("http://localhost:8080/query")

	fmt.Println("✏️ GraphQL Mutations Demo")
	fmt.Println("========================")

	// Create different types of customers
	fmt.Println("\n1. Creating individual customer...")
	individualCustomer, err := client.CreateIndividualCustomer(CreateIndividualCustomerInput{
		Name:     "Individual Test User",
		Email:    "individual@test.com",
		Password: "password123",
		PersonalInfo: &PersonalInfoInput{
			Phone: stringPtr("+1-555-0001"),
		},
	})
	if err != nil {
		fmt.Printf("❌ Failed to create individual customer: %v\n", err)
	} else {
		fmt.Printf("✅ Individual customer created: %s\n", individualCustomer.ID)
	}

	fmt.Println("\n2. Creating business customer...")
	businessCustomer, err := client.CreateBusinessCustomer(CreateBusinessCustomerInput{
		Name:        "Business Test User",
		Email:       "business@test.com",
		Password:    "password123",
		CompanyName: "Test Corp",
		BusinessInfo: &BusinessInfoInput{
			Industry:      stringPtr("Technology"),
			EmployeeCount: int32Ptr(50),
		},
	})
	if err != nil {
		fmt.Printf("❌ Failed to create business customer: %v\n", err)
	} else {
		fmt.Printf("✅ Business customer created: %s\n", businessCustomer.ID)
	}

	fmt.Println("\n3. Creating premium customer...")
	premiumCustomer, err := client.CreatePremiumCustomer(CreatePremiumCustomerInput{
		Name:        "Premium Test User",
		Email:       "premium@test.com",
		Password:    "password123",
		PremiumTier: "PLATINUM",
	})
	if err != nil {
		fmt.Printf("❌ Failed to create premium customer: %v\n", err)
	} else {
		fmt.Printf("✅ Premium customer created: %s\n", premiumCustomer.ID)
	}

	// Update customers
	if individualCustomer != nil {
		fmt.Println("\n4. Updating individual customer...")
		updateInput := UpdateCustomerInput{
			Name: stringPtr("Updated Individual User"),
		}
		client.UpdateCustomerAndPrint(individualCustomer.ID, updateInput)
	}

	if businessCustomer != nil {
		fmt.Println("\n5. Updating business customer...")
		updateInput := UpdateCustomerInput{
			CompanyName: stringPtr("Updated Test Corp"),
		}
		client.UpdateCustomerAndPrint(businessCustomer.ID, updateInput)
	}

	// Create customer with error handling
	fmt.Println("\n6. Creating customer with error handling...")
	client.CreateCustomerWithErrorHandlingAndPrint(CreateIndividualCustomerInput{
		Name:     "Error Test User",
		Email:    "error@test.com",
		Password: "password123",
	})

	// Delete customers (cleanup)
	if individualCustomer != nil {
		fmt.Println("\n7. Deleting individual customer...")
		client.DeleteCustomerAndPrint(individualCustomer.ID)
	}

	if businessCustomer != nil {
		fmt.Println("\n8. Deleting business customer...")
		client.DeleteCustomerAndPrint(businessCustomer.ID)
	}

	if premiumCustomer != nil {
		fmt.Println("\n9. Deleting premium customer...")
		client.DeleteCustomerAndPrint(premiumCustomer.ID)
	}

	fmt.Println("\n✅ Mutations demo completed!")
}

// DemoCompleteWorkflow demonstrates a complete workflow
func DemoCompleteWorkflow() {
	client := NewGraphQLClientWithToken("http://localhost:8080/query")

	fmt.Println("🔄 Complete Workflow Demo")
	fmt.Println("========================")

	// 1. Create a customer
	fmt.Println("\n1. Creating customer...")
	customer, err := client.CreateIndividualCustomer(CreateIndividualCustomerInput{
		Name:     "Workflow Test User",
		Email:    "workflow@test.com",
		Password: "password123",
		PersonalInfo: &PersonalInfoInput{
			Phone:   stringPtr("+1-555-WORKFLOW"),
			Address: stringPtr("123 Workflow St"),
		},
	})
	if err != nil {
		fmt.Printf("❌ Failed to create customer: %v\n", err)
		return
	}
	fmt.Printf("✅ Customer created with ID: %s\n", customer.ID)

	// 2. Login with the customer
	fmt.Println("\n2. Logging in...")
	loginResp, err := client.Login("workflow@test.com", "password123")
	if err != nil {
		fmt.Printf("❌ Login failed: %v\n", err)
	} else {
		fmt.Printf("✅ Login successful! Token: %s\n", loginResp.Token)
	}

	// 3. Get the customer details
	fmt.Println("\n3. Getting customer details...")
	client.GetCustomerAndPrint(customer.ID)

	// 4. Update the customer
	fmt.Println("\n4. Updating customer...")
	updateInput := UpdateCustomerInput{
		Name: stringPtr("Updated Workflow User"),
		PersonalInfo: &PersonalInfoInput{
			Phone:   stringPtr("+1-555-UPDATED"),
			Address: stringPtr("456 Updated St"),
		},
	}
	client.UpdateCustomerAndPrint(customer.ID, updateInput)

	// 5. Search for the customer
	fmt.Println("\n5. Searching for customer...")
	client.SearchCustomersAndPrint("workflow")

	// 6. Get all customers to see the updated one
	fmt.Println("\n6. Getting all customers...")
	client.GetCustomersAndPrint(10, 0)

	// 7. Delete the customer (cleanup)
	fmt.Println("\n7. Cleaning up - deleting customer...")
	client.DeleteCustomerAndPrint(customer.ID)

	fmt.Println("\n✅ Complete workflow demo finished!")
}

// RunAllDemos runs all demonstration functions
func RunAllDemos() {
	fmt.Println("🎯 Running All GraphQL Client Demos")
	fmt.Println("==================================")

	DemoAllQueries()
	fmt.Println("\n" + strings.Repeat("=", 50))
	DemoAllMutations()
	fmt.Println("\n" + strings.Repeat("=", 50))
	DemoCompleteWorkflow()

	fmt.Println("\n🎉 All demos completed!")
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func int32Ptr(i int32) *int32 {
	return &i
}
