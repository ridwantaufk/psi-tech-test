package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ridwantaufk/psi-tech-test/handlers"
	"github.com/ridwantaufk/psi-tech-test/middleware"
)

func Setup(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	api := r.Group("/api")
	{
		api.POST("/checkout", middleware.JWTMiddleware(), handlers.Checkout)
		api.GET("/users", handlers.GetUsers)

		api.GET("/users/external", handlers.GetExternalUsers)
	}
}
