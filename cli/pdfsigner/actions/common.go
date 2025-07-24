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
	"image"
	"image/color"
	"os"

	"github.com/enolgor/pdfsigner/cli/pdfsigner/actions/flags"
	"github.com/enolgor/pdfsigner/signer"
	"github.com/enolgor/pdfsigner/signer/config"
	"github.com/rotisserie/eris"
	"github.com/urfave/cli/v3"
)

var pdfArgument *cli.StringArg = &cli.StringArg{
	Name:      "pdf-file",
	UsageText: "path to pdf file",
	Config: cli.StringConfig{
		TrimSpace: true,
	},
}

func readPdf(cmd *cli.Command) (*bytes.Reader, error) {
	path := cmd.StringArg("pdf-file")
	if path == "" {
		return nil, eris.New("path to pdf file must be provided")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to read file %s", path)
	}
	return bytes.NewReader(data), nil
}

func readCertificate(path, passphrase string) (*signer.UnlockedCertificate, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to read file %s", path)
	}
	return signer.UnlockCertificate(data, passphrase)
}

func getMetadata(cmd *cli.Command) *signer.SignatureMetadata {
	return &signer.SignatureMetadata{
		Name:     flags.SignatureName(cmd),
		Location: flags.SignatureLocation(cmd),
		Reason:   flags.SignatureReason(cmd),
		Contact:  flags.SignatureContact(cmd),
	}
}

func getOptions(cmd *cli.Command) ([]func(*signer.SignatureOptions), error) {
	var options []func(*signer.SignatureOptions)
	if flags.TsaURL(cmd) != "" {
		tsa := signer.TSA{
			URL:      flags.TsaURL(cmd),
			Username: flags.TsaUser(cmd),
			Password: flags.TsaPassword(cmd),
		}
		options = append(options, signer.WithTSA(tsa))
	}
	return options, nil
}

func getConfiguration(cmd *cli.Command, pdfReader *bytes.Reader) (*config.SignatureConfiguration, error) {
	var err error
	if err = flags.LoadFonts(cmd); err != nil {
		return nil, err
	}
	options := []config.SignatureOption{}
	options = append(options, config.WidthPt(flags.Width(cmd)))
	options = append(options, config.HeightPt(flags.Height(cmd)))
	if flags.HeightFlag.IsSet() && !flags.WidthFlag.IsSet() {
		options = append(options, config.WidthPt(0))
	}
	if options, err = applyPageConfiguration(options, cmd, pdfReader); err != nil {
		return nil, err
	}
	options = append(options, config.PosXPt(flags.Xpos(cmd)))
	options = append(options, config.PosYPt(flags.Ypos(cmd)))
	options = append(options, config.Rotate(flags.Rotate(cmd)))
	options = append(options, config.Dpi(flags.Dpi(cmd)))
	options = append(options, config.Title(flags.Title(cmd)))
	options = append(options, config.DateFormat(flags.DatetimeFormat(cmd)))
	options = append(options, config.IncludeIssuer(!flags.NoIssuer(cmd)))
	options = append(options, config.IncludeSubject(!flags.NoSubject(cmd)))
	options = append(options, config.IncludeDate(!flags.NoDate(cmd)))
	options = append(options, config.SubjectKey(flags.SubjectKey(cmd)))
	options = append(options, config.IssuerKey(flags.IssuerKey(cmd)))
	options = append(options, config.DateKey(flags.DateKey(cmd)))
	extra, err := flags.ExtraLines(cmd)
	if err != nil {
		return nil, err
	}
	for _, line := range extra {
		options = append(options, config.ExtraLine(line.Key, line.Value))
	}
	options = append(options, config.BorderSizePt(flags.BorderSize(cmd)))
	var logo image.Image
	if logo, err = flags.Logo(cmd); err != nil {
		return nil, err
	}
	options = append(options, config.Logo(logo))
	options = append(options, config.LogoGrayscale(flags.LogoGrayscale(cmd)))
	options = append(options, config.LogoOpacity(flags.LogoOpacity(cmd)))
	options = append(options, config.LogoAlignment(flags.LogoAlignment(cmd)))
	options = append(options, config.EmptyLineAfterTitle(!flags.NoEmptyLineAfterTitle(cmd)))
	options = append(options, config.TitleAlignment(flags.TitleAlignment(cmd)))
	options = append(options, config.LineAlignment(flags.LineAlignment(cmd)))
	options = append(options, config.KeyAlignment(flags.KeyAlignment(cmd)))
	options = append(options, config.ValueAlignment(flags.ValueAlignment(cmd)))
	options = append(options, config.TitleFont(flags.TitleFont(cmd)))
	options = append(options, config.KeyFont(flags.KeyFont(cmd)))
	options = append(options, config.ValueFont(flags.ValueFont(cmd)))

	if options, err = applyColors(options, cmd); err != nil {
		return nil, err
	}
	return config.New(options...), err
}

func applyPageConfiguration(options []config.SignatureOption, cmd *cli.Command, pdfReader *bytes.Reader) ([]config.SignatureOption, error) {
	if flags.AddPage(cmd) {
		size, err := flags.PageSize(cmd)
		if err != nil {
			return nil, err
		}
		options = append(options, config.AddPage(size))
		options = append(options, config.PosStrict(flags.XposFlag.IsSet() || flags.YposFlag.IsSet()))
	} else {
		page := flags.Page(cmd)
		if flags.PageFlag.IsSet() {
			numpages, err := signer.GetPageCount(pdfReader)
			if err != nil {
				return nil, err
			}
			if page < 1 || page > numpages {
				return nil, eris.Errorf("invalid page number %d", page)
			}
		}
		options = append(options, config.Page(page))
		options = append(options, config.AddPage(nil))
	}
	return options, nil
}

var colorFlags = []struct {
	flag *cli.StringFlag
	fv   func(*cli.Command) (color.RGBA, error)
	fn   func(color.RGBA) config.SignatureOption
}{
	{flags.BackgroundColorFlag, flags.BackgroundColor, config.BackgroundColor},
	{flags.BorderColorFlag, flags.BorderColor, config.BorderColor},
	{flags.TitleColorFlag, flags.TitleColor, config.TitleColor},
	{flags.KeyColorFlag, flags.KeyColor, config.KeyColor},
	{flags.ValueColorFlag, flags.ValueColor, config.ValueColor},
}

func applyColors(options []config.SignatureOption, cmd *cli.Command) ([]config.SignatureOption, error) {
	for _, cf := range colorFlags {
		if !cf.flag.IsSet() {
			continue
		}
		if color, err := cf.fv(cmd); err != nil {
			return nil, err
		} else {
			options = append(options, cf.fn(color))
		}
	}
	return options, nil
}
