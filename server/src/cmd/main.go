package main

import (
	"net/http"

	"github.com/OrbitalJin/pow/internal/parser"
	"github.com/gin-gonic/gin"
)

func main() {
	parserConfig := parser.NewConfig(`!(\b\w+\b)`, `!\b\w+\b`)
	parser, err := parser.NewParser(parserConfig)

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.POST("/collect", func(c *gin.Context) {

		req := c.Query("value")

		result, err := parser.Collect(req)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "error")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"matches": result.Matches,
			"query":   result.Query,
		})
	})

	router.Run()

}
