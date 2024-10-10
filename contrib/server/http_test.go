package server_test

import (
	paruServer "github.com/AreaZero-Database/go-paru/contrib/server"
	"github.com/kataras/iris/v12"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestNewServerWithCustomShutdownTimeout(t *testing.T) {
	app := iris.New()
	server := paruServer.NewServer(app, ":27483", paruServer.WithShutdownTimeout(2*time.Minute))

	assert.Equal(t, server.ShutdownTimeout, 2*time.Minute)
}

func TestStartServerWithInvalidAddress(t *testing.T) {
	app := iris.New()
	server := paruServer.NewServer(app, "invalid address")

	err := server.Start()
	assert.Error(t, err)
}

func TestStartServerAndHandleRequest(t *testing.T) {
	app := iris.New()
	app.Get("/test", func(ctx iris.Context) {
		ctx.StatusCode(http.StatusOK)
		ctx.Writef("Test Endpoint")
	})
	app.Build()

	server := paruServer.NewServer(app, ":27483")
	go func() {
		time.Sleep(2 * time.Second)
		resp, err := http.Get("http://localhost:27483/test")
		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusOK)
		err = server.Shutdown()
		assert.NoError(t, err)
	}()

	err := server.Start()
	assert.NoError(t, err)
}
