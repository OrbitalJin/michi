package server

import (
	"log"
	"net/http"
	"time"

	"github.com/OrbitalJin/qmuxr/internal/models"
	"github.com/OrbitalJin/qmuxr/internal/service"
	"github.com/gin-gonic/gin"
)

func Search(ctx *gin.Context, psvc *service.ProviderService, hsvc *service.HistoryService) {
	query := ctx.Query("q")

	result, provider, err := psvc.CollectAndRank(query)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, "error")
		return
	}

	provider, redirect, err := psvc.ResolveWithFallback(result.Query, provider)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	ctx.Redirect(http.StatusFound, *redirect)

	if psvc.GetCfg().ShouldKeepTrack() {
		entry := &models.SearchHistoryEvent{
			Query:       result.Query,
			ProviderID:  provider.ID,
			ProviderTag: provider.Tag,
			Timestamp:   time.Now(),
		}

		go func(entry *models.SearchHistoryEvent) {
			if err := hsvc.Insert(entry); err != nil {
				log.Println("failed to insert search history entry: %w", err)
			}

		}(entry)
	}
}
