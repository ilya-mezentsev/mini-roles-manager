package response_factory

import (
	responseFactory "github.com/ilya-mezentsev/response-factory"
	sharedError "mini-roles-backend/source/domains/shared/error"
)

func DefaultServerError() responseFactory.Response {
	return responseFactory.ServerError(sharedError.ServiceError{
		Code:        sharedError.ServerErrorCode,
		Description: sharedError.ServerErrorDescription,
	})
}
