package actions

import (
	"context"
	"fmt"

	"github.com/enolgor/pdfsigner/cli/actions/flags"
	"github.com/enolgor/pdfsigner/signer/fonts"
	"github.com/urfave/cli/v3"
)

var ListFontsCommand *cli.Command = &cli.Command{
	Name:     "list-fonts",
	Usage:    "list available fonts",
	Category: "signature",
	Flags: []cli.Flag{
		flags.LoadFontFlag,
	},
	DisableSliceFlagSeparator: true,
	Action: func(ctx context.Context, cmd *cli.Command) (err error) {
		if err := flags.LoadFonts(cmd); err != nil {
			return err
		}
		for _, fontName := range fonts.ListLoadedFonts() {
			fmt.Println(fontName)
		}
		return
	},
}
