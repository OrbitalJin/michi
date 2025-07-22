package main

import (
	"net/http"

	"github.com/OrbitalJin/pow/cmd/app"
	"github.com/OrbitalJin/pow/internal/parser"
	"github.com/OrbitalJin/pow/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	parserConfig := parser.NewConfig(`!(\b\w+\b)`, `!\b\w+\b`)
	storeConfig := store.NewConfig("./database.db")
	appConfig := app.NewConfig(parserConfig, storeConfig)
	app := app.New(appConfig)

	app.Router.POST("/collect", func(c *gin.Context) {
		req := c.Query("value")
		result, err := app.Parser.Collect(req)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "error")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"matches": result.Matches,
			"query":   result.Query,
		})
	})

	app.Router.Run()

}
