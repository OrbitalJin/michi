package shortcuts

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	"github.com/atotto/clipboard"
	v2 "github.com/urfave/cli/v2"
)

func list(service service.ShortcutServiceIface) *v2.Command {
	return &v2.Command{
		Name:  "list",
		Usage: "to list shortcuts",
		Action: func(c *v2.Context) error {
			shortcuts, err := service.GetAll()

			if err != nil {
				return err
			}

			selected := fzf(shortcuts)

			if selected == nil {
				fmt.Println("No shortcut selected")
				return nil
			}

			clipboard.WriteAll(selected.URL)

			return nil
		},
	}
}
