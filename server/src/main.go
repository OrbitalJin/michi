package main

import (
	"log"
	"net/http"

	"github.com/OrbitalJin/pow/internal/app"
	"github.com/OrbitalJin/pow/internal/parser"
	"github.com/OrbitalJin/pow/internal/store"
	"github.com/gin-contrib/cors" // Import the cors package
	"github.com/gin-gonic/gin"
)

func main() {
	parserConfig := parser.NewConfig(`!(\b\w+\b)`, `!\b\w+\b`)
	storeConfig := store.NewConfig("./database.db", "g")
	appConfig := app.NewConfig(parserConfig, storeConfig)
	app := app.New(appConfig)

	app.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value of the Access-Control-Max-Age header in seconds.
	}))

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
			return
		}

		redirect, err := app.Service.ResolveWithFallback(result.Query, provider)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, "soup")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"redirect": redirect,
		})
	})

	app.Start()
}
