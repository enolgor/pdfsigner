package flags

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strings"

	"github.com/enolgor/pdfsigner/signer/config"
	"github.com/enolgor/pdfsigner/signer/fonts"
	"github.com/mazznoer/csscolorparser"
	"github.com/rotisserie/eris"
	"github.com/urfave/cli/v3"
)

// visible signature flags

var VisibleFlag = &cli.BoolFlag{
	Name:     "visible",
	Aliases:  []string{"v"},
	Value:    false,
	Usage:    "draw visible signature",
	Sources:  cli.EnvVars("VISIBLE"),
	Required: false,
	Category: "visible signature",
}

func Visible(cmd *cli.Command) bool {
	return cmd.Bool("visible")
}

var WidthFlag = &cli.Float64Flag{
	Name:     "width",
	Aliases:  []string{"w"},
	Value:    200,
	Usage:    "signature width in pt",
	Sources:  cli.EnvVars("WIDTH"),
	Required: false,
	Category: "visible signature",
}

func Width(cmd *cli.Command) float64 {
	return cmd.Float64("width")
}

var HeightFlag = &cli.Float64Flag{
	Name:     "height",
	Aliases:  []string{"h"},
	Value:    0,
	Usage:    "signature height in pt",
	Sources:  cli.EnvVars("HEIGHT"),
	Required: false,
	Category: "visible signature",
}

func Height(cmd *cli.Command) float64 {
	return cmd.Float64("height")
}

var XposFlag = &cli.Float64Flag{
	Name:     "xpos",
	Aliases:  []string{"x"},
	Value:    0,
	Usage:    "signature x position in pt (0 is left)",
	Sources:  cli.EnvVars("XPOS"),
	Required: false,
	Category: "visible signature",
}

func Xpos(cmd *cli.Command) float64 {
	return cmd.Float64("xpos")
}

var YposFlag = &cli.Float64Flag{
	Name:     "ypos",
	Aliases:  []string{"y"},
	Value:    0,
	Usage:    "signature y position in pt (0 is bottom)",
	Sources:  cli.EnvVars("YPOS"),
	Required: false,
	Category: "visible signature",
}

func Ypos(cmd *cli.Command) float64 {
	return cmd.Float64("ypos")
}

var RotateFlag = &cli.StringFlag{
	Name:     "rotate",
	Aliases:  []string{"r"},
	Value:    "0",
	Usage:    "signature rotation, one of 0, 90, 180, 270",
	Sources:  cli.EnvVars("ROTATE"),
	Required: false,
	Category: "visible signature",
	Validator: func(v string) error {
		switch v {
		case "0", "90", "180", "270":
			return nil
		default:
			return eris.Errorf("invalid rotation %s, must be one of 0, 90, 180, 270", v)
		}
	},
}

func Rotate(cmd *cli.Command) config.Rotation {
	return config.Rotation(cmd.String("rotate"))
}

var DpiFlag = &cli.Float64Flag{
	Name:     "dpi",
	Aliases:  []string{"i"},
	Value:    300,
	Usage:    "dpi",
	Sources:  cli.EnvVars("DPI"),
	Required: false,
	Category: "visible signature",
}

func Dpi(cmd *cli.Command) float64 {
	return cmd.Float64("dpi")
}

var TitleFlag = &cli.StringFlag{
	Name:     "title",
	Aliases:  []string{"t"},
	Value:    "DIGITALLY SIGNED",
	Usage:    "signature title",
	Sources:  cli.EnvVars("TITLE"),
	Required: false,
	Category: "visible signature",
}

var NoTitleFlag = &cli.BoolFlag{
	Name:     "no-title",
	Aliases:  []string{"nt"},
	Value:    false,
	Usage:    "Do not include title in signature (ignores any title flag)",
	Sources:  cli.EnvVars("NOTITLE"),
	Required: false,
	Category: "visible signature",
}

func Title(cmd *cli.Command) string {
	if cmd.Bool("no-title") {
		return ""
	}
	return cmd.String("title")
}

var DatetimeFormatFlag = &cli.StringFlag{
	Name:     "datetime-format",
	Aliases:  []string{"df"},
	Value:    "2006-01-02 15:04:05 -07:00",
	Usage:    "datetime format (in go time format)",
	Sources:  cli.EnvVars("DATETIMEFORMAT"),
	Required: false,
	Category: "visible signature",
}

func DatetimeFormat(cmd *cli.Command) string {
	return cmd.String("datetime-format")
}

var NoSubjectFlag = &cli.BoolFlag{
	Name:     "no-subject",
	Aliases:  []string{"ns"},
	Value:    false,
	Usage:    "Do not include subject in signature (ignores any subject flag)",
	Sources:  cli.EnvVars("NOSUBJECT"),
	Required: false,
	Category: "visible signature",
}

func NoSubject(cmd *cli.Command) bool {
	return cmd.Bool("no-subject")
}

var NoIssuerFlag = &cli.BoolFlag{
	Name:     "no-issuer",
	Aliases:  []string{"ni"},
	Value:    false,
	Usage:    "Do not include issuer in signature (ignores any issuer flag)",
	Sources:  cli.EnvVars("NOISSUER"),
	Required: false,
	Category: "visible signature",
}

func NoIssuer(cmd *cli.Command) bool {
	return cmd.Bool("no-issuer")
}

var NoDateFlag = &cli.BoolFlag{
	Name:     "no-date",
	Aliases:  []string{"nd"},
	Value:    false,
	Usage:    "Do not include date in signature (ignores any date flag)",
	Sources:  cli.EnvVars("NODATE"),
	Required: false,
	Category: "visible signature",
}

func NoDate(cmd *cli.Command) bool {
	return cmd.Bool("no-date")
}

var SubjectKeyFlag = &cli.StringFlag{
	Name:     "subject-key",
	Aliases:  []string{"sk"},
	Value:    "Subject:",
	Usage:    "subject key in signature",
	Sources:  cli.EnvVars("SUBJECTKEY"),
	Required: false,
	Category: "visible signature",
}

func SubjectKey(cmd *cli.Command) string {
	return cmd.String("subject-key")
}

var IssuerKeyFlag = &cli.StringFlag{
	Name:     "issuer-key",
	Aliases:  []string{"ik"},
	Value:    "Issuer:",
	Usage:    "issuer key in signature",
	Sources:  cli.EnvVars("ISSUERKEY"),
	Required: false,
	Category: "visible signature",
}

func IssuerKey(cmd *cli.Command) string {
	return cmd.String("issuer-key")
}

var DateKeyFlag = &cli.StringFlag{
	Name:     "date-key",
	Aliases:  []string{"dk"},
	Value:    "Date:",
	Usage:    "date key in signature",
	Sources:  cli.EnvVars("DATEKEY"),
	Required: false,
	Category: "visible signature",
}

func DateKey(cmd *cli.Command) string {
	return cmd.String("date-key")
}

var ExtraLinesFlag = &cli.StringSliceFlag{
	Name:     "extra-lines",
	Aliases:  []string{"el"},
	Value:    nil,
	Usage:    "extra lines in signature in the format key,value. Commas can be escaped with a backslash.",
	Sources:  cli.EnvVars("EXTRALINES"),
	Required: false,
	Category: "visible signature",
}

var re *regexp.Regexp = regexp.MustCompile(`.*?[^\\](,|$)`)
var cleanStr func(str string, last bool) string = func(str string, last bool) string {
	s := strings.TrimSpace(strings.ReplaceAll(str, "\\", ""))
	if last {
		return s
	}
	return strings.TrimSuffix(s, ",")
}

func ExtraLines(cmd *cli.Command) ([]config.TextLine, error) {
	extraLinesValue := cmd.StringSlice("extra-lines")
	if extraLinesValue == nil {
		return make([]config.TextLine, 0), nil
	}
	extraLines := make([]config.TextLine, len(extraLinesValue))

	for i, line := range cmd.StringSlice("extra-lines") {
		if !strings.Contains(line, ",") {
			extraLines[i] = config.TextLine{Key: "", Value: cleanStr(line, true)}
			continue
		}
		parts := re.FindAllString(line, -1)
		if len(parts) == 2 {
			extraLines[i] = config.TextLine{Key: cleanStr(parts[0], false), Value: cleanStr(parts[1], true)}
		} else {
			return nil, eris.Errorf("invalid extra line '%q', must be in the format key,value", line)
		}
	}
	return extraLines, nil
}

var BackgroundColorFlag = &cli.StringFlag{
	Name:     "background-color",
	Aliases:  []string{"bc"},
	Value:    "white",
	Usage:    "background color in signature, must be a valid CSS color",
	Sources:  cli.EnvVars("BACKGROUNDCOLOR"),
	Required: false,
	Category: "visible signature",
}

func BackgroundColor(cmd *cli.Command) (rgba color.RGBA, err error) {
	if rgba, err = parseColor(cmd.String("background-color")); err != nil {
		err = eris.Wrap(err, "error parsing background color")
		return
	}
	return
}

var BorderSizeFlag = &cli.FloatFlag{
	Name:     "border-size",
	Aliases:  []string{"rs"},
	Value:    1,
	Usage:    "border size in ot",
	Sources:  cli.EnvVars("BORDERSIZE"),
	Required: false,
	Category: "visible signature",
}

func BorderSize(cmd *cli.Command) float64 {
	return cmd.Float64("border-size")
}

var BorderColorFlag = &cli.StringFlag{
	Name:     "border-color",
	Aliases:  []string{"rc"},
	Value:    "black",
	Usage:    "border color in signature, must be a valid CSS color",
	Sources:  cli.EnvVars("BORDERCOLOR"),
	Required: false,
	Category: "visible signature",
}

func BorderColor(cmd *cli.Command) (rgba color.RGBA, err error) {
	if rgba, err = parseColor(cmd.String("border-color")); err != nil {
		err = eris.Wrap(err, "error parsing border color")
		return
	}
	return
}

var LogoFlag = &cli.StringFlag{
	Name:     "logo",
	Aliases:  []string{"l"},
	Value:    "",
	Usage:    "set logo path to use in signature (png only)",
	Sources:  cli.EnvVars("LOGO"),
	Required: false,
	Category: "signature",
}

func Logo(cmd *cli.Command) (image.Image, error) {
	if !LogoFlag.IsSet() {
		return nil, nil
	}
	f, err := os.Open(cmd.String("logo"))
	if err != nil {
		return nil, eris.Wrapf(err, "failed to read file %s", cmd.String("logo"))
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to decode png file %s", cmd.String("logo"))
	}
	return img, nil
}

var LogoGrayscaleFlag = &cli.BoolFlag{
	Name:     "logo-grayscale",
	Aliases:  []string{"lg"},
	Value:    false,
	Usage:    "convert logo to grayscale",
	Sources:  cli.EnvVars("LOGOGRAYSCALE"),
	Required: false,
	Category: "signature",
}

func LogoGrayscale(cmd *cli.Command) bool {
	return cmd.Bool("logo-grayscale")
}

var LogoOpacityFlag = &cli.Float64Flag{
	Name:     "logo-opacity",
	Aliases:  []string{"lo"},
	Value:    0.25,
	Usage:    "set logo opacity (0.0 to 1.0)",
	Sources:  cli.EnvVars("LOGOOPACITY"),
	Required: false,
	Category: "signature",
}

func LogoOpacity(cmd *cli.Command) float64 {
	return cmd.Float64("logo-opacity")
}

var LogoAlignmentFlag = &cli.StringFlag{
	Name:     "logo-alignment",
	Aliases:  []string{"la"},
	Value:    "center",
	Usage:    "logo alignment, one of left, center, right",
	Sources:  cli.EnvVars("LOGOALIGNMENT"),
	Required: false,
	Category: "signature",
	Validator: func(v string) error {
		switch v {
		case "left", "center", "right":
			return nil
		default:
			return eris.Errorf("invalid logo alignment %s, must be one of left, center, right", v)
		}
	},
}

func LogoAlignment(cmd *cli.Command) config.Alignment {
	return config.Alignment(cmd.String("logo-alignment"))
}

var TitleAlignmentFlag = &cli.StringFlag{
	Name:     "title-alignment",
	Aliases:  []string{"ta"},
	Value:    "center",
	Usage:    "title alignment, one of left, center, right",
	Sources:  cli.EnvVars("TITLEALIGNMENT"),
	Required: false,
	Category: "signature",
	Validator: func(v string) error {
		switch v {
		case "left", "center", "right":
			return nil
		default:
			return eris.Errorf("invalid title alignment %s, must be one of left, center, right", v)
		}
	},
}

var NoEmptyLineAfterTitleFlag = &cli.BoolFlag{
	Name:     "no-empty-line-after-title",
	Aliases:  []string{"nelt"},
	Value:    false,
	Usage:    "Do not add an empty line after the title",
	Sources:  cli.EnvVars("NOEMPTYLINEAFTERTITLE"),
	Required: false,
	Category: "signature",
}

func NoEmptyLineAfterTitle(cmd *cli.Command) bool {
	return cmd.Bool("no-empty-line-after-title")
}

func TitleAlignment(cmd *cli.Command) config.Alignment {
	return config.Alignment(cmd.String("title-alignment"))
}

var LineAlignmentFlag = &cli.StringFlag{
	Name:     "line-alignment",
	Aliases:  []string{"lia"},
	Value:    "center",
	Usage:    "line alignment, one of left, center, right. Overrides key and value alignment.",
	Sources:  cli.EnvVars("LINEALIGNMENT"),
	Required: false,
	Category: "signature",
	Validator: func(v string) error {
		switch v {
		case "left", "center", "right":
			return nil
		default:
			return eris.Errorf("invalid line alignment %s, must be one of left, center, right", v)
		}
	},
}

func LineAlignment(cmd *cli.Command) config.Alignment {
	return config.Alignment(cmd.String("line-alignment"))
}

var KeyAlignmentFlag = &cli.StringFlag{
	Name:     "key-alignment",
	Aliases:  []string{"ka"},
	Value:    "left",
	Usage:    "key column alignment, one of left, center, right",
	Sources:  cli.EnvVars("KEYALIGNMENT"),
	Required: false,
	Category: "signature",
	Validator: func(v string) error {
		switch v {
		case "left", "center", "right":
			return nil
		default:
			return eris.Errorf("invalid key alignment %s, must be one of left, center, right", v)
		}
	},
}

func KeyAlignment(cmd *cli.Command) config.Alignment {
	return config.Alignment(cmd.String("key-alignment"))
}

var ValueAlignmentFlag = &cli.StringFlag{
	Name:     "value-alignment",
	Aliases:  []string{"va"},
	Value:    "right",
	Usage:    "value column alignment, one of left, center, right",
	Sources:  cli.EnvVars("VALUEALIGNMENT"),
	Required: false,
	Category: "signature",
	Validator: func(v string) error {
		switch v {
		case "left", "center", "right":
			return nil
		default:
			return eris.Errorf("invalid value alignment %s, must be one of left, center, right", v)
		}
	},
}

func ValueAlignment(cmd *cli.Command) config.Alignment {
	return config.Alignment(cmd.String("value-alignment"))
}

var LoadFontFlag = &cli.StringSliceFlag{
	Name:     "load-font",
	Aliases:  []string{"lf"},
	Value:    nil,
	Usage:    "path to custom ttf font file",
	Sources:  cli.EnvVars("LOADFONT"),
	Required: false,
	Category: "signature",
}

func LoadFonts(cmd *cli.Command) error {
	if cmd.StringSlice("load-font") == nil {
		return nil
	}
	for _, fontPath := range cmd.StringSlice("load-font") {
		if fontPath == "" {
			continue
		}
		if err := fonts.LoadCustomFont(fontPath); err != nil {
			return eris.Wrapf(err, "failed to load custom font %s", fontPath)
		}
	}
	return nil
}

var TitleFontFlag = &cli.StringFlag{
	Name:     "title-font",
	Aliases:  []string{"tf"},
	Value:    "RobotoMono-Bold",
	Usage:    "ttf font file name, without extension (use list-fonts to see available fonts)",
	Sources:  cli.EnvVars("TITLEFONT"),
	Required: false,
	Category: "signature",
}

func TitleFont(cmd *cli.Command) string {
	return cmd.String("title-font")
}

var KeyFontFlag = &cli.StringFlag{
	Name:     "key-font",
	Aliases:  []string{"kf"},
	Value:    "RobotoMono-SemiBold",
	Usage:    "ttf font file name, without extension (use list-fonts to see available fonts)",
	Sources:  cli.EnvVars("KEYFONT"),
	Required: false,
	Category: "signature",
}

func KeyFont(cmd *cli.Command) string {
	return cmd.String("key-font")
}

var ValueFontFlag = &cli.StringFlag{
	Name:     "value-font",
	Aliases:  []string{"vf"},
	Value:    "RobotoMono-Regular",
	Usage:    "ttf font file name, without extension (use list-fonts to see available fonts)",
	Sources:  cli.EnvVars("VALUEFONT"),
	Required: false,
	Category: "signature",
}

func ValueFont(cmd *cli.Command) string {
	return cmd.String("value-font")
}

var TitleColorFlag = &cli.StringFlag{
	Name:     "title-color",
	Aliases:  []string{"tc"},
	Value:    "black",
	Usage:    "title color in signature, must be a valid CSS color",
	Sources:  cli.EnvVars("TITLECOLOR"),
	Required: false,
	Category: "visible signature",
}

func TitleColor(cmd *cli.Command) (rgba color.RGBA, err error) {
	if rgba, err = parseColor(cmd.String("title-color")); err != nil {
		err = eris.Wrap(err, "error parsing title color")
		return
	}
	return
}

var KeyColorFlag = &cli.StringFlag{
	Name:     "key-color",
	Aliases:  []string{"kc"},
	Value:    "black",
	Usage:    "key color in signature, must be a valid CSS color",
	Sources:  cli.EnvVars("KEYCOLOR"),
	Required: false,
	Category: "visible signature",
}

func KeyColor(cmd *cli.Command) (rgba color.RGBA, err error) {
	if rgba, err = parseColor(cmd.String("key-color")); err != nil {
		err = eris.Wrap(err, "error parsing key color")
		return
	}
	return
}

var ValueColorFlag = &cli.StringFlag{
	Name:     "value-color",
	Aliases:  []string{"vc"},
	Value:    "black",
	Usage:    "value color in signature, must be a valid CSS color",
	Sources:  cli.EnvVars("VALUECOLOR"),
	Required: false,
	Category: "visible signature",
}

func ValueColor(cmd *cli.Command) (rgba color.RGBA, err error) {
	if rgba, err = parseColor(cmd.String("value-color")); err != nil {
		err = eris.Wrap(err, "error parsing value color")
		return
	}
	return
}

func parseColor(css string) (rgba color.RGBA, err error) {
	var csscolor csscolorparser.Color
	if csscolor, err = csscolorparser.Parse(css); err != nil {
		err = eris.Wrapf(err, "error parsing background color %s", css)
		return
	}
	rgba.R, rgba.G, rgba.B, rgba.A = csscolor.RGBA255()
	return
}
