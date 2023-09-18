package recipe

import (
	"errors"
	"github.com/abigpotostew/go-api/internal/model"
	"github.com/abigpotostew/go-api/internal/servicedep/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRecipe(c *gin.Context, db *db.ServiceDep) error {
	var newRecipe *model.Recipe

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newRecipe); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return err
	}

	res, err := db.Database.NewDelete().Model((*model.Recipe)(nil)).Where("id = ?", newRecipe.ID).Exec(c)
	if err != nil {
		println("error deleting recipe")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	var rows int64
	rows, err = res.RowsAffected()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return err
	}
	if rows != 1 {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "rows affected != 1"})
		return errors.New("rows affected != 1")
	}
	return nil
}
