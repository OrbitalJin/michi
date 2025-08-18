package bangs

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	"github.com/atotto/clipboard"
	v2 "github.com/urfave/cli/v2"
)

func list(service service.SPServiceIface) *v2.Command {
	return &v2.Command{
		Name:  "list",
		Usage: "list bangs",
		Action: func(ctx *v2.Context) error {
			bangs, err := service.GetAll()

			if err != nil {
				return err
			}

			selected := fzf(bangs, "Copy")

			if selected == nil {
				fmt.Println("no bang selected")
				return nil
			}

			clipboard.WriteAll(selected.Tag)
			fmt.Printf("Bang `%s` copied.\n", selected.Tag)
			return nil
		},
	}
}
