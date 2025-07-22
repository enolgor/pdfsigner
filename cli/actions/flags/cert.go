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

import "github.com/urfave/cli/v3"

// certificate flags

const certCategory = "certificate"

var CertFlag = &cli.StringFlag{
	Name:      "cert",
	Aliases:   []string{"c"},
	Value:     "",
	Usage:     "path to the pkcs12 certificate file",
	Sources:   cli.EnvVars("CERT"),
	Required:  true,
	TakesFile: true,
	Category:  certCategory,
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
	Category: certCategory,
}

func Passphrase(cmd *cli.Command) string {
	return cmd.String("passphrase")
}
