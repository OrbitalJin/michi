package bangs

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/models"
	"github.com/OrbitalJin/michi/internal/service"
	fuzzy "github.com/ktr0731/go-fuzzyfinder"
	"github.com/urfave/cli/v2"
)

func Root(service service.SPServiceIface) *cli.Command {
	return &cli.Command{
		Name:  "bangs",
		Usage: "to manage bangs",
		Subcommands: []*cli.Command{
			list(service),
			delete(service),
		},
	}
}

func fzf(bangs []models.SearchProvider, prompt string) *models.SearchProvider {
	index, err := fuzzy.FindMulti(
		bangs,
		func(i int) string {
			return bangs[i].Domain
		},
		fuzzy.WithHeader("Bangs - "+prompt),
		fuzzy.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			provider := bangs[i]

			return fmt.Sprintf(
				"Site Name: %s \n"+
					"Tag: %s \n"+
					"Category: %s \n"+
					"Subcategory: %s \n"+
					"Domain: %s \n"+
					"Rank: %d \n",
				provider.SiteName,
				provider.Tag,
				provider.Category,
				provider.Subcategory,
				provider.Domain,
				provider.Rank,
			)
		}),
	)

	if err != nil {
		return nil
	}

	if len(index) > 0 {
		return &bangs[index[0]]
	}

	return nil
}
