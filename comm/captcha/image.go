// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package captcha

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
)

const (
	// Standard width and height of a captcha image.
	StdWidth  = 140
	StdHeight = 80
	// Maximum absolute skew factor of a single digit.
	maxSkew = 0.7
	// Number of background circles.
	circleCount = 20
)

type Image struct {
	*image.Paletted
	numWidth  int
	numHeight int
	dotSize   int
	rng       siprng
}

// NewImage returns a new captcha image of the given width and height with the
// given digits, where each digit must be in range 0-9.
func NewImage(id string, digits []byte, width, height int) *Image {
	m := new(Image)

	// Initialize PRNG.
	m.rng.Seed(deriveSeed(imageSeedPurpose, id, digits))

	m.Paletted = image.NewPaletted(image.Rect(0, 0, width, height), m.getRandomPalette())
	m.calculateSizes(width, height, len(digits))
	// Randomly position captcha inside the image.
	maxx := width - (m.numWidth+m.dotSize)*len(digits) - m.dotSize
	maxy := height - m.numHeight - m.dotSize*2
	var border int
	if width > height {
		border = height / 5
	} else {
		border = width / 5
	}
	x := m.rng.Int(border, maxx-border)
	y := m.rng.Int(border, maxy-border)
	// Draw digits.
	for _, n := range digits {
		m.drawDigit(font[n], x, y)
		x += m.numWidth + m.dotSize
	}
	// Draw strike-through line.
	m.strikeThrough()
	// Apply wave distortion.
	m.distort(m.rng.Float(5, 10), m.rng.Float(100, 200))
	// Fill image with random circles.
	m.fillWithCircles(circleCount, m.dotSize)
	return m
}

func (me *Image) getRandomPalette() color.Palette {
	p := make([]color.Color, circleCount+1)
	// Transparent color.
	p[0] = color.RGBA{0xFF, 0xFF, 0xFF, 0x00}
	// Primary color.
	prim := color.RGBA{
		uint8(me.rng.Intn(129)),
		uint8(me.rng.Intn(129)),
		uint8(me.rng.Intn(129)),
		0xFF,
	}
	p[1] = prim
	// Circle colors.
	for i := 2; i <= circleCount; i++ {
		p[i] = me.randomBrightness(prim, 255)
	}
	return p
}

// encodedPNG encodes an image to PNG and returns
// the result as a byte slice.
func (me *Image) encodedPNG() []byte {
	var buf bytes.Buffer
	if err := png.Encode(&buf, me.Paletted); err != nil {
		panic(err.Error())
	}
	return buf.Bytes()
}

// WriteTo writes captcha image in PNG format into the given writer.
func (me *Image) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(me.encodedPNG())
	return int64(n), err
}

func (me *Image) calculateSizes(width, height, ncount int) {
	// Goal: fit all digits inside the image.
	var border int
	if width > height {
		border = height / 4
	} else {
		border = width / 4
	}
	// Convert everything to floats for calculations.
	w := float64(width - border*2)
	h := float64(height - border*2)
	// fw takes into account 1-dot spacing between digits.
	fw := float64(fontWidth + 1)
	fh := float64(fontHeight)
	nc := float64(ncount)
	// Calculate the width of a single digit taking into account only the
	// width of the image.
	nw := w / nc
	// Calculate the height of a digit from this width.
	nh := nw * fh / fw
	// Digit too high?
	if nh > h {
		// Fit digits based on height.
		nh = h
		nw = fw / fh * nh
	}
	// Calculate dot size.
	me.dotSize = int(nh / fh)
	if me.dotSize < 1 {
		me.dotSize = 1
	}
	// Save everything, making the actual width smaller by 1 dot to account
	// for spacing between digits.
	me.numWidth = int(nw) - me.dotSize
	me.numHeight = int(nh)
}

func (me *Image) drawHorizLine(fromX, toX, y int, colorIdx uint8) {
	for x := fromX; x <= toX; x++ {
		me.SetColorIndex(x, y, colorIdx)
	}
}

func (me *Image) drawCircle(x, y, radius int, colorIdx uint8) {
	f := 1 - radius
	dfx := 1
	dfy := -2 * radius
	xo := 0
	yo := radius

	me.SetColorIndex(x, y+radius, colorIdx)
	me.SetColorIndex(x, y-radius, colorIdx)
	me.drawHorizLine(x-radius, x+radius, y, colorIdx)

	for xo < yo {
		if f >= 0 {
			yo--
			dfy += 2
			f += dfy
		}
		xo++
		dfx += 2
		f += dfx
		me.drawHorizLine(x-xo, x+xo, y+yo, colorIdx)
		me.drawHorizLine(x-xo, x+xo, y-yo, colorIdx)
		me.drawHorizLine(x-yo, x+yo, y+xo, colorIdx)
		me.drawHorizLine(x-yo, x+yo, y-xo, colorIdx)
	}
}

func (me *Image) fillWithCircles(n, maxradius int) {
	maxx := me.Bounds().Max.X
	maxy := me.Bounds().Max.Y
	for i := 0; i < n; i++ {
		colorIdx := uint8(me.rng.Int(1, circleCount-1))
		r := me.rng.Int(1, maxradius)
		me.drawCircle(me.rng.Int(r, maxx-r), me.rng.Int(r, maxy-r), r, colorIdx)
	}
}

func (me *Image) strikeThrough() {
	maxx := me.Bounds().Max.X
	maxy := me.Bounds().Max.Y
	y := me.rng.Int(maxy/3, maxy-maxy/3)
	amplitude := me.rng.Float(5, 20)
	period := me.rng.Float(80, 180)
	dx := 2.0 * math.Pi / period
	for x := 0; x < maxx; x++ {
		xo := amplitude * math.Cos(float64(y)*dx)
		yo := amplitude * math.Sin(float64(x)*dx)
		for yn := 0; yn < me.dotSize; yn++ {
			r := me.rng.Int(0, me.dotSize)
			me.drawCircle(x+int(xo), y+int(yo)+(yn*me.dotSize), r/2, 1)
		}
	}
}

func (me *Image) drawDigit(digit []byte, x, y int) {
	skf := me.rng.Float(-maxSkew, maxSkew)
	xs := float64(x)
	r := me.dotSize / 2
	y += me.rng.Int(-r, r)
	for yo := 0; yo < fontHeight; yo++ {
		for xo := 0; xo < fontWidth; xo++ {
			if digit[yo*fontWidth+xo] != blackChar {
				continue
			}
			me.drawCircle(x+xo*me.dotSize, y+yo*me.dotSize, r, 1)
		}
		xs += skf
		x = int(xs)
	}
}

func (me *Image) distort(amplude float64, period float64) {
	w := me.Bounds().Max.X
	h := me.Bounds().Max.Y

	oldm := me.Paletted
	newm := image.NewPaletted(image.Rect(0, 0, w, h), oldm.Palette)

	dx := 2.0 * math.Pi / period
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			xo := amplude * math.Sin(float64(y)*dx)
			yo := amplude * math.Cos(float64(x)*dx)
			newm.SetColorIndex(x, y, oldm.ColorIndexAt(x+int(xo), y+int(yo)))
		}
	}
	me.Paletted = newm
}

func (me *Image) randomBrightness(c color.RGBA, max uint8) color.RGBA {
	minc := min3(c.R, c.G, c.B)
	maxc := max3(c.R, c.G, c.B)
	if maxc > max {
		return c
	}
	n := me.rng.Intn(int(max-maxc)) - int(minc)
	return color.RGBA{
		uint8(int(c.R) + n),
		uint8(int(c.G) + n),
		uint8(int(c.B) + n),
		uint8(c.A),
	}
}

func min3(x, y, z uint8) (m uint8) {
	m = x
	if y < m {
		m = y
	}
	if z < m {
		m = z
	}
	return
}

func max3(x, y, z uint8) (m uint8) {
	m = x
	if y > m {
		m = y
	}
	if z > m {
		m = z
	}
	return
}
