package rest

// @title {{ServiceName}}
// @version 1.0

// @BasePath /api/v1
// @query.collection.format multi

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	"{{GoModule}}/internal/rest/docs"
)

func RegisterDebugRouter(swaggerHost string) http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	p := ginprometheus.NewPrometheus("{{ServiceName}}")
	p.MetricsPath = "/metrics"
	p.Use(r)

	pprof.Register(r, "pprof")

	if swaggerHost != "" {
		docs.SwaggerInfo.Host = swaggerHost
	}

	r.GET("/swagger", func(c *gin.Context) { c.Redirect(301, "swagger/index.html") })
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
