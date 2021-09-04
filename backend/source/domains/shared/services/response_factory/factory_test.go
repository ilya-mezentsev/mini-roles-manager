package response_factory

import (
	"github.com/stretchr/testify/assert"
	sharedError "mini-roles-backend/source/domains/shared/error"
	"net/http"
	"testing"
)

func TestDefaultServerError(t *testing.T) {
	response := DefaultServerError()

	assert.True(t, response.HasData())
	assert.Equal(t, http.StatusInternalServerError, response.HttpStatus())
	assert.Equal(t, sharedError.ServerErrorCode, response.Data().(sharedError.ServiceError).Code)
	assert.Equal(t, sharedError.ServerErrorDescription, response.Data().(sharedError.ServiceError).Description)
}
