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
	"io"
	"time"

	"github.com/enolgor/pdfsigner/cli/pdfsigner/actions/flags"
	"github.com/enolgor/pdfsigner/signer"
	"github.com/enolgor/pdfsigner/signer/config"
	"github.com/urfave/cli/v3"
)

var SignCommand *cli.Command = &cli.Command{
	Name:      "sign",
	Usage:     "sign a pdf (default subcommand if none is provided)",
	Category:  "signature",
	Aliases:   []string{"s"},
	Arguments: []cli.Argument{pdfArgument},
	Flags: []cli.Flag{
		flags.CertFlag,
		flags.PassphraseFlag,
		flags.SignedOutputFlag,
		flags.ForceWriteFlag,
		flags.DatetimeFlag,
		flags.LocationFlag,
		flags.SignatureNameFlag,
		flags.SignatureReasonFlag,
		flags.SignatureLocationFlag,
		flags.SignatureContactFlag,
		flags.PageFlag,
		flags.AddPageFlag,
		flags.PageSizeFlag,

		flags.VisibleFlag,
		flags.WidthFlag,
		flags.HeightFlag,
		flags.XposFlag,
		flags.YposFlag,
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
		flags.BackgroundColorFlag,
		flags.BorderSizeFlag,
		flags.BorderColorFlag,
		flags.LogoFlag,
		flags.LogoGrayscaleFlag,
		flags.LogoOpacityFlag,
		flags.LogoAlignmentFlag,
		flags.NoEmptyLineAfterTitleFlag,
		flags.TitleAlignmentFlag,
		flags.LineAlignmentFlag,
		flags.KeyAlignmentFlag,
		flags.ValueAlignmentFlag,
		flags.LoadFontFlag,
		flags.TitleFontFlag,
		flags.KeyFontFlag,
		flags.ValueFontFlag,
		flags.TitleColorFlag,
		flags.KeyColorFlag,
		flags.ValueColorFlag,
	},
	DisableSliceFlagSeparator: true,
	Action: func(ctx context.Context, cmd *cli.Command) (err error) {
		var cert *signer.UnlockedCertificate
		var pdf *bytes.Reader
		var signed io.WriteCloser
		var conf *config.SignatureConfiguration
		var metadata *signer.SignatureMetadata
		var date time.Time

		if cert, err = readCertificate(flags.Cert(cmd), flags.Passphrase(cmd)); err != nil {
			return
		}
		if pdf, err = readPdf(cmd); err != nil {
			return
		}
		if signed, err = flags.SignedOutput(cmd); err != nil {
			return
		}
		defer signed.Close()
		metadata = getMetadata(cmd)
		date = flags.Datetime(cmd)
		if !flags.Visible(cmd) {
			err = signer.Sign(cert, pdf, signed, date, metadata)
		} else {
			if conf, err = getConfiguration(cmd, pdf); err != nil {
				return
			}
			err = signer.SignVisual(cert, pdf, signed, date, metadata, conf)
		}
		return
	},
}
