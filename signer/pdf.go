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

package signer

import (
	"bytes"
	"strconv"

	"github.com/enolgor/pdfsigner/signer/config"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/rotisserie/eris"
)

func addLastPage(pdfReader *bytes.Reader, dim *config.Dim) (*bytes.Reader, int, error) {
	count, err := GetPageCount(pdfReader)
	if err != nil {
		return nil, 0, err
	}
	buff := new(bytes.Buffer)
	pageConf := &pdfcpu.PageConfiguration{
		PageDim: dim,
		InpUnit: types.POINTS,
	}
	if err := api.InsertPages(pdfReader, buff, []string{strconv.Itoa(count)}, false, pageConf, nil); err != nil {
		return nil, 0, err
	}
	return bytes.NewReader(buff.Bytes()), count + 1, err
}

func GetPageCount(r *bytes.Reader) (int, error) {
	count, err := api.PageCount(r, nil)
	return count, eris.Wrap(err, "failed to get page count")
}

func GetPageDimensionsPt(r *bytes.Reader, pageNum int) (float64, float64, error) {
	if pageNum < 0 {
		return 0, 0, eris.New("invalid page number")
	}
	dims, err := api.PageDims(r, nil)
	if err != nil {
		return 0, 0, eris.Wrap(err, "failed to get page dimensions")
	}
	if pageNum >= len(dims) {
		return 0, 0, eris.New("invalid page number")
	}
	return dims[pageNum].Width, dims[pageNum].Height, nil
}
