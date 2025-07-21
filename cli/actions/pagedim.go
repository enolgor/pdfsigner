package actions

import (
	"bytes"
	"context"
	"fmt"

	"github.com/enolgor/pdfsigner/cli/actions/flags"
	"github.com/enolgor/pdfsigner/signer"
	"github.com/urfave/cli/v3"
)

var PageDimCommand *cli.Command = &cli.Command{
	Name:      "page-dim",
	Usage:     "get page dimensions in pt",
	Category:  "pdf",
	Aliases:   []string{"pd"},
	Arguments: []cli.Argument{pdfArgument},
	Flags: []cli.Flag{
		flags.PageFlag,
	},
	Action: func(ctx context.Context, cmd *cli.Command) (err error) {
		var pdf *bytes.Reader
		var width, height float64
		if pdf, err = readPdf(cmd); err != nil {
			return
		}
		if width, height, err = signer.GetPageDimensionsPt(pdf, flags.Page(cmd)-1); err != nil {
			return
		}
		fmt.Println(width, height)
		return
	},
}
