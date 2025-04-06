package http

import (
	"github.com/gin-gonic/gin"
)

func (s *HTTPServer) registerRoutes(engine *gin.Engine) {
	r := engine.Use(s.server.RequireCorrelationID, s.server.RequireRequestID)

	r.GET("/", s.wrapped(s.viewOriginHandler))
}
