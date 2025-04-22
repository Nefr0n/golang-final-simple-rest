package main

import (
	"github.com/Skapar/simple-rest/app"
	_ "github.com/Skapar/simple-rest/docs"
)

// main запускает приложение
// @title Simple REST API
// @version 1.0
// @description Это простое REST API приложение на Go с использованием Gin и GORM
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

func main() {
	// Create a new App
	application := app.NewApp()

	application.Log.Infof("Starting server on port %s", application.Cfg.ListenHttpPort)
	application.Log.Fatal(application.Router.Run(":" + application.Cfg.ListenHttpPort))
}
