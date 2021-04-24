package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		msg := "Request"
		if len(c.Errors) > 0 {
			msg = c.Errors.String()
		}

		end := time.Now()
		latency := end.Sub(start)

		dumplogger := log.With().
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("client-ip", c.ClientIP()).
			Str("path", path).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Int("body-size", c.Writer.Size()).
			Str("user-agent", c.Request.UserAgent()).
			Logger()

		if c.Request.PostForm != nil {
			dumplogger = dumplogger.With().
				Interface("form", c.Request.Form).
				Logger()
		}
		if len(c.Errors) == 0 {
			dumplogger.Info().
				Msg(msg)
		} else {
			dumplogger.Error().Strs("errors", c.Errors.Errors()).
				Msg(msg)
		}
	}
}
