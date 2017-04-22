package router

import (
	"github.com/labstack/echo"
	"github.com/stryveapp/stryve-api/controller"
)

// RegisterRoutes registers all API routes for the app
func RegisterRoutes(svr *echo.Echo) {
	// AUTH ROUTES
	svr.POST("/auth/login", controller.Login)
	svr.GET("/auth/logout", controller.Logout)
	svr.POST("/auth/register", controller.Register)

	// USER ROUTES
	svr.GET("/v1/users", controller.GetUsers)
	svr.GET("/v1/users/:id", controller.GetUser)
	svr.POST("/v1/users", controller.CreateUser)
	svr.PUT("/v1/users/:id", controller.UpdateUser)
	svr.DELETE("/v1/users/:id", controller.DeleteUser)

	// COMMUNITY ROUTES
	svr.GET("/v1/communities", controller.GetCommunities)
	svr.GET("/v1/communities/:id", controller.GetCommunity)
	svr.POST("/v1/communities", controller.CreateCommunity)
	svr.PUT("/v1/communities/:id", controller.UpdateCommunity)
	svr.DELETE("/v1/communities/:id", controller.DeleteCommunity)
}
