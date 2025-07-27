package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/OrbitalJin/qmuxr/internal/models"
	"github.com/OrbitalJin/qmuxr/internal/parser"
	"github.com/OrbitalJin/qmuxr/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	QueryParser     *parser.QueryParser
	ProviderService *service.SearchProviderService
	HistoryService  *service.HistoryService
	QueryParam      string
}

func NewHandler(
	qp *parser.QueryParser,
	psvc *service.SearchProviderService,
	hsvc *service.HistoryService,
	queryParam string,

) *Handler {

	return &Handler{
		QueryParser:     qp,
		ProviderService: psvc,
		HistoryService:  hsvc,
		QueryParam:      queryParam,
	}
}

func (h *Handler) Root(ctx *gin.Context) {
	query := ctx.Query(h.QueryParam)

	if query == "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Query parameter 'q' is required."})
		return
	}

	action := h.QueryParser.ParseAction(query)

	switch action.Type {
	case parser.BANG:
		{
			h.handleBang(ctx, action.Result)
		}
	case parser.SHORTCUT:
		{
			log.Println("shortcut!")
		}
	}
}

func (h *Handler) handleBang(ctx *gin.Context, result *parser.Result) {
	best := h.ProviderService.Rank(result)

	provider, redirect, err := h.ProviderService.ResolveWithFallback(result.Query, best)

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

	if h.ProviderService.GetCfg().ShouldKeepTrack() {
		entry := &models.SearchHistoryEvent{
			Query:       result.Query,
			ProviderID:  provider.ID,
			ProviderTag: provider.Tag,
			Timestamp:   time.Now(),
		}

		go h.logSearchHistoryAsync(entry)
	}
}

func (h *Handler) logSearchHistoryAsync(entry *models.SearchHistoryEvent) {
	if err := h.HistoryService.Insert(entry); err != nil {
		log.Printf(
			"failed to insert search history entry for query '%s': %v",
			entry.Query,
			fmt.Errorf("insertion error: %w", err),
		)
	}
}
