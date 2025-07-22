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
	"github.com/enolgor/pdfsigner/signer"
	"github.com/enolgor/pdfsigner/signer/config"
	"github.com/urfave/cli/v3"
)

var SignatureDimCommand *cli.Command = &cli.Command{
	Name:     "signature-dim",
	Usage:    "get visible signature dimensions in pt",
	Category: "signature",
	Aliases:  []string{"sd"},
	Flags: []cli.Flag{
		flags.CertFlag,
		flags.PassphraseFlag,
		flags.DatetimeFlag,
		flags.LocationFlag,

		flags.WidthFlag,
		flags.HeightFlag,
		flags.RotateFlag,
		flags.DpiFlag,
		flags.TitleFlag,
		flags.NoTitleFlag,
		flags.DatetimeFormatFlag,
		flags.NoSubjectFlag,
		flags.NoIssuerFlag,
		flags.NoDateFlag,
		flags.SubjectKeyFlag,
		flags.IssuerKeyFlag,
		flags.DateKeyFlag,
		flags.ExtraLinesFlag,
		flags.LoadFontFlag,
		flags.TitleFontFlag,
		flags.KeyFontFlag,
		flags.ValueFontFlag,
		flags.NoEmptyLineAfterTitleFlag,
	},
	DisableSliceFlagSeparator: true,
	Action: func(ctx context.Context, cmd *cli.Command) (err error) {
		var cert *signer.UnlockedCertificate
		var conf *config.SignatureConfiguration
		var widthPt, heightPt float64
		if cert, err = readCertificate(flags.Cert(cmd), flags.Passphrase(cmd)); err != nil {
			return
		}
		if conf, err = getConfiguration(cmd, nil); err != nil {
			return
		}
		if widthPt, heightPt, err = signer.CalculateSignatureDim(flags.Datetime(cmd), cert, conf); err != nil {
			return
		}
		fmt.Println(widthPt, heightPt)
		return
	},
}
