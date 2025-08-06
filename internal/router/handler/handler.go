package handler

import (
	"log"
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
	queryParser     parser.QueryParserIface
	providerService service.SPServiceIface
	historyService  service.HistoryServiceIface
	shortcutService service.ShortcutServiceIface
	sessionService  service.SessionServiceIface
	QueryParam      string
}

func NewHandler(
	qp parser.QueryParserIface,
	psvc service.SPServiceIface,
	hsvc service.HistoryServiceIface,
	scsvc service.ShortcutServiceIface,
	seshSvc service.SessionServiceIface,
	queryParam string,

) *Handler {

	return &Handler{
		queryParser:     qp,
		providerService: psvc,
		historyService:  hsvc,
		shortcutService: scsvc,
		sessionService:  seshSvc,
		QueryParam:      queryParam,
	}
}

func (h *Handler) Favicon(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNoContent)
}

func (h *Handler) Error(ctx *gin.Context, message string) {
	log.Println(message)
	ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to load session: " + message})
}

func (h *Handler) SessionOpened(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "session_success.html",
		gin.H{
			"DefaultEngine": "https://www.google.com/",
		},
	)
}

func (h *Handler) completeSearchRequest(
	ctx *gin.Context,
	redirectURL string,
	result *parser.Result,
	provider *models.SearchProvider,
) {

	ctx.Redirect(http.StatusFound, redirectURL)

	if h.providerService.GetCfg().ShouldKeepTrack() {
		go h.logSearchHistoryAsync(result, provider)
	}
}
