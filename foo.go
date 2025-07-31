package main

import (
	"embed"
	"html/template" // We'll use html/template directly
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var embeddedTemplates embed.FS

func main() {
	r := gin.Default()

	templatesFS, err := fs.Sub(embeddedTemplates, "templates")

	if err != nil {
		panic(err)
	}
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}

	tmpl := template.New("")
	tmpl = tmpl.Funcs(funcMap)
	tmpl = template.Must(tmpl.ParseFS(templatesFS, "*.html"))

	r.SetHTMLTemplate(tmpl)

	// This endpoint will redirect to an intermediate page
	r.GET("/multi-redirect", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/intermediate-redirect-page")
	})

	// This endpoint serves the intermediate HTML page with dynamic URLs
	r.GET("/foo", func(c *gin.Context) {
		urls := []string{
			"https://www.bing.com",
			"https://go.dev",
			"https://en.wikipedia.org/wiki/Go_(programming_language)",
		}

		c.HTML(http.StatusOK, "session.html", gin.H{
			"URLs": urls,
		})
	})

	r.Run(":8080")
}
