package http

import (
	"github.com/kataras/iris/v12"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.StatusCode(http.StatusOK)
		ctx.Writef("Hello, World!")
	})
	app.Build()

	server := NewServer(app, ":27483", func(server *Server) {
		server.ShutdownTimeout = time.Second
	})

	go func() {
		time.Sleep(2 * time.Second)

		resp, err := http.Get("http://localhost:27483")
		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusOK)

		time.Sleep(1 * time.Second)
		err = server.Shutdown()
		assert.NoError(t, err)
		t.Log("shutdown completed")
	}()

	_ = server.Start()
}
