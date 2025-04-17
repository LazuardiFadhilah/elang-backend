package main

import (
	"os"

	"github.com/LazuardiFadhilah/elang-backend/config"
	"github.com/LazuardiFadhilah/elang-backend/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()

	router.SetupRoutes(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
