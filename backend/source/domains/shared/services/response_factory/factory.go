package response_factory

import "mini-roles-backend/source/domains/shared/interfaces"

func DefaultResponse() interfaces.HttpResponse {
	return defaultResponse{}
}

func SuccessResponse(data interface{}) interfaces.HttpResponse {
	return successResponse{defaultResponse{data}}
}

func ServerError(data interface{}) interfaces.HttpResponse {
	return serverErrorResponse{defaultResponse{data}}
}

func ClientError(data interface{}) interfaces.HttpResponse {
	return clientErrorResponse{defaultResponse{data}}
}

func ForbiddenError(data interface{}) interfaces.HttpResponse {
	return forbiddenErrorResponse{defaultResponse{data}}
}

func UnauthorizedError(data interface{}) interfaces.HttpResponse {
	return unauthorizedErrorResponse{defaultResponse{data}}
}
