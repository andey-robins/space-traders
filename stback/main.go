package main

import (
	"github.com/andey-robins/space-traders/stback/controllers"
	"github.com/andey-robins/space-traders/stback/repositories"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// first we need to define the router completely, because it is used when we
	// register controllers into the router
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"X-Requested-With",
			"Accept",
			"X-CSRF-Token",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Content-Type",
		},
	}))

	// firm hierarchy of repo < service < controller < model
	// so ensure we build dependencies in that order.
	repositories.RegisterRepositories()
	// we don't register services because they are owned by each controller
	// and loaded by the controller on creation
	router = controllers.RegisterControllers(router)

	// everything is set up now, we can run the backend
	router.Run(":8080")

	// unwind objects and models in the reverse order as above
	repositories.CloseDb()
}
