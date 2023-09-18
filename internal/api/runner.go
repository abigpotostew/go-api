package api

import (
	"fmt"
	"github.com/abigpotostew/go-api/internal/model/recipe"
	"github.com/abigpotostew/go-api/internal/servicedep"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/uptrace/bun/driver/pgdriver"
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

func wsHandler(dep *servicedep.ServiceDep) gin.HandlerFunc {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	return func(c *gin.Context) {
		ln := pgdriver.NewListener(dep.DB.Database)
		if err := ln.Listen(c, "recipes:updated", "recipes:created"); err != nil {
			panic(err)
		}
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "failed to upgrade connection"})
			return
		}
		defer conn.Close()

		for notif := range ln.Channel() {
			msg := fmt.Sprintf("%s: %s", notif.Channel, notif.Payload)
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				fmt.Println("could not write message", err.Error())
				break
			}
		}
	}
}

func NewApiHandler(router *gin.Engine, dep *servicedep.ServiceDep) error {
	router.GET("/recipes", getRecipesHandler(dep))
	router.POST("/recipes", postRecipeHandler(dep))
	router.DELETE("/recipes", deleteRecipeHandler(dep))
	router.GET("/ws", wsHandler(dep))
	return nil
}
