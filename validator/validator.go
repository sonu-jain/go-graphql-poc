package validator

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidationError represents a structured validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors holds multiple validation errors
type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

func (e ValidationErrors) Error() string {
	var messages []string
	for _, err := range e.Errors {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
}

// NewValidationError creates a new validation error
func NewValidationError(field, message, code string) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
		Code:    code,
	}
}

// NewValidationErrors creates a new validation errors collection
func NewValidationErrors(errors ...ValidationError) *ValidationErrors {
	return &ValidationErrors{Errors: errors}
}

// Email validation
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail validates email format
func ValidateEmail(email string) *ValidationError {
	if email == "" {
		return &ValidationError{
			Field:   "email",
			Message: "Email is required",
			Code:    "REQUIRED_FIELD",
		}
	}

	if len(email) > 255 {
		return &ValidationError{
			Field:   "email",
			Message: "Email must not exceed 255 characters",
			Code:    "MAX_LENGTH_EXCEEDED",
		}
	}

	if !emailRegex.MatchString(email) {
		return &ValidationError{
			Field:   "email",
			Message: "Invalid email format",
			Code:    "INVALID_FORMAT",
		}
	}

	return nil
}

// ValidateName validates name field
func ValidateName(name string) *ValidationError {
	if name == "" {
		return &ValidationError{
			Field:   "name",
			Message: "Name is required",
			Code:    "REQUIRED_FIELD",
		}
	}

	if len(name) < 2 {
		return &ValidationError{
			Field:   "name",
			Message: "Name must be at least 2 characters long",
			Code:    "MIN_LENGTH",
		}
	}

	if len(name) > 100 {
		return &ValidationError{
			Field:   "name",
			Message: "Name must not exceed 100 characters",
			Code:    "MAX_LENGTH_EXCEEDED",
		}
	}

	return nil
}

// ValidateID validates ID field
func ValidateID(id string) *ValidationError {
	if id == "" {
		return &ValidationError{
			Field:   "id",
			Message: "ID is required",
			Code:    "REQUIRED_FIELD",
		}
	}

	if !regexp.MustCompile(`^\d+$`).MatchString(id) {
		return &ValidationError{
			Field:   "id",
			Message: "ID must be a valid number",
			Code:    "INVALID_FORMAT",
		}
	}

	return nil
}

// ValidateCustomerCreate validates customer creation input
func ValidateCustomerCreate(name, email string) error {
	var errors []ValidationError

	if err := ValidateName(name); err != nil {
		errors = append(errors, *err)
	}

	if err := ValidateEmail(email); err != nil {
		errors = append(errors, *err)
	}

	if len(errors) > 0 {
		return NewValidationErrors(errors...)
	}

	return nil
}

// ValidateCustomerUpdate validates customer update input
func ValidateCustomerUpdate(id string, name, email *string) error {
	var errors []ValidationError

	if err := ValidateID(id); err != nil {
		errors = append(errors, *err)
	}

	// At least one field must be provided for update
	if name == nil && email == nil {
		errors = append(errors, ValidationError{
			Field:   "input",
			Message: "At least one field (name or email) must be provided for update",
			Code:    "MISSING_UPDATE_FIELDS",
		})
	}

	if name != nil {
		if err := ValidateName(*name); err != nil {
			errors = append(errors, *err)
		}
	}

	if email != nil {
		if err := ValidateEmail(*email); err != nil {
			errors = append(errors, *err)
		}
	}

	if len(errors) > 0 {
		return NewValidationErrors(errors...)
	}

	return nil
}

// ValidatePagination validates pagination parameters
func ValidatePagination(page, offset *int32) error {
	var errors []ValidationError

	if page != nil && *page < 0 {
		errors = append(errors, ValidationError{
			Field:   "page",
			Message: "Page limit must be a positive number",
			Code:    "INVALID_VALUE",
		})
	}

	if page != nil && *page > 100 {
		errors = append(errors, ValidationError{
			Field:   "page",
			Message: "Page limit must not exceed 100",
			Code:    "MAX_VALUE_EXCEEDED",
		})
	}

	if offset != nil && *offset < 0 {
		errors = append(errors, ValidationError{
			Field:   "offset",
			Message: "Offset must be a non-negative number",
			Code:    "INVALID_VALUE",
		})
	}

	if len(errors) > 0 {
		return NewValidationErrors(errors...)
	}

	return nil
}
