package shortcuts

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	v2 "github.com/urfave/cli/v2"
)

func delete(service service.ShortcutServiceIface) *v2.Command {
	return &v2.Command{
		Name: "delete",
		Flags: []v2.Flag{
			&v2.StringFlag{
				Name:     "alias",
				Usage:    "alias for the shortcut",
				Required: true,
			},
		},
		Usage: "delete a shortcut",
		Action: func(c *v2.Context) error {
			alias := c.String("alias")

			err := service.DeleteFromAlias(alias)

			if err != nil {
				return err
			}

			fmt.Println("Shortcut deleted.")
			return nil
		},
	}
}
