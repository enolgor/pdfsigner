package actions

import (
	"context"
	"fmt"

	"github.com/enolgor/pdfsigner/cli/actions/flags"
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
