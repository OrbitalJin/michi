package main

import (
	"log"
	"net/http"

	"github.com/OrbitalJin/pow/internal/app"
	"github.com/OrbitalJin/pow/internal/parser"
	"github.com/OrbitalJin/pow/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	parserConfig := parser.NewConfig(`!(\b\w+\b)`, `!\b\w+\b`)
	storeConfig := store.NewConfig("./database.db")
	appConfig := app.NewConfig(parserConfig, storeConfig)
	app := app.New(appConfig)

	app.Router.GET("/provider", func(ctx *gin.Context) {
		tag := ctx.Query("tag")
		provider, err := app.Service.GetByTag(tag)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, "error")
		}

		ctx.JSON(http.StatusOK, provider)
	})

	app.Router.POST("/collect", func(ctx *gin.Context) {
		value := ctx.Query("value")
		result, provider, err := app.Service.CollectAndRank(value)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, "error")
		}

		redirect, err := app.Service.Resolve(result.Query, provider)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, "error")
		}

		ctx.JSON(http.StatusOK, gin.H{
			"redirect": redirect,
		})
	})

	app.Start()
}
