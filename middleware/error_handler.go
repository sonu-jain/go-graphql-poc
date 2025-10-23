package middleware

import (
	"context"
	"errors"
	"fmt"
	"go-graphql-poc/validator"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ErrorPresenter formats errors in a consistent way
func ErrorPresenter(ctx context.Context, err error) *gqlerror.Error {
	// Get the GraphQL error
	gqlErr := graphql.DefaultErrorPresenter(ctx, err)

	// Check if it's a validation error
	var validationErrs *validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		// Format validation errors with proper structure
		gqlErr.Message = "Validation failed"
		gqlErr.Extensions = map[string]interface{}{
			"code":             "VALIDATION_ERROR",
			"validationErrors": validationErrs.Errors,
		}
		return gqlErr
	}

	// Check for single validation error
	var validationErr validator.ValidationError
	if errors.As(err, &validationErr) {
		gqlErr.Message = "Validation failed"
		gqlErr.Extensions = map[string]interface{}{
			"code": "VALIDATION_ERROR",
			"validationErrors": []validator.ValidationError{
				validationErr,
			},
		}
		return gqlErr
	}

	// Handle database errors
	if isDatabaseError(err) {
		gqlErr.Message = formatDatabaseError(err)
		gqlErr.Extensions = map[string]interface{}{
			"code": getDatabaseErrorCode(err),
		}
		return gqlErr
	}

	// Default error handling
	if gqlErr.Extensions == nil {
		gqlErr.Extensions = make(map[string]interface{})
	}

	// Add default error code if not present
	if _, ok := gqlErr.Extensions["code"]; !ok {
		gqlErr.Extensions["code"] = "INTERNAL_ERROR"
	}

	return gqlErr
}

// isDatabaseError checks if the error is a database-related error
func isDatabaseError(err error) bool {
	errMsg := err.Error()
	return contains(errMsg, "record not found") ||
		contains(errMsg, "duplicate") ||
		contains(errMsg, "constraint") ||
		contains(errMsg, "database")
}

// formatDatabaseError formats database errors into user-friendly messages
func formatDatabaseError(err error) string {
	errMsg := err.Error()

	if contains(errMsg, "record not found") {
		return "The requested resource was not found"
	}

	if contains(errMsg, "duplicate") || contains(errMsg, "unique constraint") {
		return "A record with the same information already exists"
	}

	return "A database error occurred"
}

// getDatabaseErrorCode returns the appropriate error code for database errors
func getDatabaseErrorCode(err error) string {
	errMsg := err.Error()

	if contains(errMsg, "record not found") {
		return "NOT_FOUND"
	}

	if contains(errMsg, "duplicate") || contains(errMsg, "unique constraint") {
		return "DUPLICATE_ENTRY"
	}

	return "DATABASE_ERROR"
}

// contains checks if a string contains a substring (case-insensitive)
func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr || len(substr) == 0 ||
			fmt.Sprintf("%s", s) != "" && len(s) > 0 &&
				stringContains(s, substr))
}

func stringContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
