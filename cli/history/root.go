package history

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/models"
	"github.com/OrbitalJin/michi/internal/service"
	fzf "github.com/ktr0731/go-fuzzyfinder"
	v2 "github.com/urfave/cli/v2"
)

func Root(service service.HistoryServiceIface) *v2.Command {
	return &v2.Command{
		Name:  "history",
		Usage: "to manage history",
		Subcommands: []*v2.Command{
			list(service),
			delete(service),
		},
	}
}

func fzfHistory(history []models.SearchHistoryEvent) *models.SearchHistoryEvent {
	index, err := fzf.FindMulti(
		history,
		func(i int) string {
			return history[i].Query

		},
		fzf.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Query: %s \nProvider: (%s) \nTimeStamp: %s",
				history[i].Query,
				history[i].ProviderTag,
				history[i].Timestamp,
			)
		}))

	if err != nil {
		return nil
	}

	if len(index) > 0 {
		return &history[index[0]]
	}

	return nil
}
