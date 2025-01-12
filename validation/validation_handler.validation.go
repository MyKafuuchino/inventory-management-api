package validation

import (
	"inventory-management/utils"
	"net/http"
)

func ValidationHandler[T any](body T) error {
	err := Validate.Struct(body)
	if err != nil {
		return utils.NewCustomError(http.StatusBadRequest, "Validation error", utils.GetValidationMessages(err)...)
	}
	return nil
}
