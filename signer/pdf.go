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
