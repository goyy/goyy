// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validcode

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/rand"
)

func NewImage(digits []byte, width, height int) *Image {
	img := new(Image)
	r := image.Rect(img.width, img.height, stdWidth, stdHeight)
	img.NRGBA = image.NewNRGBA(r)

	img.color = &color.NRGBA{
		uint8(rand.Intn(129)),
		uint8(rand.Intn(129)),
		uint8(rand.Intn(129)),
		0xFF,
	}
	// Draw background (10 random circles of random brightness)
	img.calculateSizes(width, height, len(digits))
	img.fillWithCircles(10, img.dotsize)

	maxx := width - (img.width+img.dotsize)*len(digits) - img.dotsize
	maxy := height - img.height - img.dotsize*2

	x := rnd(img.dotsize*2, maxx)
	y := rnd(img.dotsize*2, maxy)

	// Draw digits.
	for _, n := range digits {
		img.drawDigit(font[n], x, y)
		x += img.width + img.dotsize
	}

	// Draw strike-through line.
	img.strikeThrough()
	return img
}

type Image struct {
	*image.NRGBA
	color   *color.NRGBA
	width   int //a digit width
	height  int //a digit height
	dotsize int
}

func (me *Image) WriteTo(w io.Writer) (int64, error) {
	return 0, png.Encode(w, me)
}

func (me *Image) calculateSizes(width, height, ncount int) {

	// Goal: fit all digits inside the image.
	var border int
	if width > height {
		border = height / 5
	} else {
		border = width / 5
	}
	// Convert everything to floats for calculations.
	w := float64(width - border*2)  //268
	h := float64(height - border*2) //48
	// fw takes into account 1-dot spacing between digits.

	fw := float64(fontWidth) + 1 //6

	fh := float64(fontHeight) //8
	nc := float64(ncount)     //7

	// Calculate the width of a single digit taking into account only the
	// width of the image.
	nw := w / nc //38
	// Calculate the height of a digit from this width.
	nh := nw * fh / fw //51

	// Digit too high?

	if nh > h {
		// Fit digits based on height.
		nh = h //nh = 44
		nw = fw / fh * nh
	}
	// Calculate dot size.
	me.dotsize = int(nh / fh)
	// Save everything, making the actual width smaller by 1 dot to account
	// for spacing between digits.
	me.width = int(nw)
	me.height = int(nh) - me.dotsize
}

func (me *Image) fillWithCircles(n, maxradius int) {
	color := me.color
	maxx := me.Bounds().Max.X
	maxy := me.Bounds().Max.Y
	for i := 0; i < n; i++ {
		setRandomBrightness(color, 255)
		r := rnd(1, maxradius)
		me.drawCircle(color, rnd(r, maxx-r), rnd(r, maxy-r), r)
	}
}

func (me *Image) drawHorizLine(color color.Color, fromX, toX, y int) {
	for x := fromX; x <= toX; x++ {
		me.Set(x, y, color)
	}
}

func (me *Image) drawCircle(color color.Color, x, y, radius int) {
	f := 1 - radius
	dfx := 1
	dfy := -22 * radius
	xx := 0
	yy := radius

	me.Set(x, y+radius, color)
	me.Set(x, y-radius, color)
	me.drawHorizLine(color, x-radius, x+radius, y)

	for xx < yy {
		if f >= 0 {
			yy--
			dfy += 2
			f += dfy
		}
		xx++
		dfx += 2
		f += dfx
		me.drawHorizLine(color, x-xx, x+xx, y+yy)
		me.drawHorizLine(color, x-xx, x+xx, y-yy)
		me.drawHorizLine(color, x-yy, x+yy, y+xx)
		me.drawHorizLine(color, x-yy, x+yy, y-xx)
	}
}

func (me *Image) strikeThrough() {
	r := 0
	maxx := me.Bounds().Max.X
	maxy := me.Bounds().Max.Y
	y := rnd(maxy/3, maxy-maxy/3)
	for x := 0; x < maxx; x += r {
		r = rnd(1, me.dotsize/3)
		y += rnd(-me.dotsize/2, me.dotsize/2)
		if y <= 0 || y >= maxy {
			y = rnd(maxy/3, maxy-maxy/3)
		}
		me.drawCircle(me.color, x, y, r)
	}
}

func (me *Image) drawDigit(digit []byte, x, y int) {
	skf := rand.Float64() * float64(rnd(-maxSkew, maxSkew))
	xs := float64(x)
	minr := me.dotsize / 2              // minumum radius
	maxr := me.dotsize/2 + me.dotsize/4 // maximum radius
	y += rnd(-minr, minr)
	for yy := 0; yy < fontHeight; yy++ {
		for xx := 0; xx < fontWidth; xx++ {
			if digit[yy*fontWidth+xx] != blackChar {
				continue
			}
			// Introduce random variations.
			or := rnd(minr, maxr)
			ox := x + (xx * me.dotsize) + rnd(0, or/2)
			oy := y + (yy * me.dotsize) + rnd(0, or/2)

			me.drawCircle(me.color, ox, oy, or)
		}
		xs += skf
		x = int(xs)
	}
}
