package sessions

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/models"
	"github.com/OrbitalJin/michi/internal/service"
	v2 "github.com/urfave/cli/v2"
)

func create(service service.SessionServiceIface) *v2.Command {
	aliasFlag := &v2.StringFlag{
		Name:     "alias",
		Usage:    "alias for the session",
		Required: true,
	}

	urlFlag := &v2.StringSliceFlag{
		Name:     "url",
		Usage:    "urls for the session (can be specified multiple times)",
		Required: true,
	}

	return &v2.Command{
		Name:  "create",
		Usage: "create a new session",
		Flags: []v2.Flag{
			aliasFlag,
			urlFlag,
		},
		Action: func(c *v2.Context) error {
			alias := c.String("alias")
			urls := c.StringSlice("url")

			err := service.Insert(&models.Session{
				Alias: alias,
				URLs:  urls,
			})

			if err != nil {
				return err
			}

			fmt.Printf("Successfully created session `%s` with URLs: %v\n", alias, urls)
			return nil
		},
	}
}
