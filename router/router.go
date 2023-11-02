package router

import (
	"be-62test/db"
	"be-62test/handler"
	"be-62test/repository"
	"be-62test/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	routes := gin.Default()
	db := db.DbConnect()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Add the origins you want to allow
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	routes.Use(cors.New(config));

	businessRepository := repository.NewBusinessRepository(db)
	businessService := service.NewBusinessService(businessRepository)
	businessHandler := handler.NewBusinessHandler(businessService)

	routes.GET("/business", businessHandler.GetBusinesses)
	routes.POST("/business", businessHandler.PostBusiness)
	routes.PUT("/business/:business_id", businessHandler.UpdateBusiness)
	routes.DELETE("/business/:business_id", businessHandler.DeleteBusiness)
	routes.GET("/business/search", businessHandler.SearchBusiness)
	
	routes.Run()
}