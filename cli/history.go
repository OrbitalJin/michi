package cli

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/urfave/cli/v2"
)

func history(services *service.Services) *cli.Command {
	return &cli.Command{
		Name: "history",
		Action: func(ctx *cli.Context) error {
			history, err := services.GetHistoryService().GetRecentHistory(30)

			if err != nil {
				return err
			}

			idx, err := fuzzyfinder.FindMulti(
				history,
				func(i int) string {
					return history[i].Query
				},
				fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
					if i == -1 {
						return ""
					}
					return fmt.Sprintf("Track: %s (%s)\nAlbum: %s",
						history[i].Query,
						history[i].ProviderTag,
						history[i].Timestamp,
					)
				}))
			if err != nil {
				return err
			}

			fmt.Printf("selected: %v\n", idx)
			return nil
		},
	}
}
