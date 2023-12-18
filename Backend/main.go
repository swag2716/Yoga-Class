package main

import (
	"log"
	"os"
	"yoga-class/controllers"
	"yoga-class/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	// Connect to MongoDB
	database.DBinstance()

	// Initialize Gin
	router := gin.New()
	router.Use(gin.Logger())

	// Routes
	router.POST("enroll", controllers.EnrollController)
	router.GET("participants", controllers.GetParticipantsController)
	router.PATCH("update/{id}", controllers.UpdateBatchController)

	// Run the server
	router.Run(":" + port)
}
