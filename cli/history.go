package cli

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/models"
	"github.com/OrbitalJin/michi/internal/service"
	fzf "github.com/ktr0731/go-fuzzyfinder"
	v2 "github.com/urfave/cli/v2"
)

func history(service service.HistoryServiceIface) *v2.Command {
	return &v2.Command{
		Name:  "history",
		Usage: "to manage history",
		Subcommands: []*v2.Command{
			list(service),
		},
	}
}

func list(service service.HistoryServiceIface) *v2.Command {
	allFlag := &v2.BoolFlag{
		Name:  "all",
		Usage: "list all history",
	}

	limitFlag := &v2.IntFlag{
		Name:  "limit",
		Usage: "limit of history",
	}

	return &v2.Command{
		Name:  "list",
		Usage: "list history",
		Flags: []v2.Flag{
			allFlag,
			limitFlag,
		},
		Action: func(ctx *v2.Context) error {
			var history []models.SearchHistoryEvent
			var err error = nil

			all := ctx.Bool("all")
			limit := ctx.Int("limit")

			if all || limit < 1 {
				history, err = service.GetAllHistory()
			} else {
				history, err = service.GetRecentHistory(limit)
			}

			if err != nil {
				return err
			}

			selected := fzfHistory(history)
			fmt.Println(selected)

			return nil
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

	return &history[index[0]]

}
