package public

import (
	"crypto/rand"
	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
	"log"
	"math"
	"math/big"
)

const (
	defaultDpi      = 100
	defaultFontSize = 25
)

type Captcha struct {
	W, H, CodeLen int
	FontSize      float64
	Dpi           int
}

// 实例化验证码
func NewCaptcha(w, h, CodeLen int) *Captcha {
	return &Captcha{W: w, H: h, CodeLen: CodeLen}
}

// 输出
func (captcha *Captcha) OutPut() (string, *image.RGBA) {
	img := captcha.initCanvas()
	return captcha.doImage(img)
}

// 初始化画布
func (captcha *Captcha) initCanvas() *image.RGBA {
	dest := image.NewRGBA(image.Rect(0, 0, captcha.W, captcha.H))

	// 随机色
	r := uint8(255) // uint8(captcha.RangeRand(50, 250))
	g := uint8(255) // uint8(captcha.RangeRand(50, 250))
	b := uint8(255) // uint8(captcha.RangeRand(50, 250))

	// 填充背景色
	for x := 0; x < captcha.W; x++ {
		for y := 0; y < captcha.H; y++ {
			dest.Set(x, y, color.RGBA{r, g, b, 255}) // 设定alpha图片的透明度
		}
	}

	return dest
}

// 处理图像
func (captcha *Captcha) doImage(dest *image.RGBA) (string, *image.RGBA) {
	gc := draw2dimg.NewGraphicContext(dest)

	defer gc.Close()
	defer gc.FillStroke()

	captcha.setFont(gc)
	captcha.doPoint(gc)
	captcha.doLine(gc)
	captcha.doSinLine(gc)

	var codeStr string

	codeStr = RandString(captcha.CodeLen, CAPITAL, LOWER_CASE, NUMBER)
	captcha.doCode(gc, codeStr)

	return codeStr, dest
}

// 增加干扰线
func (captcha *Captcha) doLine(gc *draw2dimg.GraphicContext) {
	// 设置干扰线
	for n := 0; n < 5; n++ {
		// gc.SetLineWidth(float64(captcha.RangeRand(1, 2)))
		gc.SetLineWidth(1)

		// 随机背景色
		r := uint8(captcha.RangeRand(0, 255))
		g := uint8(captcha.RangeRand(0, 255))
		b := uint8(captcha.RangeRand(0, 255))

		gc.SetStrokeColor(color.RGBA{r, g, b, 255})

		// 初始化位置
		gc.MoveTo(float64(captcha.RangeRand(0, int64(captcha.W)+10)), float64(captcha.RangeRand(0, int64(captcha.H)+5)))
		gc.LineTo(float64(captcha.RangeRand(0, int64(captcha.W)+10)), float64(captcha.RangeRand(0, int64(captcha.H)+5)))

		gc.Stroke()
	}
}

// 增加干扰点
func (captcha *Captcha) doPoint(gc *draw2dimg.GraphicContext) {
	for n := 0; n < 50; n++ {
		gc.SetLineWidth(float64(captcha.RangeRand(1, 3)))

		// 随机色
		r := uint8(captcha.RangeRand(0, 255))
		g := uint8(captcha.RangeRand(0, 255))
		b := uint8(captcha.RangeRand(0, 255))

		gc.SetStrokeColor(color.RGBA{r, g, b, 255})

		x := captcha.RangeRand(0, int64(captcha.W)+10) + 1
		y := captcha.RangeRand(0, int64(captcha.H)+5) + 1

		gc.MoveTo(float64(x), float64(y))
		gc.LineTo(float64(x+captcha.RangeRand(1, 2)), float64(y+captcha.RangeRand(1, 2)))

		gc.Stroke()
	}
}

// 增加正弦干扰线
func (captcha *Captcha) doSinLine(gc *draw2dimg.GraphicContext) {
	h1 := captcha.RangeRand(-12, 12)
	h2 := captcha.RangeRand(-1, 1)
	w2 := captcha.RangeRand(5, 20)
	h3 := captcha.RangeRand(5, 10)

	h := float64(captcha.H)
	w := float64(captcha.W)

	// 随机色
	r := uint8(captcha.RangeRand(128, 255))
	g := uint8(captcha.RangeRand(128, 255))
	b := uint8(captcha.RangeRand(128, 255))

	gc.SetStrokeColor(color.RGBA{r, g, b, 255})
	gc.SetLineWidth(float64(captcha.RangeRand(2, 4)))

	var i float64
	for i = -w / 2; i < w/2; i = i + 0.1 {
		y := h/float64(h3)*math.Sin(i/float64(w2)) + h/2 + float64(h1)

		gc.LineTo(i+w/2, y)

		if h2 == 0 {
			gc.LineTo(i+w/2, y+float64(h2))
		}
	}

	gc.Stroke()
}

// 验证码字符设置到图像上
func (captcha *Captcha) doCode(gc *draw2dimg.GraphicContext, code string) {
	for l := 0; l < len(code); l++ {
		y := captcha.RangeRand(int64(captcha.FontSize)-1, int64(captcha.H)+6)
		x := captcha.RangeRand(1, 20)

		// 随机色
		r := uint8(captcha.RangeRand(0, 200))
		g := uint8(captcha.RangeRand(0, 200))
		b := uint8(captcha.RangeRand(0, 200))

		gc.SetFillColor(color.RGBA{r, g, b, 255})
		gc.FillStringAt(string(code[l]), float64(x)+captcha.FontSize*float64(l), float64(int64(captcha.H)-y)+captcha.FontSize)
		gc.Stroke()
	}
}

// 设置相关字体
func (captcha *Captcha) setFont(gc *draw2dimg.GraphicContext) {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Println(err)
		return
	}
	gc.FontCache = draw2d.NewSyncFolderFontCache("")
	gc.FontCache.Store(draw2d.FontData{Name: "t", Family: 0, Style: draw2d.FontStyleNormal}, font)
	gc.SetFontData(draw2d.FontData{Name: "t", Style: draw2d.FontStyleNormal})

	// 设置清晰度
	if captcha.Dpi <= 0 {
		captcha.Dpi = defaultDpi
	}
	gc.SetDPI(captcha.Dpi)

	// 设置字体大小
	if captcha.FontSize <= 0 {
		captcha.FontSize = defaultFontSize
	}
	gc.SetFontSize(captcha.FontSize)
}

// 获取区间[-m, n]的随机数
func (captcha *Captcha) RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))

		return min + result.Int64()
	}
}
