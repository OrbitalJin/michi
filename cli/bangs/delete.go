package bangs

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/service"
	v2 "github.com/urfave/cli/v2"
)

func delete(service service.SPServiceIface) *v2.Command {
	return &v2.Command{
		Name:  "delete",
		Usage: "delete a bang",
		Action: func(ctx *v2.Context) error {
			bangs, err := service.GetAll()

			if err != nil {
				return err
			}

			selected := fzf(bangs, "Delete")

			if selected == nil {
				fmt.Println("no bang selected")
				return nil
			}

			err = service.Delete(selected.ID)

			if err != nil {
				return err
			}

			fmt.Printf("Bang `%s` deleted.\n", selected.Tag)
			return nil

		},
	}
}
