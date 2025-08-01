package stamps

import (
	"image/color"

	"github.com/enolgor/pdfsigner/signer/config"
)

type RGBA = color.RGBA
type TextLine = config.TextLine
type Rotation = config.Rotation
type Alignment = config.Alignment

type StampConfig struct {
	Title               string     `json:"title"`
	DateFormat          string     `json:"dateFormat"`
	IncludeSubject      bool       `json:"includeSubject"`
	IncludeIssuer       bool       `json:"includeIssuer"`
	IncludeDate         bool       `json:"includeDate"`
	SubjectKey          string     `json:"subjectKey"`
	IssuerKey           string     `json:"issuerKey"`
	DateKey             string     `json:"dateKey"`
	ExtraLines          []TextLine `json:"extraLines,omitempty"`
	Dpi                 float64    `json:"dpi"`
	BackgroundColor     RGBA       `json:"backgroundColor"`
	WidthPt             float64    `json:"widthPt"`
	HeightPt            float64    `json:"heightPt"`
	PosXPt              float64    `json:"posXPt"`
	PosYPt              float64    `json:"posYPt"`
	PosStrict           bool       `json:"posStrict"`
	Rotate              Rotation   `json:"rotate"`
	BorderSizePt        float64    `json:"borderSizePt"`
	BorderColor         RGBA       `json:"borderColor"`
	Logo                string     `json:"logo,omitempty"`
	LogoOpacity         float64    `json:"logoOpacity"`
	LogoGrayScale       bool       `json:"logoGrayScale"`
	LogoAlignment       Alignment  `json:"logoAlignment"`
	EmptyLineAfterTitle bool       `json:"emptyLineAfterTitle"`
	TitleAlignment      Alignment  `json:"titleAlignment"`
	LineAlignment       Alignment  `json:"lineAlignment"`
	KeyAlignment        Alignment  `json:"keyAlignment"`
	ValueAlignment      Alignment  `json:"valueAlignment"`
	TitleFont           string     `json:"titleFont"`
	KeyFont             string     `json:"keyFont"`
	ValueFont           string     `json:"valueFont"`
	TitleColor          RGBA       `json:"titleColor"`
	KeyColor            RGBA       `json:"keyColor"`
	ValueColor          RGBA       `json:"valueColor"`
}

const (
	ROTATE_0   Rotation = "0"
	ROTATE_90  Rotation = "90"
	ROTATE_180 Rotation = "180"
	ROTATE_270 Rotation = "270"
)

var AllRotations = []struct {
	Value  Rotation
	TSName string
}{
	{ROTATE_0, "ROTATE_0"},
	{ROTATE_90, "ROTATE_90"},
	{ROTATE_180, "ROTATE_180"},
	{ROTATE_270, "ROTATE_270"},
}

const (
	LEFT   Alignment = "left"
	CENTER Alignment = "center"
	RIGHT  Alignment = "right"
)

var AllAlignments = []struct {
	Value  Alignment
	TSName string
}{
	{LEFT, "LEFT"},
	{CENTER, "CENTER"},
	{RIGHT, "RIGHT"},
}

func (sc *StampConfig) FromConfig(cfg *config.SignatureConfiguration) (err error) {
	sc.Title = cfg.Title
	sc.DateFormat = cfg.DateFormat
	sc.IncludeSubject = cfg.IncludeSubject
	sc.IncludeIssuer = cfg.IncludeIssuer
	sc.IncludeDate = cfg.IncludeDate
	sc.SubjectKey = cfg.SubjectKey
	sc.IssuerKey = cfg.IssuerKey
	sc.DateKey = cfg.DateKey
	sc.ExtraLines = cfg.ExtraLines
	sc.Dpi = cfg.Dpi
	sc.BackgroundColor = cfg.BackgroundColor
	sc.WidthPt = cfg.WidthPt
	sc.HeightPt = cfg.HeightPt
	sc.PosXPt = cfg.PosXPt
	sc.PosYPt = cfg.PosYPt
	sc.PosStrict = cfg.PosStrict
	sc.Rotate = cfg.Rotate
	sc.BorderSizePt = cfg.BorderSizePt
	sc.BorderColor = cfg.BorderColor
	if cfg.Logo == nil {
		sc.Logo = ""
	} else {
		data, err := cfg.Logo.MarshalJSON()
		if err != nil {
			return err
		}
		sc.Logo = string(data)
	}
	sc.LogoOpacity = cfg.LogoOpacity
	sc.LogoGrayScale = cfg.LogoGrayScale
	sc.LogoAlignment = cfg.LogoAlignment
	sc.EmptyLineAfterTitle = cfg.EmptyLineAfterTitle
	sc.TitleAlignment = cfg.TitleAlignment
	sc.LineAlignment = cfg.LineAlignment
	sc.KeyAlignment = cfg.KeyAlignment
	sc.ValueAlignment = cfg.ValueAlignment
	sc.TitleFont = cfg.TitleFont
	sc.KeyFont = cfg.KeyFont
	sc.ValueFont = cfg.ValueFont
	sc.TitleColor = cfg.TitleColor
	sc.KeyColor = cfg.KeyColor
	sc.ValueColor = cfg.ValueColor
	return
}

func (sc *StampConfig) ToConfig() (*config.SignatureConfiguration, error) {
	cfg := &config.SignatureConfiguration{}
	cfg.Title = sc.Title
	cfg.DateFormat = sc.DateFormat
	cfg.IncludeSubject = sc.IncludeSubject
	cfg.IncludeIssuer = sc.IncludeIssuer
	cfg.IncludeDate = sc.IncludeDate
	cfg.SubjectKey = sc.SubjectKey
	cfg.IssuerKey = sc.IssuerKey
	cfg.DateKey = sc.DateKey
	cfg.ExtraLines = sc.ExtraLines
	cfg.Dpi = sc.Dpi
	cfg.BackgroundColor = sc.BackgroundColor
	cfg.WidthPt = sc.WidthPt
	cfg.HeightPt = sc.HeightPt
	cfg.PosXPt = sc.PosXPt
	cfg.PosYPt = sc.PosYPt
	cfg.PosStrict = sc.PosStrict
	cfg.Rotate = sc.Rotate
	cfg.BorderSizePt = sc.BorderSizePt
	cfg.BorderColor = sc.BorderColor
	if sc.Logo == "" {
		cfg.Logo = nil
	} else {
		ji := &config.JImage{}
		if err := ji.UnmarshalJSON([]byte(sc.Logo)); err != nil {
			return nil, err
		}
	}
	cfg.LogoOpacity = sc.LogoOpacity
	cfg.LogoGrayScale = sc.LogoGrayScale
	cfg.LogoAlignment = sc.LogoAlignment
	cfg.EmptyLineAfterTitle = sc.EmptyLineAfterTitle
	cfg.TitleAlignment = sc.TitleAlignment
	cfg.LineAlignment = sc.LineAlignment
	cfg.KeyAlignment = sc.KeyAlignment
	cfg.ValueAlignment = sc.ValueAlignment
	cfg.TitleFont = sc.TitleFont
	cfg.KeyFont = sc.KeyFont
	cfg.ValueFont = sc.ValueFont
	cfg.TitleColor = sc.TitleColor
	cfg.KeyColor = sc.KeyColor
	cfg.ValueColor = sc.ValueColor
	return cfg, nil
}
