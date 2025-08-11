package handler

import (
	"net/http"

	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/gin-gonic/gin"
)

func (h *Handler) handleSession(ctx *gin.Context, action *parser.QueryAction) {
	result := action.Result

	if result == nil || len(result.Matches) == 0 {
		respondWithError(
			ctx,
			http.StatusInternalServerError,
			"handleSession: Received malformed parser result for query '%s'. Result: %+v.",
			"Failed to process your session request. Please try again.",
			nil,
			ctx.Query(h.QueryParam), result,
		)
		return
	}

	alias := result.Matches[0]

	session, err := h.sessionService.GetFromAlias(alias)

	if err != nil {
		respondWithError(
			ctx,
			http.StatusInternalServerError,
			"handleSession: Error retrieving session for alias '%s': %v",
			"Failed to retrieve session. Please try again later.",
			err,
			alias,
		)
		return
	}

	if session == nil {
		h.handleDefaultSearch(ctx, action)
		return
	}

	urls := session.URLs

	if len(urls) == 0 {
		respondWithError(
			ctx,
			http.StatusInternalServerError,
			"handleSession: Session for alias '%s' has no URLs",
			"Failed to retrieve session. Please try again later.",
			nil,
			alias,
		)
		return
	}

	ctx.HTML(
		http.StatusOK,
		"session_open.html",
		gin.H{
			"URLs": urls,
		},
	)
}
