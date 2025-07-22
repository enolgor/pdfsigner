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
