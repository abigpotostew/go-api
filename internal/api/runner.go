package api

import (
	"github.com/abigpotostew/go-api/internal/model/recipe"
	"github.com/abigpotostew/go-api/internal/servicedep"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getRecipesHandler(dep *servicedep.ServiceDep) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := recipe.ListRecipes(c, dep.DB)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.IndentedJSON(http.StatusOK, data)
	}
}

func postRecipeHandler(dep *servicedep.ServiceDep) gin.HandlerFunc {
	return func(c *gin.Context) {
		newRecipe, err := recipe.CreateRecipe(c, dep.DB)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.IndentedJSON(http.StatusOK, newRecipe)
	}
}

func deleteRecipeHandler(dep *servicedep.ServiceDep) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := recipe.DeleteRecipe(c, dep.DB)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func NewApiHandler(router *gin.Engine, dep *servicedep.ServiceDep) error {
	router.GET("/recipes", getRecipesHandler(dep))
	router.POST("/recipes", postRecipeHandler(dep))
	router.DELETE("/recipes", deleteRecipeHandler(dep))
	return nil
}
