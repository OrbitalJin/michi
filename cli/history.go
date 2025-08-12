package cli

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/models"
	"github.com/OrbitalJin/michi/internal/service"
	"github.com/atotto/clipboard"
	fzf "github.com/ktr0731/go-fuzzyfinder"
	v2 "github.com/urfave/cli/v2"
)

func history(service service.HistoryServiceIface) *v2.Command {
	return &v2.Command{
		Name:  "history",
		Usage: "to manage history",
		Subcommands: []*v2.Command{
			list(service),
			delete(service),
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
			err = clipboard.WriteAll(selected.Query)

			if err != nil {
				return err
			}

			fmt.Printf("Selection copied to clipboard: %s\n", selected.Query)

			return nil
		},
	}
}

func delete(service service.HistoryServiceIface) *v2.Command {
	lastFlag := &v2.IntFlag{
		Name:  "last",
		Usage: "purge the last (n) entries",
	}

	return &v2.Command{
		Name:  "delete",
		Usage: "delete an entry",
		Flags: []v2.Flag{
			lastFlag,
		},
		Action: func(ctx *v2.Context) error {
			last := ctx.Int("last")

			if last > 0 {
				history, err := service.GetRecentHistory(last)

				if err != nil {
					return err
				}

				for _, entry := range history {
					service.DeleteEntry(entry.ID)
				}

				fmt.Printf("Last (%d) have been successfully purged.\n", last)
				return nil
			}

			history, err := service.GetAllHistory()

			if err != nil {
				return err
			}

			selected := fzfHistory(history)

			if selected == nil {
				fmt.Println("No entry selected.")
				return nil
			}

			err = service.DeleteEntry(selected.ID)

			if err != nil {
				return err
			}

			fmt.Println("Deleted successfully.")
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

	if len(index) > 0 {
		return &history[index[0]]
	}

	return nil
}
