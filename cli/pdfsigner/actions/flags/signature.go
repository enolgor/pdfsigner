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
	"io"
	"os"
	"time"

	"github.com/rotisserie/eris"
	"github.com/urfave/cli/v3"
)

// signature flags

const signatureCategory = "signature"

var SignedOutputFlag = &cli.StringFlag{
	Name:     "out",
	Aliases:  []string{"o"},
	Value:    "",
	Usage:    "set filename and path of signed pdf, otherwise stdout is used",
	Sources:  cli.EnvVars("OUT"),
	Required: false,
	Category: signatureCategory,
}

func SignedOutput(cmd *cli.Command) (io.WriteCloser, error) {
	path := cmd.String(SignedOutputFlag.Name)
	if path == "" {
		return os.Stdout, nil
	}
	fileflags := os.O_CREATE | os.O_WRONLY
	if ForceWrite(cmd) {
		fileflags |= os.O_TRUNC
	} else {
		fileflags |= os.O_EXCL
	}
	f, err := os.OpenFile(path, fileflags, os.ModePerm)
	if err != nil {
		if os.IsExist(err) {
			return nil, eris.Wrapf(err, "file %s already exists", path)
		}
		return nil, eris.Wrapf(err, "failed to create file %s", path)
	}
	return f, nil
}

var ForceWriteFlag = &cli.BoolFlag{
	Name:     "force",
	Aliases:  []string{"f"},
	Value:    false,
	Usage:    "force overwrite of existing file",
	Sources:  cli.EnvVars("FORCE"),
	Required: false,
	Category: signatureCategory,
}

func ForceWrite(cmd *cli.Command) bool {
	return cmd.Bool(ForceWriteFlag.Name)
}

var DatetimeFlag *cli.TimestampFlag = &cli.TimestampFlag{
	Name:    "datetime",
	Aliases: []string{"d"},
	Value:   time.Now(),
	Usage:   "date and time in RFC3339 format (defaults to now)",
	Sources: cli.EnvVars("DATETIME"),
	Config: cli.TimestampConfig{
		Layouts: []string{time.RFC3339},
	},
	HideDefault: true,
	Category:    signatureCategory,
}

func Datetime(cmd *cli.Command) time.Time {
	return cmd.Timestamp(DatetimeFlag.Name)
}

var LocationFlag *cli.StringFlag = &cli.StringFlag{
	Name:     "location",
	Aliases:  []string{"l"},
	Value:    "UTC",
	Usage:    "datetime zone location (defaults to UTC)",
	Sources:  cli.EnvVars("LOCATION"),
	Required: false,
	Category: signatureCategory,
}

func Location(cmd *cli.Command) (*time.Location, error) {
	loc, err := time.LoadLocation(cmd.String(LocationFlag.Name))
	if err != nil {
		return nil, eris.Wrapf(err, "location %s not found", cmd.String(LocationFlag.Name))
	}
	return loc, nil
}

var SignatureNameFlag = &cli.StringFlag{
	Name:     "signature-name",
	Aliases:  []string{"sn"},
	Value:    "",
	Usage:    "signature name",
	Sources:  cli.EnvVars("SIGNATURE_NAME"),
	Required: false,
	Category: signatureCategory,
}

func SignatureName(cmd *cli.Command) string {
	return cmd.String(SignatureNameFlag.Name)
}

var SignatureReasonFlag = &cli.StringFlag{
	Name:     "signature-reason",
	Aliases:  []string{"sr"},
	Value:    "",
	Usage:    "signature reason",
	Sources:  cli.EnvVars("SIGNATURE_REASON"),
	Required: false,
	Category: signatureCategory,
}

func SignatureReason(cmd *cli.Command) string {
	return cmd.String(SignatureReasonFlag.Name)
}

var SignatureLocationFlag = &cli.StringFlag{
	Name:     "signature-location",
	Aliases:  []string{"sl"},
	Value:    "",
	Usage:    "signature location",
	Sources:  cli.EnvVars("SIGNATURE_LOCATION"),
	Required: false,
	Category: signatureCategory,
}

func SignatureLocation(cmd *cli.Command) string {
	return cmd.String(SignatureLocationFlag.Name)
}

var SignatureContactFlag = &cli.StringFlag{
	Name:     "signature-contact",
	Aliases:  []string{"sc"},
	Value:    "",
	Usage:    "signature contact details",
	Sources:  cli.EnvVars("SIGNATURE_CONTACT"),
	Required: false,
	Category: signatureCategory,
}

func SignatureContact(cmd *cli.Command) string {
	return cmd.String(SignatureContactFlag.Name)
}

var TsaURLFlag = &cli.StringFlag{
	Name:     "tsa-url",
	Value:    "",
	Usage:    "URL of the timestamp authority (TSA) to use for signing",
	Sources:  cli.EnvVars("TSA_URL"),
	Required: false,
	Category: signatureCategory,
}

func TsaURL(cmd *cli.Command) string {
	return cmd.String(TsaURLFlag.Name)
}

var TsaUserFlag = &cli.StringFlag{
	Name:     "tsa-user",
	Value:    "",
	Usage:    "Username for the TSA, if required",
	Sources:  cli.EnvVars("TSA_USER"),
	Required: false,
	Category: signatureCategory,
}

func TsaUser(cmd *cli.Command) string {
	return cmd.String(TsaUserFlag.Name)
}

var TsaPasswordFlag = &cli.StringFlag{
	Name:     "tsa-password",
	Value:    "",
	Usage:    "Password for the TSA, if required",
	Sources:  cli.EnvVars("TSA_PASSWORD"),
	Required: false,
	Category: signatureCategory,
}

func TsaPassword(cmd *cli.Command) string {
	return cmd.String(TsaPasswordFlag.Name)
}
