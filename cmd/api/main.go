package main

import (
	"github.com/abigpotostew/go-api/internal/api"
	"github.com/abigpotostew/go-api/internal/database"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := database.ConnectDatabase(dsn)
	if err != nil {
		println("Error connecting to database")
		panic(err)
	}
	router := gin.Default()
	err = api.NewApiHandler(router, db)
	if err != nil {
		println("Error creating API handler")
		panic(err)
	}
	err = router.Run("localhost:8080")
	if err != nil {
		println("Error running API server")
		panic(err)
	}
}
