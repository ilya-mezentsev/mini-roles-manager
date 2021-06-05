package response_factory

import sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"

func DefaultResponse() sharedInterfaces.Response {
	return defaultResponse{}
}

func SuccessResponse(data interface{}) sharedInterfaces.Response {
	return successResponse{defaultResponse{data}}
}

func ServerError(data interface{}) sharedInterfaces.Response {
	return serverErrorResponse{defaultResponse{data}}
}

func EmptyServerError() sharedInterfaces.Response {
	return ServerError(nil)
}

func ClientError(data interface{}) sharedInterfaces.Response {
	return clientErrorResponse{defaultResponse{data}}
}

func EmptyClientError() sharedInterfaces.Response {
	return ClientError(nil)
}

func ForbiddenError(data interface{}) sharedInterfaces.Response {
	return forbiddenErrorResponse{defaultResponse{data}}
}

func UnauthorizedError(data interface{}) sharedInterfaces.Response {
	return unauthorizedErrorResponse{defaultResponse{data}}
}
