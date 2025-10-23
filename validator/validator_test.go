package validator

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
		errCode string
	}{
		{"Valid email", "test@example.com", false, ""},
		{"Valid email with subdomain", "user@mail.example.com", false, ""},
		{"Empty email", "", true, "REQUIRED_FIELD"},
		{"Invalid format - no @", "invalid.email.com", true, "INVALID_FORMAT"},
		{"Invalid format - no domain", "test@", true, "INVALID_FORMAT"},
		{"Invalid format - no local part", "@example.com", true, "INVALID_FORMAT"},
		{"Too long email", "verylongemailaddressthatexceedsthemaximumlengthverylongemailaddressthatexceedsthemaximumlengthverylongemailaddressthatexceedsthemaximumlengthverylongemailaddressthatexceedsthemaximumlengthverylongemailaddressthatexceedsthemaximumlengthverylongemailaddressthatexceedsthemaximumlength@example.com", true, "MAX_LENGTH_EXCEEDED"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Code != tt.errCode {
				t.Errorf("ValidateEmail() error code = %v, want %v", err.Code, tt.errCode)
			}
		})
	}
}

func TestValidateName(t *testing.T) {
	tests := []struct {
		name    string
		nameVal string
		wantErr bool
		errCode string
	}{
		{"Valid name", "John Doe", false, ""},
		{"Valid short name", "Jo", false, ""},
		{"Empty name", "", true, "REQUIRED_FIELD"},
		{"Too short name", "J", true, "MIN_LENGTH"},
		{"Too long name", "ThisIsAVeryLongNameThatExceedsTheMaximumLengthAllowedForACustomerNameFieldInTheDatabaseAndShouldFailThisIsAVeryLongNameThatExceedsTheMaximum", true, "MAX_LENGTH_EXCEEDED"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateName(tt.nameVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Code != tt.errCode {
				t.Errorf("ValidateName() error code = %v, want %v", err.Code, tt.errCode)
			}
		})
	}
}

func TestValidateID(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr bool
		errCode string
	}{
		{"Valid ID", "123", false, ""},
		{"Valid single digit ID", "1", false, ""},
		{"Empty ID", "", true, "REQUIRED_FIELD"},
		{"Invalid ID - letters", "abc", true, "INVALID_FORMAT"},
		{"Invalid ID - mixed", "12a34", true, "INVALID_FORMAT"},
		{"Invalid ID - negative", "-123", true, "INVALID_FORMAT"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Code != tt.errCode {
				t.Errorf("ValidateID() error code = %v, want %v", err.Code, tt.errCode)
			}
		})
	}
}

func TestValidateCustomerCreate(t *testing.T) {
	tests := []struct {
		name     string
		nameVal  string
		email    string
		wantErr  bool
		errCount int
	}{
		{"Valid input", "John Doe", "john@example.com", false, 0},
		{"Invalid name only", "J", "john@example.com", true, 1},
		{"Invalid email only", "John Doe", "invalid-email", true, 1},
		{"Both invalid", "J", "invalid-email", true, 2},
		{"Both empty", "", "", true, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCustomerCreate(tt.nameVal, tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCustomerCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				validationErrs, ok := err.(*ValidationErrors)
				if !ok {
					t.Errorf("Expected ValidationErrors type")
					return
				}
				if len(validationErrs.Errors) != tt.errCount {
					t.Errorf("Expected %d errors, got %d", tt.errCount, len(validationErrs.Errors))
				}
			}
		})
	}
}

func TestValidatePagination(t *testing.T) {
	page10 := int32(10)
	page150 := int32(150)
	pageNeg := int32(-5)
	offset0 := int32(0)
	offsetNeg := int32(-10)

	tests := []struct {
		name     string
		page     *int32
		offset   *int32
		wantErr  bool
		errCount int
	}{
		{"Valid pagination", &page10, &offset0, false, 0},
		{"Nil values", nil, nil, false, 0},
		{"Page too large", &page150, nil, true, 1},
		{"Negative page", &pageNeg, nil, true, 1},
		{"Negative offset", nil, &offsetNeg, true, 1},
		{"Both invalid", &pageNeg, &offsetNeg, true, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePagination(tt.page, tt.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePagination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				validationErrs, ok := err.(*ValidationErrors)
				if !ok {
					t.Errorf("Expected ValidationErrors type")
					return
				}
				if len(validationErrs.Errors) != tt.errCount {
					t.Errorf("Expected %d errors, got %d", tt.errCount, len(validationErrs.Errors))
				}
			}
		})
	}
}
