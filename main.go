package main

import (
	"github.com/SethukumarJ/trx/src/handlers"
	"github.com/SethukumarJ/trx/src/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @title signup/signin form API
// @version 1.0
// @description This is an Event Management project. You can visit the GitHub repository at https://github.com/SethukumarJ/TRX-backend

// @contact.name API Support
// @contact.url sethukumarj.com
// @contact.email sethukumarj.76@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:5000
// @BasePath /
// @query.collection.format multi
func main() {
	engine := gin.Default()

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	middlewareUser := middleware.NewMiddlewareUser()
	user := engine.Group("/user")
	{
		user.POST("/signup", handlers.UserSignup)
		user.POST("/login", handlers.UserLogin)
		user.POST("/token-refresh", handlers.UserRefreshToken)
		user.Use(middlewareUser.AuthorizeJwt())
		{
			user.PATCH("/profile", handlers.UpdateProfile)
			user.GET("/get", handlers.GetUser)
		}
	}

	engine.Run(":5000")
}
