package http

import (
	"github.com/gin-gonic/gin"

	"github.com/yeencloud/bpt-service/internal/ports"
	service "github.com/yeencloud/lib-base"
	"github.com/yeencloud/lib-base/transaction"
	httpserver "github.com/yeencloud/lib-httpserver"
)

type HTTPServer struct {
	server *httpserver.HttpServer

	usecases ports.Usecases

	trx transaction.TransactionInterface
}

func NewHTTPServer(server *httpserver.HttpServer, usecases ports.Usecases, trx transaction.TransactionInterface) *HTTPServer {
	httpServer := &HTTPServer{
		server:   server,
		usecases: usecases,
		trx:      trx,
	}

	httpServer.registerRoutes(server.Gin)

	return httpServer
}

func (s *HTTPServer) wrapped(handler service.WrappedHandlerFunc) gin.HandlerFunc {
	return service.HandleWithTransaction(s.server, s.trx, handler)
}
