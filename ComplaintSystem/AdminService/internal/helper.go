package internal

import (
	"ComplaintSystem/AdminService/internal/types"
	_ "ComplaintSystem/AdminService/internal/types"
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
	"unicode"
)

func ToAdminResponse(admin *types.Admin) *types.AdminResponseModel {
	return &types.AdminResponseModel{
		Id:           admin.Id,
		CompanyName:  admin.CompanyName,
		Category:     admin.Category,
		Email:        admin.Email,
		Phone:        admin.Phone,
		Address:      admin.Address,
		EmployeeRole: admin.EmployeeRole,
		EmployeeId:   admin.EmployeeId,
	}
}
func ValidatePhone(admin *types.AdminRequestModel) error {

	phone := admin.Phone

	if phone == "" {
		return errors.New("Phone is required")
	}

	if !strings.HasPrefix(phone, "5") {

		return errors.New("Phone must start 5")
	}
	return nil
}

func ValidateCompanyNameFormat(admin *types.AdminRequestModel) error {
	errors := make(map[string]string)
	if admin.CompanyName != "" {
		if !isFirstLetterUpperCase(admin.CompanyName) {
			errors["CompanyName"] = "Company Name must start with an uppercase letter"
		}
		if !areRemainingLettersLowerCase(admin.CompanyName) {
			errors["CompanyName"] = "All characters after the first one must be lowercase"
		}
		if containsDigit(admin.CompanyName) {
			errors["CompanyName"] = "Company name contains a number"
		}
	}
	if len(errors) > 0 {
		//return fmt.Errorf("validation errors: %v", errors)
		return &ValidationError{Errors: errors}
	}

	return nil
}
func containsDigit(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return true
		}
	}
	return false
}
func areRemainingLettersLowerCase(s string) bool {
	for _, c := range s[1:] {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

func isFirstLetterUpperCase(s string) bool {
	if len(s) > 0 {
		return strings.ToUpper(s[:1]) == s[:1]
	}
	return false
}
func ValidateAdmin(admin *types.AdminRequestModel, validate *validator.Validate) error {
	validationErrors := make(map[string]string)

	if err := ValidatePhone(admin); err != nil {
		validationErrors["Age"] = err.Error()
	}

	if err := ValidateCompanyNameFormat(admin); err != nil {
		// Use the errors from ValidateFirstLetterUpperCase directly
		if valErr, ok := err.(*ValidationError); ok {
			for field, msg := range valErr.Errors {
				validationErrors[field] = msg
			}
		}
	}

	if err := validate.Struct(admin); err != nil {
		if fieldErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range fieldErrors {
				if fieldError.Tag() == "required" {
					validationErrors[fieldError.Field()] = "This field is required"
				}

			}
		}
	}

	if len(validationErrors) > 0 {
		return &ValidationError{Errors: validationErrors}
	}

	return nil
}

// Custom validation error structure
type ValidationError struct {
	Errors map[string]string
}

func (e *ValidationError) Error() string {
	return "Validation failed "
}
