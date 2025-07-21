package draw

import (
	"image"
	"image/color"
	"math"
	"unicode/utf8"

	"github.com/enolgor/pdfsigner/signer/config"
	"github.com/enolgor/pdfsigner/signer/fonts"
	"github.com/fogleman/gg"
	"github.com/rotisserie/eris"
	"golang.org/x/image/font"
)

type rect struct {
	ypadpt   float64
	xpadpt   float64
	vspacept float64
}

func (r *rect) Draw(text []config.TextLine, conf *config.SignatureConfiguration) (image.Image, error) {
	var titleFace font.Face
	keyFace, valFace, fontSize, textwidth, textheight, err := r.calculateFontsAndUnpaddedSizesToFitPixels(text, conf)
	if err != nil {
		return nil, err
	}
	if titleFace, err = fonts.LoadFontFace(conf.TitleFont, conf.Dpi, fontSize); err != nil {
		return nil, err
	}
	xpad, ypad, vspace := r.getPaddings(text, conf)
	unpaddedBounds := r.getUnpaddedBounds(r.countLines(text, conf), textwidth, textheight)
	imageBounds := r.getImageBounds(unpaddedBounds, xpad, ypad, vspace)
	lineHeight := float64(unpaddedBounds.Dy()) / float64(r.countLines(text, conf))
	dc := gg.NewContextForRGBA(image.NewRGBA(imageBounds))
	dc.SetColor(conf.BackgroundColor)
	dc.Clear()
	r.drawBorder(dc, conf)
	if conf.Logo != nil && conf.Logo.Image != nil {
		logo, err := redrawLogo(conf.Logo.Image, unpaddedBounds, conf.LogoOpacity, conf.LogoGrayScale)
		if err != nil {
			return nil, err
		}
		var logox int
		switch conf.LogoAlignment {
		case config.LEFT:
			logox = int(xpad)
		case config.CENTER:
			logox = (imageBounds.Dx() - logo.Bounds().Dx()) / 2
		case config.RIGHT:
			logox = imageBounds.Dx() - logo.Bounds().Dx() - int(xpad)
		default:
			logox = int(xpad)
		}
		dc.DrawImage(logo, logox, int(ypad))
	}

	dc.SetFontFace(keyFace)
	longestKeyW, _ := dc.MeasureString(r.longest(text, lineKey))
	dc.SetFontFace(valFace)
	longestValueW, _ := dc.MeasureString(r.longest(text, lineValue))
	c := 1
	if conf.Title != "" {
		r.writeTitle(dc, conf.Title, xpad, ypad, lineHeight, titleFace, conf.TitleColor, conf.TitleAlignment, imageBounds)
		if conf.EmptyLineAfterTitle {
			c = 3
		} else {
			c = 2
		}
	}
	for i, text := range text {
		r.writeLine(dc, text, xpad, ypad, vspace, longestKeyW, longestValueW, i+c, lineHeight, keyFace, valFace, conf.KeyColor, conf.ValueColor, conf.LineAlignment, conf.KeyAlignment, conf.ValueAlignment)
	}
	return dc.Image(), nil
}

func (r *rect) drawBorder(dc *gg.Context, conf *config.SignatureConfiguration) {
	if conf.BorderSizePt > 0 {
		borderSize := PtsToPixels(conf.BorderSizePt, conf.Dpi)
		dc.SetLineWidth(borderSize)
		dc.SetLineCapSquare()
		dc.SetStrokeStyle(gg.NewSolidPattern(conf.BorderColor))
		x, y := borderSize/2, borderSize/2
		w, h := float64(dc.Width())-borderSize, float64(dc.Height())-borderSize
		dc.DrawLine(x, y, x+w, y)
		dc.DrawLine(x+w, y, x+w, y+h)
		dc.DrawLine(x+w, y+h, x, y+h)
		dc.DrawLine(x, y+h, x, y)
		//dc.DrawRectangle(x, y, w, h) //always rounds at least 3 corners
		dc.Stroke()
	}
}

func (r *rect) writeLine(dc *gg.Context, text config.TextLine, xpad, ypad, vspace float64, longestKeyW, longestValueW float64, n int, lineHeight float64, keyfont, valuefont font.Face, keyColor, valueColor color.Color, lineAlignment, keyAlignment, valueAlignment config.Alignment) {
	dc.SetFontFace(keyfont)
	keyW, _ := dc.MeasureString(text.Key)
	dc.SetFontFace(valuefont)
	valueW, _ := dc.MeasureString(text.Value)
	var keyX, valueX float64
	switch lineAlignment {
	case config.LEFT:
		keyX = xpad
		valueX = keyW + xpad + vspace
	case config.RIGHT:
		valueX = float64(dc.Width()) - valueW - xpad
		keyX = valueX - vspace - keyW
	default:
		switch keyAlignment {
		case config.CENTER:
			keyX = xpad + (longestKeyW-keyW)/2
		case config.RIGHT:
			keyX = xpad + longestKeyW - keyW
		default:
			keyX = xpad
		}
		switch valueAlignment {
		case config.CENTER:
			valueX = longestKeyW + xpad + vspace + (longestValueW-valueW)/2
		case config.RIGHT:
			valueX = longestKeyW + xpad + vspace + longestValueW - valueW
		default:
			valueX = longestKeyW + xpad + vspace
		}
	}
	y := float64(n)*lineHeight - lineHeight*0.25 + ypad
	dc.SetColor(keyColor)
	dc.SetFontFace(keyfont)
	dc.DrawString(text.Key, keyX, y)
	dc.SetColor(valueColor)
	dc.SetFontFace(valuefont)
	dc.DrawString(text.Value, valueX, y)
}

func (r *rect) writeTitle(dc *gg.Context, title string, xpad, ypad, lineHeight float64, font font.Face, color color.Color, alignment config.Alignment, bounds image.Rectangle) {
	dc.SetFontFace(font)
	var titlex float64
	switch alignment {
	case config.LEFT:
		titlex = xpad
	case config.CENTER:
		titlew, _ := dc.MeasureString(title)
		titlex = (float64(bounds.Dx()) - titlew) / 2
	case config.RIGHT:
		titlew, _ := dc.MeasureString(title)
		titlex = float64(bounds.Dx()) - titlew - xpad
	}
	y := lineHeight - lineHeight*0.25 + ypad
	dc.SetColor(color)
	dc.DrawString(title, titlex, y)
}

func (r *rect) CalculateExactPixelSize(text []config.TextLine, conf *config.SignatureConfiguration) (float64, float64, error) {
	_, _, _, textwidth, textheight, err := r.calculateFontsAndUnpaddedSizesToFitPixels(text, conf)
	xpad, ypad, vspace := r.getPaddings(text, conf)
	unpaddedBounds := r.getUnpaddedBounds(r.countLines(text, conf), textwidth, textheight)
	imageBounds := r.getImageBounds(unpaddedBounds, xpad, ypad, vspace)
	return float64(imageBounds.Dx()), float64(imageBounds.Dy()), err
}

func (r *rect) countLines(text []config.TextLine, conf *config.SignatureConfiguration) int {
	if len(text) == 0 && conf.Title != "" {
		return 1
	}
	countTitle := 1
	if conf.Title == "" {
		countTitle = 0
	} else if conf.EmptyLineAfterTitle {
		countTitle = 2
	}
	return len(text) + countTitle
}

func (r *rect) getPaddings(text []config.TextLine, conf *config.SignatureConfiguration) (xpad float64, ypad float64, vspace float64) {
	xpad = PtsToPixels(r.xpadpt, conf.Dpi)
	ypad = PtsToPixels(r.ypadpt, conf.Dpi)
	vspace = PtsToPixels(r.vspacept, conf.Dpi)
	if len(text) == 0 || r.allKeysAreEmpty(text) {
		vspace = 0
	}
	return
}

func (r *rect) allKeysAreEmpty(text []config.TextLine) bool {
	for _, line := range text {
		if line.Key != "" {
			return false
		}
	}
	return true
}

func (r *rect) getUnpaddedBounds(lines int, textwidth, textheight float64) image.Rectangle {
	return image.Rect(0, 0, int(math.Ceil(textwidth)), int(math.Ceil(textheight))*lines)
}

func (r *rect) getImageBounds(unpaddedBounds image.Rectangle, xpad, ypad, vspace float64) image.Rectangle {
	return image.Rect(0, 0, unpaddedBounds.Dx()+int(xpad*2)+int(vspace), unpaddedBounds.Dy()+int(ypad*2))
}

func (r *rect) calculateFontsAndUnpaddedSizesToFitPixels(text []config.TextLine, conf *config.SignatureConfiguration) (keyFace, valFace font.Face, fontSize, width, height float64, err error) {
	longestLine := config.TextLine{}
	longestLine.Key = r.longest(text, lineKey)
	longestLine.Value = r.longest(text, lineValue)
	if utf8.RuneCountInString(conf.Title) > utf8.RuneCountInString(longestLine.Key+longestLine.Value) {
		longestLine.Key = conf.Title
		longestLine.Value = ""
		return r.getFontsAndRealSizeToFitPixels(conf.TitleFont, conf.ValueFont, longestLine, conf.WidthPt, conf.Dpi)
	}
	return r.getFontsAndRealSizeToFitPixels(conf.KeyFont, conf.ValueFont, longestLine, conf.WidthPt, conf.Dpi)
}

func (r *rect) longest(texts []config.TextLine, f func(config.TextLine) string) (longest string) {
	var c, l int
	for _, text := range texts {
		s := f(text)
		c = utf8.RuneCountInString(s)
		if c > l {
			l = c
			longest = s
		}
	}
	return
}

func (r *rect) getFontsAndRealSizeToFitPixels(keyFont, valFont string, line config.TextLine, maxWidthPt float64, dpi float64) (keyFace, valFace font.Face, fontSize, width, height float64, err error) {
	dc := gg.NewContext(0, 0)
	min := 1.0
	max := 500.0 // reasonable upper limit
	var best float64 = min
	maxWidthPix := PtsToPixels(maxWidthPt, dpi)
	for range 20 { // binary search with 20 iterations
		mid := (min + max) / 2
		if keyFace, err = fonts.LoadFontFace(keyFont, dpi, mid); err != nil {
			return
		}
		if valFace, err = fonts.LoadFontFace(valFont, dpi, mid); err != nil {
			return
		}
		dc.SetFontFace(keyFace)
		kw, _ := dc.MeasureString(line.Key)
		dc.SetFontFace(valFace)
		vw, _ := dc.MeasureString(line.Value)
		if kw+vw > maxWidthPix {
			max = mid
		} else {
			best = mid
			min = mid
		}
	}
	fontSize = best
	if keyFace, err = fonts.LoadFontFace(keyFont, dpi, best); err != nil {
		return
	}
	if valFace, err = fonts.LoadFontFace(valFont, dpi, best); err != nil {
		return
	}
	dc.SetFontFace(keyFace)
	kw, kHeight := dc.MeasureString(line.Key)
	dc.SetFontFace(valFace)
	vw, vHeight := dc.MeasureString(line.Value)
	width = kw + vw
	height = math.Max(kHeight, vHeight)
	return
}

func (r *rect) RotateImage(img image.Image, conf *config.SignatureConfiguration) (image.Image, error) {
	var dc *gg.Context
	switch conf.Rotate {
	case config.ROTATE_0:
		return img, nil
	case config.ROTATE_90:
		dc = gg.NewContextForRGBA(image.NewRGBA(image.Rect(0, 0, img.Bounds().Dy(), img.Bounds().Dx())))
		dc.Translate(float64(img.Bounds().Dy()), 0)
		dc.Rotate(math.Pi / 2)
		dc.DrawImage(img, 0, 0)
	case config.ROTATE_180:
		dc = gg.NewContextForRGBA(image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy())))
		dc.RotateAbout(math.Pi, float64(img.Bounds().Dx())/2, float64(img.Bounds().Dy())/2)
		dc.DrawImage(img, 0, 0)
	case config.ROTATE_270:
		dc = gg.NewContextForRGBA(image.NewRGBA(image.Rect(0, 0, img.Bounds().Dy(), img.Bounds().Dx())))
		dc.Translate(0, float64(img.Bounds().Dx()))
		dc.Rotate(-math.Pi / 2)
		dc.DrawImage(img, 0, 0)
	default:
		return nil, eris.Errorf("invalid rotation %s", conf.Rotate)
	}
	return dc.Image(), nil
}
