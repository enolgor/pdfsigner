// MIT License
//
// Copyright (c) 2025 @enolgor
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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

const pdfCategory = "pdf"

var PageFlag *cli.IntFlag = &cli.IntFlag{
	Name:     "page",
	Aliases:  []string{"p"},
	Value:    1,
	Usage:    "page of the pdf file (1-indexed)",
	Sources:  cli.EnvVars("PAGE"),
	Required: false,
	Category: pdfCategory,
}

func Page(cmd *cli.Command) int {
	return cmd.Int(PageFlag.Name)
}

var AddPageFlag = &cli.BoolFlag{
	Name:     "add-page",
	Aliases:  []string{"a"},
	Value:    false,
	Usage:    "add a page to the end of the pdf file (ignores page flag)",
	Sources:  cli.EnvVars("ADD_PAGE"),
	Category: pdfCategory,
}

func AddPage(cmd *cli.Command) bool {
	return cmd.Bool(AddPageFlag.Name)
}

var PageSizeFlag = &cli.StringFlag{
	Name:     "page-size",
	Aliases:  []string{"ps"},
	Value:    "A4",
	Usage:    "page size if add-page is set",
	Sources:  cli.EnvVars("PAGE_SIZE"),
	Required: false,
	Category: pdfCategory,
}

func PageSize(cmd *cli.Command) (*config.Dim, error) {
	pageSize := cmd.String(PageSizeFlag.Name)
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
