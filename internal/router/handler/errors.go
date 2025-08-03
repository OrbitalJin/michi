package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func respondWithError(
	ctx *gin.Context,
	statusCode int,
	logMsgFormat string,
	userMsg string,
	err error,
	args ...any,
) {

	logArgs := make([]any, 0, len(args)+1)
	if err != nil {
		logArgs = append(logArgs, err)
	}
	logArgs = append(logArgs, args...)

	log.Printf(logMsgFormat, logArgs...)

	ctx.HTML(statusCode, "error.html", gin.H{"error": userMsg})
}
