package error_handler

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Execution struct {
	Message	string
	StatusCode int
	Details string
}

func (e *Execution) Error() string {
	return e.Message
}

func (e *Execution) GqlError(ctx context.Context) *gqlerror.Error {
	return &gqlerror.Error{
		Message: e.Message,
		Path: graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"details": e.Details,
			"status_code": e.StatusCode,
		},
	}
}

func NewValidationError(  message string, details error, statusCode int) *Execution {
	return &Execution{
		Message	: message,
		StatusCode : statusCode,
		Details :	details.Error(),
	}
}