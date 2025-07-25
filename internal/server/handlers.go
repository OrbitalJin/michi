package server

import (
	"log"
	"net/http"

	"github.com/OrbitalJin/qmuxr/internal/service"
	"github.com/gin-gonic/gin"
)

func Search(ctx *gin.Context, service *service.ProviderService) {
	query := ctx.Query("q")

	result, provider, err := service.CollectAndRank(query)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, "error")
		return
	}

	redirect, err := service.ResolveWithFallback(result.Query, provider)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, "soup")
		return
	}

	ctx.Redirect(http.StatusFound, *redirect)
}
