package api

import (
	"github.com/abigpotostew/go-api/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func getRecipesHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := []model.Recipe{}
		result := db.Find(&data)
		if result.Error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error})
			return
		}
		c.IndentedJSON(http.StatusOK, data)
	}
}

func postRecipeHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newRecipe *model.Recipe

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := c.BindJSON(&newRecipe); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
			return
		}
		newRecipe.Id = 0
		tx := db.Save(newRecipe)
		if tx.Error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": tx.Error})
			return
		}
		c.IndentedJSON(http.StatusOK, newRecipe)
	}
}

func NewApiHandler(router *gin.Engine, db *gorm.DB) error {
	router.GET("/recipes", getRecipesHandler(db))
	router.POST("/recipes", postRecipeHandler(db))
	return nil
}
