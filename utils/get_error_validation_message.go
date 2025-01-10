package utils

import "github.com/go-playground/validator/v10"

func GetErrorValidationMessages(err error) []string {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errors = append(errors, err.Field()+" is required")
		case "email":
			errors = append(errors, err.Field()+" must be a valid email")
		case "min":
			errors = append(errors, err.Field()+" must be at least "+err.Param()+" characters")
		case "max":
			errors = append(errors, err.Field()+" must be at most "+err.Param()+" characters")
		case "gte":
			errors = append(errors, err.Field()+" must be greater than or equal to "+err.Param())
		case "lte":
			errors = append(errors, err.Field()+" must be less than or equal to "+err.Param())
		default:
			errors = append(errors, err.Field()+" is invalid")
		}
	}
	return errors
}
