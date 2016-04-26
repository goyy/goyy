// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package captcha

import (
	"bytes"
	"strconv"

	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/controller"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

const (
	paramkey   = "captcha"
	sessionkey = "gs-captcha-key"
)

type Controller struct {
	ImgWidth  int
	ImgHeight int
	ImgLength int
	controller.JSONController
}

// Build captcha
func (me *Controller) Build(c xhttp.Context) {
	c.ResponseWriter().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.ResponseWriter().Header().Set("Pragma", "no-cache")
	c.ResponseWriter().Header().Set("Expires", "0")
	c.ResponseWriter().Header().Set("Content-Type", "image/png")
	if me.ImgWidth <= 0 {
		me.ImgWidth = StdWidth
	}
	if me.ImgHeight <= 0 {
		me.ImgHeight = StdHeight
	}
	if me.ImgLength <= 0 {
		me.ImgLength = DefaultLen
	}
	if v, err := WriteImage(c.ResponseWriter(), sessionkey, me.ImgLength, me.ImgWidth, me.ImgHeight); err == nil {
		var b bytes.Buffer
		for _, val := range v {
			b.WriteString(strconv.Itoa(int(val)))
		}
		logger.Println("{captcha:" + b.String() + "}")
		c.Session().Set(sessionkey, b.String())
	} else {
		logger.Error(err.Error())
		me.Error(c, err)
		return
	}
}

// Ajax verify captcha
func (me *Controller) Verify(c xhttp.Context) {
	if Verify(c, false) {
		c.JSON(xhttp.StatusOK, me.SuccessMessage(c, "Verification code matching correctly"))
		return
	}
	c.JSON(xhttp.StatusOK, me.FaultMessage(c, "Verification code matching error"))
	return
}

// verify captcha
func Verify(c xhttp.Context, clear bool) bool {
	if c == nil {
		return false
	}
	v := c.Param(paramkey)
	if strings.IsBlank(v) {
		return false
	}
	val, err := c.Session().Get(sessionkey)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	if strings.IsBlank(val) {
		return false
	}
	if v == val {
		if clear {
			if err = c.Session().Delete(sessionkey); err != nil {
				logger.Error(err.Error())
			}
		}
		return true
	}
	return false
}
