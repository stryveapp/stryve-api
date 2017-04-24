package router

import (
	"github.com/labstack/echo"
	"github.com/stryveapp/stryve-api/controller"
	"github.com/stryveapp/stryve-api/database"
)

// RegisterRoutes registers all API routes for the app
func RegisterRoutes(e *echo.Echo) {
	h := &controller.Handler{
		DB: database.NewConnection(),
	}

	// AUTH ROUTES
	e.POST("/auth/login", h.Login)
	e.GET("/auth/logout", h.Logout)
	e.POST("/auth/register", h.Register)

	// USER ROUTES
	e.GET("/v1/users", h.GetUsers)
	e.GET("/v1/users/:id", h.GetUser)
	e.POST("/v1/users", h.CreateUser)
	e.PUT("/v1/users/:id", h.UpdateUser)
	e.DELETE("/v1/users/:id", h.DeleteUser)

	// COMMUNITY ROUTES
	e.GET("/v1/communities", h.GetCommunities)
	e.GET("/v1/communities/:id", h.GetCommunity)
	e.POST("/v1/communities", h.CreateCommunity)
	e.PUT("/v1/communities/:id", h.UpdateCommunity)
	e.DELETE("/v1/communities/:id", h.DeleteCommunity)
}
