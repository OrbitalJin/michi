package handler

import (
	"log"
	"net/http"

	"github.com/OrbitalJin/qmuxr/internal/parser"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Root(ctx *gin.Context) {
	query := ctx.Query(h.QueryParam)

	if query == "" {
		respondWithError(
			ctx,
			http.StatusBadRequest,
			"Root: Query parameter '%s' is required. No query provided.",
			"Query parameter 'q' is required.",
			nil,
			h.QueryParam,
		)
		return
	}

	action := h.queryParser.ParseAction(query)

	log.Println(action.Type)

	switch action.Type {

	case parser.BANG:
		h.handleBang(ctx, action)

	case parser.SHORTCUT:
		h.handleShortcut(ctx, action)

	case parser.SESSION:
		h.handleSession(ctx, action)

	case parser.DEFAULT:
		h.handleDefaultSearch(ctx, action)

	default:
		respondWithError(
			ctx,
			http.StatusBadRequest,
			"Root: Unhandled action type '%v' for query: '%s'. This indicates a parser issue or malformed query.",
			"Couldn't understand your query. Please check the format.",
			nil,
			action.Type, query,
		)
	}
}
