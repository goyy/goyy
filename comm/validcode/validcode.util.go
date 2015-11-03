// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validcode

import (
	crand "crypto/rand"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
	"image/color"
	"io"
	"math/rand"
	"strconv"
)

func Judge(c xhttp.Context) bool {
	p := c.Param("validcode")
	if v, err := c.Session().Get("validcode"); err == nil && p == v {
		return true
	}
	return false
}

// New returns a new random string of the standard length, consisting of
// standard characters.
func New() (string, error) {
	return NewLenChars(StdLen, StdChars)
}

// NewLen returns a new random string of the provided length, consisting of
// standard characters.
func NewLen(length int) (string, error) {
	return NewLenChars(length, StdChars)
}

// NewLenChars returns a new random string of the provided length, consisting
// of the provided byte slice of allowed characters (maximum 256).
func NewLenChars(length int, chars []byte) (string, error) {
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(crand.Reader, r); err != nil {
			return "", err
		}
		for _, c := range r {
			if c >= maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b), nil
			}
		}
	}
	return "", errors.New("unreachable")
}

func setRandomBrightness(c *color.NRGBA, max uint8) {
	minc := min3(c.R, c.G, c.B)
	maxc := max3(c.R, c.G, c.B)
	if maxc > max {
		return
	}
	n := rand.Intn(int(max-maxc)) - int(minc)
	c.R = uint8(int(c.R) + n)
	c.G = uint8(int(c.G) + n)
	c.B = uint8(int(c.B) + n)
}

func min3(x, y, z uint8) (o uint8) {
	o = x
	if y < o {
		o = y
	}
	if z < o {
		o = z
	}
	return
}

func max3(x, y, z uint8) (o uint8) {
	o = x
	if y > o {
		o = y
	}
	if z > o {
		o = z
	}
	return
}

// rnd returns a random number in range [from, to].
func rnd(from, to int) int {
	//println(to+1-from)
	return rand.Intn(to+1-from) + from
}

func createImage(c xhttp.Context) {
	d := make([]byte, 4)
	s, err := NewLen(4)
	if err != nil {
		c.Error(500)
		return
	}
	validcode := ""
	d = []byte(s)
	for v := range d {
		d[v] %= 10
		validcode += strconv.FormatInt(int64(d[v]), 32)
	}
	c.ResponseWriter().Header().Set("Content-Type", "image/png")
	c.Session().Set("validcode", validcode)
	NewImage(d, 90, 40).WriteTo(c.ResponseWriter())
}
