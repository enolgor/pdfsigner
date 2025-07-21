package main

import (
	"context"
	"fmt"
	"os"

	"github.com/enolgor/pdfsigner/cli/actions"
	"github.com/rotisserie/eris"
	"github.com/urfave/cli/v3"
)

var Version = "dev"

func main() {
	cli.HelpFlag = &cli.BoolFlag{
		Name:        "help",
		Usage:       "show help",
		HideDefault: true,
		Local:       true,
	}
	cmd := &cli.Command{
		Name:  "pdfsigner",
		Usage: "A tool to sign pdfs",
		Commands: []*cli.Command{
			actions.PageCountCommand,
			actions.PageDimCommand,
			actions.SignatureDimCommand,
			actions.SignCommand,
			actions.ListFontsCommand,
		},
		DefaultCommand: actions.SignCommand.Name,
		Version:        Version,
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "fatal error:%s\n", eris.ToString(err, true))
		os.Exit(1)
	}
}
