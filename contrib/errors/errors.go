package errors

import (
	paruErrors "github.com/AreaZero-Database/go-paru/errors"
)

// New create error
func New(code int, reason string) *paruErrors.Error {
	return &paruErrors.Error{Code: code, Reason: reason}
}

// BadRequest new BadRequest error
func BadRequest(reason string) *paruErrors.Error {
	return New(400, reason)
}

// IsBadRequest determines if err is BadRequest error.
func IsBadRequest(err *paruErrors.Error) bool {
	return err.Code == 400
}

// Unauthorized new Unauthorized error
func Unauthorized(reason string) *paruErrors.Error {
	return New(401, reason)
}

// IsUnauthorized determines if err is Unauthorized error.
func IsUnauthorized(err *paruErrors.Error) bool {
	return err.Code == 401
}

// Forbidden new Forbidden error
func Forbidden(reason string) *paruErrors.Error {
	return New(403, reason)
}

// IsForbidden determines if err is Forbidden error.
func IsForbidden(err *paruErrors.Error) bool {
	return err.Code == 403
}

// NotFound new NotFound error
func NotFound(reason string) *paruErrors.Error {
	return New(404, reason)
}

// IsNotFound determines if err is NotFound error.
func IsNotFound(err *paruErrors.Error) bool {
	return err.Code == 404
}

// Conflict new Conflict error
func Conflict(reason string) *paruErrors.Error {
	return New(409, reason)
}

// IsConflict determines if err is Conflict error.
func IsConflict(err *paruErrors.Error) bool {
	return err.Code == 409
}

// InternalServer new InternalServer error
func InternalServer(reason string) *paruErrors.Error {
	return New(500, reason)
}

// IsInternalServer determines if err is InternalServer error.
func IsInternalServer(err *paruErrors.Error) bool {
	return err.Code == 500
}

// ServiceUnavailable new ServiceUnavailable error
func ServiceUnavailable(reason string) *paruErrors.Error {
	return New(503, reason)
}

// IsServiceUnavailable determines if err is ServiceUnavailable error.
func IsServiceUnavailable(err *paruErrors.Error) bool {
	return err.Code == 503
}

// GatewayTimeout new GatewayTimeout error
func GatewayTimeout(reason string) *paruErrors.Error {
	return New(504, reason)
}

// IsGatewayTimeout determines if err is GatewayTimeout error.
func IsGatewayTimeout(err *paruErrors.Error) bool {
	return err.Code == 504
}

// ClientClosed new ClientClosed error
func ClientClosed(reason string) *paruErrors.Error {
	return New(499, reason)
}

// IsClientClosed determines if err is ClientClosed error.
func IsClientClosed(err *paruErrors.Error) bool {
	return err.Code == 499
}
