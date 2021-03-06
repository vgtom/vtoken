package routes

import (
	"github.com/vgtom/vtoken/api/controller"
	"github.com/vgtom/vtoken/lib"
	"github.com/vgtom/vtoken/middlewares"
)

type PublicRoutes struct {
	handler         lib.RequestHandler
	tokenController controller.ITokenController
	rateLimiter     middlewares.IRateLimiter
}

func (s PublicRoutes) Setup() {
	publicRoutes := s.handler.Gin.Group("/api/v1/public")
	publicRoutes.Use(s.rateLimiter.RateLimit)
	{
		publicRoutes.POST("/token/validate", s.tokenController.ValidateToken)
	}
}

func NewPublicRoutes(handler lib.RequestHandler, tokenController controller.ITokenController, rateLimiter middlewares.IRateLimiter) PublicRoutes {
	return PublicRoutes{
		handler:         handler,
		tokenController: tokenController,
		rateLimiter:     rateLimiter,
	}
}
