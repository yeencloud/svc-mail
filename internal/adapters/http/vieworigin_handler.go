package http

import (
	"github.com/gin-gonic/gin"

	"github.com/yeencloud/bpt-service/internal/domain"
)

func (s *HTTPServer) viewOriginHandler(c *gin.Context) (any, error) {
	origins, err := s.usecases.Viewed(c, domain.ViewOrigin{
		IP:        c.ClientIP(),
		Useragent: c.Request.UserAgent(),
	})

	if err != nil {
		return nil, err
	}

	return origins, nil
}
