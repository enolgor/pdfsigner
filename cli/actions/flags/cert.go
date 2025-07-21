package flags

import "github.com/urfave/cli/v3"

// certificate flags

var CertFlag = &cli.StringFlag{
	Name:      "cert",
	Aliases:   []string{"c"},
	Value:     "",
	Usage:     "path to the pkcs12 certificate file",
	Sources:   cli.EnvVars("CERT"),
	Required:  true,
	TakesFile: true,
	Category:  "certificate",
}

func Cert(cmd *cli.Command) string {
	return cmd.String("cert")
}

var PassphraseFlag = &cli.StringFlag{
	Name:     "passphrase",
	Aliases:  []string{"s"},
	Value:    "",
	Usage:    "passphrase OF the pkcs12 certificate file",
	Sources:  cli.EnvVars("PASSPHRASE"),
	Required: true,
	Category: "certificate",
}

func Passphrase(cmd *cli.Command) string {
	return cmd.String("passphrase")
}
