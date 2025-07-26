package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/OrbitalJin/qmuxr/internal/models"
	"github.com/OrbitalJin/qmuxr/internal/service"
	"github.com/gin-gonic/gin"
)

var QueryParam = "q"

func Search(ctx *gin.Context, psvc *service.ProviderService, hsvc *service.HistoryService) {
	query := ctx.Query(QueryParam)

	if query == "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Query parameter 'q' is required."})
		return
	}

	result, provider, err := psvc.CollectAndRank(query)

	if err != nil {
		log.Printf("failed to collect and rank for query '%s': %v", query, err)
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to process search query. Please try again later."},
		)
		return
	}

	provider, redirect, err := psvc.ResolveWithFallback(result.Query, provider)

	if err != nil {
		log.Printf(
			"failed to resolve redirect for query '%s' with provider '%s': %v",
			result.Query,
			provider.Tag,
			err,
		)
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Could not determine search destination. Please try again later."},
		)
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

		go logSearchHistoryAsync(hsvc, entry)
	}
}

func logSearchHistoryAsync(hsvc *service.HistoryService, entry *models.SearchHistoryEvent) {
	if err := hsvc.Insert(entry); err != nil {
		log.Printf(
			"failed to insert search history entry for query '%s': %v",
			entry.Query,
			fmt.Errorf("insertion error: %w", err),
		)
	}
}
