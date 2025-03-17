package router

import (
	"DobryySoul/project-with-API-interaction/internal/http/handler"
	"DobryySoul/project-with-API-interaction/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *handler.Handlers) {
	r.Use(gin.Recovery())
	// r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
	// 	Formatter: func(p gin.LogFormatterParams) string {
	// 		return fmt.Sprintf("[%s] %s %s %d %s %s\n",
	// 			p.TimeStamp.Format(time.RFC1123),
	// 			p.ClientIP,
	// 			p.Method,
	// 			p.StatusCode,
	// 			p.Latency,
	// 			p.Request.UserAgent(),
	// 		)
	// 	},
	// }))

	r.POST("/register", h.RegisterUser)
	r.POST("/login", h.LoginUser)
	r.GET("/protected/info", middlewares.AuthMiddleware(h.Cfg.JWTSecret), func(c *gin.Context) {
		c.Request.Header.Set("Authorization", "Bearer "+c.GetHeader("Authorization"))
		c.JSON(200, gin.H{"message": "protected route"})
	})
}
