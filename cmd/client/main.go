package main

import (
	"flag"
	"fmt"
	"go-graphql-poc/client"
	"os"
)

func main() {
	var (
		action       = flag.String("action", "create", "Action to perform")
		name         = flag.String("name", "Test User", "Customer name")
		email        = flag.String("email", "test@example.com", "Email address")
		password     = flag.String("password", "password123", "Password")
		id           = flag.String("id", "", "Customer ID (for get, update, delete)")
		query        = flag.String("query", "", "Search query")
		customerType = flag.String("type", "INDIVIDUAL", "Customer type (INDIVIDUAL, BUSINESS, PREMIUM)")
		status       = flag.String("status", "ACTIVE", "Customer status (ACTIVE, INACTIVE, SUSPENDED, PENDING)")
		tier         = flag.String("tier", "GOLD", "Premium tier")
		page         = flag.Int("page", 10, "Page size")
		offset       = flag.Int("offset", 0, "Offset")
		help         = flag.Bool("help", false, "Show help")
	)
	flag.Parse()

	if *help {
		showHelp()
		return
	}

	fmt.Printf("🚀 GraphQL Client - %s\n", *action)
	fmt.Println("=========================")

	graphqlClient := client.NewGraphQLClientWithToken("http://localhost:8080/query")

	switch *action {
	case "create":
		client.CreateCustomerAndPrint(*name, *email, *password)
	case "login":
		client.LoginAndPrint(*email, *password)
	case "get":
		if *id == "" {
			fmt.Println("❌ ID is required for get action")
			os.Exit(1)
		}
		graphqlClient.GetCustomerAndPrint(*id)
	case "get-all":
		graphqlClient.GetCustomersAndPrint(*page, *offset)
	case "search":
		if *query == "" {
			fmt.Println("❌ Query is required for search action")
			os.Exit(1)
		}
		graphqlClient.SearchCustomersAndPrint(*query)
	case "get-by-type":
		customerType := client.CustomerType(*customerType)
		customers, err := graphqlClient.GetCustomersByType(customerType, *page, *offset)
		if err != nil {
			fmt.Printf("❌ Failed to get customers by type: %v\n", err)
		} else {
			fmt.Printf("✅ Found %d customers of type %s\n", len(customers), string(customerType))
			for i, customer := range customers {
				fmt.Printf("  %d. %+v\n", i+1, customer)
			}
		}
	case "get-by-status":
		status := client.CustomerStatus(*status)
		customers, err := graphqlClient.GetCustomersByStatus(status, *page, *offset)
		if err != nil {
			fmt.Printf("❌ Failed to get customers by status: %v\n", err)
		} else {
			fmt.Printf("✅ Found %d customers with status %s\n", len(customers), string(status))
			for i, customer := range customers {
				fmt.Printf("  %d. %+v\n", i+1, customer)
			}
		}
	case "get-premium-by-tier":
		customers, err := graphqlClient.GetPremiumCustomersByTier(*tier, *page, *offset)
		if err != nil {
			fmt.Printf("❌ Failed to get premium customers by tier: %v\n", err)
		} else {
			fmt.Printf("✅ Found %d premium customers with tier %s\n", len(customers), *tier)
			for i, customer := range customers {
				fmt.Printf("  %d. %+v\n", i+1, customer)
			}
		}
	case "update":
		if *id == "" {
			fmt.Println("❌ ID is required for update action")
			os.Exit(1)
		}
		updateInput := client.UpdateCustomerInput{
			Name:  stringPtr(*name),
			Email: stringPtr(*email),
		}
		graphqlClient.UpdateCustomerAndPrint(*id, updateInput)
	case "delete":
		if *id == "" {
			fmt.Println("❌ ID is required for delete action")
			os.Exit(1)
		}
		graphqlClient.DeleteCustomerAndPrint(*id)
	case "create-business":
		graphqlClient.CreateBusinessCustomerAndPrint(*name, *email, *password, "Test Company", nil)
	case "create-premium":
		graphqlClient.CreatePremiumCustomerAndPrint(*name, *email, *password, *tier)
	case "demo-queries":
		client.DemoAllQueries()
	case "demo-mutations":
		client.DemoAllMutations()
	case "demo-workflow":
		client.DemoCompleteWorkflow()
	case "demo-all":
		client.RunAllDemos()
	case "logout":
		if err := client.ClearToken(); err != nil {
			fmt.Printf("❌ Failed to logout: %v\n", err)
		} else {
			fmt.Println("✅ Logged out successfully")
		}
	default:
		fmt.Printf("Unknown action: %s\n", *action)
		showHelp()
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Println("GraphQL Client - Complete API Testing")
	fmt.Println("====================================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run main.go [options]")
	fmt.Println()
	fmt.Println("Actions:")
	fmt.Println("  -action string")
	fmt.Println("        Action to perform (default: create)")
	fmt.Println()
	fmt.Println("Basic Actions:")
	fmt.Println("  create              - Create individual customer")
	fmt.Println("  login               - Login with email/password")
	fmt.Println("  logout              - Logout and clear saved token")
	fmt.Println("  get                 - Get customer by ID")
	fmt.Println("  get-all             - Get all customers")
	fmt.Println("  search              - Search customers")
	fmt.Println("  update              - Update customer")
	fmt.Println("  delete              - Delete customer")
	fmt.Println()
	fmt.Println("Advanced Actions:")
	fmt.Println("  get-by-type         - Get customers by type")
	fmt.Println("  get-by-status       - Get customers by status")
	fmt.Println("  get-premium-by-tier - Get premium customers by tier")
	fmt.Println("  create-business     - Create business customer")
	fmt.Println("  create-premium      - Create premium customer")
	fmt.Println()
	fmt.Println("Demo Actions:")
	fmt.Println("  demo-queries        - Demo all queries")
	fmt.Println("  demo-mutations      - Demo all mutations")
	fmt.Println("  demo-workflow       - Demo complete workflow")
	fmt.Println("  demo-all            - Run all demos")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -name string")
	fmt.Println("        Customer name (default: Test User)")
	fmt.Println("  -email string")
	fmt.Println("        Email address (default: test@example.com)")
	fmt.Println("  -password string")
	fmt.Println("        Password (default: password123)")
	fmt.Println("  -id string")
	fmt.Println("        Customer ID (for get, update, delete)")
	fmt.Println("  -query string")
	fmt.Println("        Search query (for search)")
	fmt.Println("  -type string")
	fmt.Println("        Customer type: INDIVIDUAL, BUSINESS, PREMIUM (default: INDIVIDUAL)")
	fmt.Println("  -status string")
	fmt.Println("        Customer status: ACTIVE, INACTIVE, SUSPENDED, PENDING (default: ACTIVE)")
	fmt.Println("  -tier string")
	fmt.Println("        Premium tier (default: GOLD)")
	fmt.Println("  -page int")
	fmt.Println("        Page size (default: 10)")
	fmt.Println("  -offset int")
	fmt.Println("        Offset (default: 0)")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  # Basic operations")
	fmt.Println("  go run main.go -action create -name \"John Doe\" -email \"john@example.com\"")
	fmt.Println("  go run main.go -action login -email \"john@example.com\" -password \"mypassword\"")
	fmt.Println("  go run main.go -action get -id \"1\"")
	fmt.Println("  go run main.go -action search -query \"john\"")
	fmt.Println()
	fmt.Println("  # Advanced operations")
	fmt.Println("  go run main.go -action get-by-type -type INDIVIDUAL")
	fmt.Println("  go run main.go -action get-by-status -status ACTIVE")
	fmt.Println("  go run main.go -action create-business -name \"Company\" -email \"company@test.com\"")
	fmt.Println("  go run main.go -action update -id \"1\" -name \"Updated Name\"")
	fmt.Println("  go run main.go -action delete -id \"1\"")
	fmt.Println()
	fmt.Println("  # Demo operations")
	fmt.Println("  go run main.go -action demo-queries")
	fmt.Println("  go run main.go -action demo-mutations")
	fmt.Println("  go run main.go -action demo-workflow")
	fmt.Println("  go run main.go -action demo-all")
	fmt.Println()
}

func stringPtr(s string) *string {
	return &s
}
