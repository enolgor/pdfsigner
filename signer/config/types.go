package config

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image"
	"image/png"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Dim = types.Dim

var PaperSize map[string]*Dim = types.PaperSize

type Rotation string

const (
	ROTATE_0   Rotation = "0"
	ROTATE_90  Rotation = "90"
	ROTATE_180 Rotation = "180"
	ROTATE_270 Rotation = "270"
)

type Alignment string

const (
	LEFT   Alignment = "left"
	CENTER Alignment = "center"
	RIGHT  Alignment = "right"
)

type TextLine struct {
	Key   string
	Value string
}

type JImage struct {
	Image image.Image
}

func (ji JImage) MarshalJSON() ([]byte, error) {
	if ji.Image == nil {
		return json.Marshal(nil)
	}
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, ji.Image); err != nil {
		return nil, err
	}
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	return json.Marshal(encoded)
}

func (ji *JImage) UnmarshalJSON(data []byte) error {
	var b64 string
	if err := json.Unmarshal(data, &b64); err != nil {
		return err
	}
	decoded, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return err
	}
	img, err := png.Decode(bytes.NewReader(decoded))
	if err != nil {
		return err
	}
	ji.Image = img
	return nil
}
