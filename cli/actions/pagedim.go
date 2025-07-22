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
