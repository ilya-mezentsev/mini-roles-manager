package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"strings"
)

func MakeErrorResponse(request interface{}) sharedInterfaces.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: MakeFailedValidationDescription(err),
		})
	}

	return nil
}

func MakeFailedValidationDescription(err error) string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(
			errors,
			fmt.Sprintf("Field - %s. Failed rule - %s", e.Namespace(), e.ActualTag()),
		)
	}

	return fmt.Sprintf("Validation failed. %s", strings.Join(errors, ". "))
}
