package actions

import (
	"bytes"
	"context"
	"fmt"

	"github.com/enolgor/pdfsigner/signer"
	"github.com/urfave/cli/v3"
)

var PageCountCommand *cli.Command = &cli.Command{
	Name:      "page-count",
	Usage:     "count pages in a pdf",
	Category:  "pdf",
	Aliases:   []string{"pc"},
	Arguments: []cli.Argument{pdfArgument},
	Action: func(ctx context.Context, cmd *cli.Command) (err error) {
		var pdf *bytes.Reader
		var count int
		if pdf, err = readPdf(cmd); err != nil {
			return
		}
		if count, err = signer.GetPageCount(pdf); err != nil {
			return
		}
		fmt.Println(count)
		return
	},
}
