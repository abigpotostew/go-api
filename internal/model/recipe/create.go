package recipe

import (
	"github.com/abigpotostew/go-api/internal/model"
	"github.com/abigpotostew/go-api/internal/servicedep/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRecipe(c *gin.Context, db *db.ServiceDep) (*model.Recipe, error) {
	var newRecipe *model.Recipe

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newRecipe); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return nil, err
	}

	newRecipe.ID = 0
	_, err := db.Database.NewInsert().Model(newRecipe).
		Returning("*").
		Exec(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return nil, err
	}
	return newRecipe, nil
}
