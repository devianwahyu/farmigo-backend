package request

import (
	"github.com/devianwahyu/farmigo/model/response"
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=1"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	RoleID   uint   `json:"role_id" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type ChangePasswordRequest struct {
	Email       string `json:"email" validate:"required,email"`
	OldPassword string `json:"old_password" validate:"required,min=6"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

var authValidate = validator.New()

// Input Validations
func ValidateRegisterStruct(registerData RegisterRequest) []*response.ErrorResponse {
	var errors []*response.ErrorResponse
	err := authValidate.Struct(registerData)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element response.ErrorResponse
			element.FailedFields = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateLoginStruct(loginData LoginRequest) []*response.ErrorResponse {
	var errors []*response.ErrorResponse
	err := authValidate.Struct(loginData)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element response.ErrorResponse
			element.FailedFields = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateChangePasswordStruct(changePasswordData ChangePasswordRequest) []*response.ErrorResponse {
	var errors []*response.ErrorResponse
	err := authValidate.Struct(changePasswordData)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element response.ErrorResponse
			element.FailedFields = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
