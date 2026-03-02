package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetRouter(conn *pgx.Conn) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Using middlewares to get the database connection for all requests.
	router.Use(func(ctx *gin.Context) {
		ctx.Set("DbConn", conn)
	})

	api := router.Group("/api")

	api.GET("/persons", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"id":    "sd43242ij2094230ßi423",
			"name":  "test user",
			"email": "test@ccf.com",
		})
	})

	return router
}
