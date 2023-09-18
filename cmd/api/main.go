package main

import (
	"github.com/abigpotostew/go-api/internal/api"
	"github.com/abigpotostew/go-api/internal/servicedep"
	"github.com/gin-gonic/gin"
)

func main() {
	deps := servicedep.NewServiceDep()
	router := gin.Default()
	err := api.NewApiHandler(router, deps)
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
