package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logger(params gin.LogFormatterParams) string {

	// Get the request body
	requestBody := params.Request.Body
	fmt.Println("_____", requestBody)

	return fmt.Sprintf("%s | [%s] | %s |%s |%s \n",
		params.ClientIP,
		params.TimeStamp.Format("2006-01-02 15:04:05"),
		params.Method,
		params.Path,
		requestBody,
	)

}

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authtoken := ctx.GetHeader("Authorization")

		if authtoken != "token" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		// If authentication token is valid, proceed to the next middleware or handlers
		ctx.Next()

	}
}

// BasicAuthMiddleware provides basic authentication for Swagger UI
func BasicAuthMiddleware(username, password string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, pass, hasAuth := ctx.Request.BasicAuth()
		if hasAuth && user == "admin" && pass == "admin" {
			ctx.Next()
			return
		}
		ctx.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}
