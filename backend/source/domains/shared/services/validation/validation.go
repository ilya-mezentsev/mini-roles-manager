package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	responseFactory "github.com/ilya-mezentsev/response-factory"
	sharedError "mini-roles-backend/source/domains/shared/error"
	"strings"
)

func MakeErrorResponse(request interface{}) responseFactory.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return responseFactory.ClientError(sharedError.ServiceError{
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
