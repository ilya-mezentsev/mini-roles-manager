package response_factory

import "mini-roles-backend/source/domains/shared/interfaces"

func DefaultResponse() interfaces.Response {
	return defaultResponse{}
}

func SuccessResponse(data interface{}) interfaces.Response {
	return successResponse{defaultResponse{data}}
}

func ServerError(data interface{}) interfaces.Response {
	return serverErrorResponse{defaultResponse{data}}
}

func ClientError(data interface{}) interfaces.Response {
	return clientErrorResponse{defaultResponse{data}}
}

func ForbiddenError(data interface{}) interfaces.Response {
	return forbiddenErrorResponse{defaultResponse{data}}
}

func UnauthorizedError(data interface{}) interfaces.Response {
	return unauthorizedErrorResponse{defaultResponse{data}}
}
