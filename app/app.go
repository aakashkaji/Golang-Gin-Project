package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aakashkaji/empolyee-go/app/internal/domain"
	"github.com/aakashkaji/empolyee-go/app/internal/middleware"
	"github.com/aakashkaji/empolyee-go/app/internal/service"
	_ "github.com/aakashkaji/empolyee-go/docs" // Import the generated docs
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	listenAddress string
	db            *sql.DB
}

func NewServer(listenAddress string, db *sql.DB) *Server {
	return &Server{listenAddress: listenAddress, db: db}
}

func (s Server) StartServer() {
	router := gin.Default()
	empolyeeDB := domain.NewEmpolyeeDB(s.db)

	eh := EmpHandler{service: *service.NewEmpolyeeService(*empolyeeDB)}

	EmpolyeeSetupRoutes(router, eh)
	router.Use(gin.LoggerWithFormatter(middleware.Logger))

	log.Fatal(http.ListenAndServe(s.listenAddress, router))

}

func EmpolyeeSetupRoutes(r *gin.Engine, eh EmpHandler) {

	swaggerGroup := r.Group("/swagger", middleware.BasicAuthMiddleware("admin", "password"))

	swaggerGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	empRoutes := r.Group("/api/v1")
	empRoutes.Use(middleware.Authentication())

	{
		empRoutes.GET("/ping", eh.TestHandler)
		empRoutes.GET("/empolyees", eh.GetAllEmp)

		empRoutes.GET("/empolyees/:id", eh.GetEmpById)
		empRoutes.POST("/empolyees", eh.CreateEmpHandler)
		empRoutes.DELETE("/empolyees/:id", eh.DeleteEmpHandler)
		empRoutes.PUT("/empolyees/:id", eh.UpdateEmpHandler)

	}

}
