package handler

import (
	"net/http"

	"github.com/OrbitalJin/qmuxr/internal/models"
	"github.com/OrbitalJin/qmuxr/internal/parser"
	"github.com/OrbitalJin/qmuxr/internal/service"
	"github.com/gin-gonic/gin"
)

type HandlerIface interface {
	Root(ctx *gin.Context)
	completeSearchRequest(ctx *gin.Context, redirectURL string, result *parser.Result, provider *models.SearchProvider)
	handleBang(ctx *gin.Context, action *parser.QueryAction)
	handleShortcut(ctx *gin.Context, action *parser.QueryAction)
	handleSession(ctx *gin.Context, action *parser.QueryAction)
	handleDefaultSearch(ctx *gin.Context, action *parser.QueryAction)
}

type Handler struct {
	QueryParser     parser.QueryParserIface
	ProviderService service.SPServiceIface
	HistoryService  service.HistoryServiceIface
	ShortcutService service.ShortcutServiceIface
	QueryParam      string
}

func NewHandler(
	qp parser.QueryParserIface,
	psvc service.SPServiceIface,
	hsvc service.HistoryServiceIface,
	scsvc service.ShortcutServiceIface,
	queryParam string,

) *Handler {

	return &Handler{
		QueryParser:     qp,
		ProviderService: psvc,
		HistoryService:  hsvc,
		ShortcutService: scsvc,
		QueryParam:      queryParam,
	}
}

func (h *Handler) completeSearchRequest(
	ctx *gin.Context,
	redirectURL string,
	result *parser.Result,
	provider *models.SearchProvider,
) {

	ctx.Redirect(http.StatusFound, redirectURL)

	if h.ProviderService.GetCfg().ShouldKeepTrack() {
		go h.logSearchHistoryAsync(result, provider)
	}
}
