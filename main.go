package main

import (
	"fmt"
	"os"

	"github.com/aakashkaji/empolyee-go/app"
	"github.com/aakashkaji/empolyee-go/config"
)

// @title           Empolyee Example API
// @version         1.0.0
// @description     This is a sample empolyee Api.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1/

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	env := os.Getenv("ENVIRONMENT")

	fmt.Println(env, "envv")

	fmt.Println("Start gin server")
	db := config.InitDB()
	// swag init
	server := app.NewServer(":8000", db)

	server.StartServer()

}
