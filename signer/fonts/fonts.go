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

package fonts

import (
	"embed"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/rotisserie/eris"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed data
var embeddedData embed.FS

var fontCache map[string]*opentype.Font = make(map[string]*opentype.Font)
var embeddedFonts map[string]string = make(map[string]string)
var systemFonts map[string]string = make(map[string]string)
var customFonts map[string]string = make(map[string]string)

type LoadedFont string

func (f LoadedFont) String() string {
	source := "unavailable"
	if _, ok := customFonts[string(f)]; ok {
		source = "custom"
	} else if _, ok := embeddedFonts[string(f)]; ok {
		source = "embedded"
	} else if _, ok := systemFonts[string(f)]; ok {
		source = "system"
	}
	return string(f) + " (" + source + ")"
}

var foundFonts []LoadedFont = make([]LoadedFont, 0)

var ErrFontNotFound error = eris.New("font not found")
var ErrFontFailedToLoad error = eris.New("font failed to load")

func init() {
	fs.WalkDir(embeddedData, "data", walkFn(embeddedFonts))
	for _, dir := range getSystemFontDirectories() {
		filepath.WalkDir(dir, walkFn(systemFonts))
	}
}

func walkFn(fontMap map[string]string) func(path string, d fs.DirEntry, err error) error {
	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			loadFont(fontMap, path)
		}
		return nil
	}
}

func loadFont(fontMap map[string]string, path string) {
	ext := filepath.Ext(path)
	if ext != ".ttf" {
		return
	}
	name := strings.TrimSuffix(filepath.Base(path), ext)
	fontMap[name] = path
	if !slices.Contains(foundFonts, LoadedFont(name)) {
		foundFonts = append(foundFonts, LoadedFont(name))
	}
}

func ListLoadedFonts() []LoadedFont {
	return foundFonts
}

func LoadCustomFont(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	if _, err := os.Stat(absPath); err != nil {
		return eris.Wrapf(err, "failed to read file %s", absPath)
	}
	loadFont(customFonts, absPath)
	return nil
}

func readFont(name string) (data []byte, err error) {
	var found bool
	data, found, err = readFontFromTable(customFonts, os.ReadFile, name)
	if !found {
		data, found, err = readFontFromTable(embeddedFonts, embeddedData.ReadFile, name)
	}
	if !found {
		data, found, err = readFontFromTable(systemFonts, os.ReadFile, name)
	}
	if !found {
		err = ErrFontNotFound
	} else if err != nil {
		err = errors.Join(ErrFontFailedToLoad, err)
	}
	return
}

func readFontFromTable(table map[string]string, readfile func(name string) ([]byte, error), name string) (data []byte, found bool, err error) {
	var path string
	if path, found = table[name]; found {
		data, err = readfile(path)
	}
	return
}

func LoadFontFace(name string, dpi float64, size float64) (font.Face, error) {
	var v *opentype.Font
	var ok bool
	if v, ok = fontCache[name]; !ok {
		data, err := readFont(name)
		if err != nil {
			return nil, err
		}
		if v, err = opentype.Parse(data); err != nil {
			return nil, errors.Join(ErrFontFailedToLoad, err)
		}
	}
	fontCache[name] = v
	return opentype.NewFace(v, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}
