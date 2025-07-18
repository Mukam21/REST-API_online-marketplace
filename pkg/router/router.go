package router

import (
	"github.com/Mukam21/REST-API_online-marketplace/pkg/handlers"
	"github.com/Mukam21/REST-API_online-marketplace/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	auth := handlers.NewAuthHandler(db)
	order := handlers.NewOrderHandler(db)

	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

	orderGroup := r.Group("/orders")
	orderGroup.Use(middleware.JWTAuthMiddleware())
	{
		orderGroup.POST("", order.CreateOrder)
	}

	r.GET("/orders", order.GetOrders)

	return r
}
