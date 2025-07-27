package handler

import (
	"net/http"

	"github.com/OrbitalJin/qmuxr/internal/parser"
	"github.com/gin-gonic/gin"
)

func (h *Handler) handleBang(ctx *gin.Context, action *parser.QueryAction) {
	result := action.Result

	if result == nil {
		respondWithError(
			ctx,
			http.StatusInternalServerError,
			"handleBang: Parser result is nil for query '%s'.",
			"An unexpected error occurred while processing your bang query.",
			nil,
			action.RawQuery,
		)
		return
	}

	best := h.ProviderService.Rank(result)

	provider, redirect, err := h.ProviderService.ResolveAndFallback(
		result.Query,
		best,
	)

	if err != nil || redirect == nil {
		providerTag := "N/A"
		if provider != nil {
			providerTag = provider.Tag
		}
		respondWithError(
			ctx,
			http.StatusInternalServerError,
			"handleBang: Failed to resolve redirect for query '%s' with provider '%s': %v",
			"Could not determine search destination. Please try again later.",
			err,
			result.Query, providerTag,
		)
		return
	}

	h.completeSearchRequest(ctx, *redirect, result, provider)
}
