package errors_test

import (
	paruErrors "github.com/AreaZero-Database/go-paru/contrib/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBadRequest(t *testing.T) {
	err := paruErrors.BadRequest("bad request")
	assert.Equal(t, err.Code, 400)
	assert.Equal(t, err.Reason, "bad request")
}

func TestIsBadRequest(t *testing.T) {
	err := paruErrors.New(400, "bad request")
	assert.True(t, paruErrors.IsBadRequest(err))

	err = paruErrors.New(401, "unauthorized")
	assert.False(t, paruErrors.IsBadRequest(err))
}

func TestUnauthorized(t *testing.T) {
	err := paruErrors.Unauthorized("unauthorized")
	assert.Equal(t, err.Code, 401)
	assert.Equal(t, err.Reason, "unauthorized")
}

func TestIsUnauthorized(t *testing.T) {
	err := paruErrors.New(401, "unauthorized")
	assert.True(t, paruErrors.IsUnauthorized(err))

	err = paruErrors.New(403, "forbidden")
	assert.False(t, paruErrors.IsUnauthorized(err))
}

func TestForbidden(t *testing.T) {
	err := paruErrors.Forbidden("forbidden")
	assert.Equal(t, err.Code, 403)
	assert.Equal(t, err.Reason, "forbidden")
}

func TestIsForbidden(t *testing.T) {
	err := paruErrors.New(403, "forbidden")
	assert.True(t, paruErrors.IsForbidden(err))

	err = paruErrors.New(404, "not found")
	assert.False(t, paruErrors.IsForbidden(err))
}

func TestNotFound(t *testing.T) {
	err := paruErrors.NotFound("not found")
	assert.Equal(t, err.Code, 404)
	assert.Equal(t, err.Reason, "not found")
}

func TestIsNotFound(t *testing.T) {
	err := paruErrors.New(404, "not found")
	assert.True(t, paruErrors.IsNotFound(err))

	err = paruErrors.New(409, "conflict")
	assert.False(t, paruErrors.IsNotFound(err))
}

func TestConflict(t *testing.T) {
	err := paruErrors.Conflict("conflict")
	assert.Equal(t, err.Code, 409)
	assert.Equal(t, err.Reason, "conflict")
}

func TestIsConflict(t *testing.T) {
	err := paruErrors.New(409, "conflict")
	assert.True(t, paruErrors.IsConflict(err))

	err = paruErrors.New(500, "internal server error")
	assert.False(t, paruErrors.IsConflict(err))
}

func TestInternalServer(t *testing.T) {
	err := paruErrors.InternalServer("internal server error")
	assert.Equal(t, err.Code, 500)
	assert.Equal(t, err.Reason, "internal server error")
}

func TestIsInternalServer(t *testing.T) {
	err := paruErrors.New(500, "internal server error")
	assert.True(t, paruErrors.IsInternalServer(err))

	err = paruErrors.New(503, "service unavailable")
	assert.False(t, paruErrors.IsInternalServer(err))
}

func TestServiceUnavailable(t *testing.T) {
	err := paruErrors.ServiceUnavailable("service unavailable")
	assert.Equal(t, err.Code, 503)
	assert.Equal(t, err.Reason, "service unavailable")
}

func TestIsServiceUnavailable(t *testing.T) {
	err := paruErrors.New(503, "service unavailable")
	assert.True(t, paruErrors.IsServiceUnavailable(err))

	err = paruErrors.New(504, "gateway timeout")
	assert.False(t, paruErrors.IsServiceUnavailable(err))
}

func TestGatewayTimeout(t *testing.T) {
	err := paruErrors.GatewayTimeout("gateway timeout")
	assert.Equal(t, err.Code, 504)
	assert.Equal(t, err.Reason, "gateway timeout")
}

func TestIsGatewayTimeout(t *testing.T) {
	err := paruErrors.New(504, "gateway timeout")
	assert.True(t, paruErrors.IsGatewayTimeout(err))

	err = paruErrors.New(499, "client closed")
	assert.False(t, paruErrors.IsGatewayTimeout(err))
}

func TestClientClosed(t *testing.T) {
	err := paruErrors.ClientClosed("client closed")
	assert.Equal(t, err.Code, 499)
	assert.Equal(t, err.Reason, "client closed")
}

func TestIsClientClosed(t *testing.T) {
	err := paruErrors.New(499, "client closed")
	assert.True(t, paruErrors.IsClientClosed(err))

	err = paruErrors.New(400, "bad request")
	assert.False(t, paruErrors.IsClientClosed(err))
}
