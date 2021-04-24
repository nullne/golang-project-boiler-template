package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"{{GoModule}}/internal/{{ServiceName}}"

	"{{GoModule}}/pkg/middleware"
)

func RegisterRouter(c {{ServiceName}}.{{title ServiceName}}) http.Handler {
	h := handler{
		{{ServiceName}}: c,
	}

	r := gin.New()

	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	v1.Use(middleware.Logger())

	v1.GET("hello", h.Hello)
	return r
}
