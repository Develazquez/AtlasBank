package core

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
)


func SetupCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Orígenes permitidos
		allowedOrigins := []string{
			"http://localhost:3000",
			"http://localhost:5173",
			"http://localhost:4200",
			"http://localhost:8080",
			os.Getenv("FRONTEND_URL"),
		}

		origin := c.Request.Header.Get("Origin")
		for _, allowed := range allowedOrigins {
			if origin == allowed && allowed != "" {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Max-Age", (12 * time.Hour).String())

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
