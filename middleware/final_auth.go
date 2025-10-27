package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"go-graphql-poc/auth"
)

// FinalAuthMiddleware creates a simple and reliable authentication middleware
func FinalAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only apply to GraphQL requests
		if r.URL.Path != "/query" {
			next.ServeHTTP(w, r)
			return
		}

		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, `{"errors":[{"message":"Failed to read request body"}]}`, http.StatusBadRequest)
			return
		}

		// Restore the body for the next handler
		r.Body = io.NopCloser(strings.NewReader(string(body)))

		query := string(body)

		// Debug logging
		fmt.Printf("DEBUG: Received query: %s\n", query)
		fmt.Printf("DEBUG: Request headers: %v\n", r.Header)

		// Check if this is a public operation
		if isPublicQuery(query) {
			fmt.Printf("DEBUG: Allowing public query\n")
			next.ServeHTTP(w, r)
			return
		}

		fmt.Printf("DEBUG: Query requires authentication\n")

		// For protected operations, validate the JWT token
		token := extractTokenFromHeader(r)
		if token == "" {
			http.Error(w, `{"errors":[{"message":"Authorization token required","extensions":{"code":"UNAUTHENTICATED"}}]}`, http.StatusUnauthorized)
			return
		}

		// Validate the token
		claims, err := auth.ValidateToken(token)
		if err != nil {
			http.Error(w, `{"errors":[{"message":"Invalid or expired token","extensions":{"code":"UNAUTHENTICATED"}}]}`, http.StatusUnauthorized)
			return
		}

		// Add user information to the request context
		ctx := context.WithValue(r.Context(), "user_id", claims.CustomerID)
		ctx = context.WithValue(ctx, "user_email", claims.Email)

		// Continue with the authenticated request
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// isPublicQuery checks if the query is public (doesn't require authentication)
func isPublicQuery(query string) bool {
	// Check for introspection queries - but be more specific
	if strings.Contains(query, "__schema") ||
		strings.Contains(query, "__type") ||
		strings.Contains(query, "__typename") ||
		strings.Contains(query, "__directive") ||
		strings.Contains(query, "__field") ||
		strings.Contains(query, "__inputValue") ||
		strings.Contains(query, "__enumValue") {
		// Only allow if it's ONLY introspection, not mixed with other queries
		if isOnlyIntrospection(query) {
			return true
		}
	}

	// Check for public operations
	publicOps := []string{
		"createIndividualCustomer",
		"createBusinessCustomer",
		"createPremiumCustomer",
		"createCustomerWithErrorHandling",
		"login",
	}

	for _, op := range publicOps {
		if strings.Contains(query, op) {
			return true
		}
	}

	return false
}

// isOnlyIntrospection checks if the query contains only introspection fields
func isOnlyIntrospection(query string) bool {
	// Remove whitespace and newlines for easier checking
	cleanQuery := strings.ReplaceAll(strings.ReplaceAll(query, " ", ""), "\n", "")

	// Check if the query contains any non-introspection operations
	nonIntrospectionOps := []string{
		"customers",
		"customer",
		"customersByType",
		"searchCustomers",
		"getCustomerWithErrorHandling",
		"customersByStatus",
		"premiumCustomersByTier",
		"updateCustomer",
		"deleteCustomer",
	}

	for _, op := range nonIntrospectionOps {
		if strings.Contains(cleanQuery, op) {
			return false
		}
	}

	return true
}

// extractTokenFromHeader extracts the JWT token from the Authorization header
func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	// Check if it's a Bearer token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

// GetUserIDFromContext extracts the user ID from the request context
func GetUserIDFromContext(ctx context.Context) (uint, error) {
	userID, ok := ctx.Value("user_id").(uint)
	if !ok {
		return 0, fmt.Errorf("user not authenticated")
	}
	return userID, nil
}

// GetUserEmailFromContext extracts the user email from the request context
func GetUserEmailFromContext(ctx context.Context) (string, error) {
	email, ok := ctx.Value("user_email").(string)
	if !ok {
		return "", fmt.Errorf("user not authenticated")
	}
	return email, nil
}
