package flags

import (
	"errors"
	"strconv"
	"strings"

	"github.com/enolgor/pdfsigner/signer/config"
	"github.com/rotisserie/eris"
	"github.com/urfave/cli/v3"
)

// pdf flags

var PageFlag *cli.IntFlag = &cli.IntFlag{
	Name:     "page",
	Aliases:  []string{"p"},
	Value:    1,
	Usage:    "page of the pdf file (1-indexed)",
	Sources:  cli.EnvVars("PAGE"),
	Required: false,
	Category: "pdf",
}

func Page(cmd *cli.Command) int {
	return cmd.Int("page")
}

var AddPageFlag = &cli.BoolFlag{
	Name:     "add-page",
	Aliases:  []string{"a"},
	Value:    false,
	Usage:    "add a page to the end of the pdf file (ignores page flag)",
	Sources:  cli.EnvVars("ADD_PAGE"),
	Category: "pdf",
}

func AddPage(cmd *cli.Command) bool {
	return cmd.Bool("add-page")
}

var PageSizeFlag = &cli.StringFlag{
	Name:     "page-size",
	Aliases:  []string{"ps"},
	Value:    "A4",
	Usage:    "page size if add-page is set",
	Sources:  cli.EnvVars("PAGE_SIZE"),
	Required: false,
	Category: "pdf",
}

func PageSize(cmd *cli.Command) (*config.Dim, error) {
	pageSize := cmd.String("page-size")
	parts := strings.Split(pageSize, ",")
	var width, height float64
	var errs, err error
	if len(parts) >= 2 {
		width, err = strconv.ParseFloat(parts[0], 64)
		errs = errors.Join(errs, err)
		height, err = strconv.ParseFloat(parts[1], 64)
		errs = errors.Join(errs, err)
	}
	if len(parts) < 2 || errs != nil {
		size, ok := config.PaperSize[pageSize]
		if !ok {
			return nil, eris.Errorf("invalid page size %q", pageSize)
		}
		width = size.Width
		height = size.Height
	}
	return &config.Dim{Width: width, Height: height}, nil
}
