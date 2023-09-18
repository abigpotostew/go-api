package recipe

import (
	"github.com/abigpotostew/go-api/internal/model"
	"github.com/abigpotostew/go-api/internal/servicedep/db"
	"github.com/gin-gonic/gin"
)

func ListRecipes(c *gin.Context, db *db.ServiceDep) (*[]model.Recipe, error) {
	data := []model.Recipe{}
	err := db.Database.NewSelect().Model(&data).
		//ColumnExpr("lower(name)").
		//Where("? = ?", bun.Ident("id"), "some-id").
		Scan(c.Request.Context())

	return &data, err
}
