package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/OrbitalJin/qmuxr/internal/models"
	"github.com/OrbitalJin/qmuxr/internal/parser"
)

func (h *Handler) logSearchHistoryAsync(result *parser.Result, provider *models.SearchProvider) {
	if result == nil || provider == nil {
		log.Printf(
			"logSearchHistoryAsync: Skipping history log due to missing result or provider. Result: %+v, Provider: %+v",
			result,
			provider,
		)
		return
	}

	entry := &models.SearchHistoryEvent{
		Query:       result.Query,
		ProviderID:  provider.ID,
		ProviderTag: provider.Tag,
		Timestamp:   time.Now(),
	}

	if err := h.historyService.Insert(entry); err != nil {
		log.Printf(
			"failed to insert search history entry for query '%s': %v",
			entry.Query,
			fmt.Errorf("insertion error: %w", err),
		)
	}
}
