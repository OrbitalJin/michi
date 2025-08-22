package handler

import (
	"net/http"

	"github.com/OrbitalJin/michi/internal/models"
	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/public"
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
	queryParser parser.QueryParserIface
	services    *service.Services
	QueryParam  string
}

func NewHandler(
	qp parser.QueryParserIface,
	services *service.Services,
	queryParam string,

) *Handler {

	return &Handler{
		queryParser: qp,
		services:    services,
		QueryParam:  queryParam,
	}
}

func (h *Handler) Favicon(ctx *gin.Context) {
		data, err := public.Content.ReadFile("assets/favicon.svg")
		if err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Data(http.StatusOK, "image/svg+xml", data)
}

func (h *Handler) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func (h *Handler) Error(ctx *gin.Context) {
	message := ctx.Query("message")

	if message != "" {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": message})
	}

	ctx.HTML(http.StatusInternalServerError, "error.html", nil)
}

func (h *Handler) SessionOpened(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "session_success.html", nil)
}

func (h *Handler) completeSearchRequest(
	ctx *gin.Context,
	redirectURL string,
	result *parser.Result,
	provider *models.SearchProvider,
) {

	ctx.Redirect(http.StatusFound, redirectURL)

	if h.services.GetProvidersService().GetCfg().ShouldKeepTrack() {
		go h.logSearchHistoryAsync(result, provider)
	}
}
