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
	"context"
	"fmt"

	"github.com/enolgor/pdfsigner/cli/pdfsigner/actions/flags"
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
