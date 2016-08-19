// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package controller

import (
	"strconv"

	"github.com/tealeg/xlsx"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"
	"gopkg.in/goyy/goyy.v0/data/service"
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"gopkg.in/goyy/goyy.v0/util/uuids"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type JSONController struct {
	baseController
	pre
	post
	Settings
	Mgr service.Service
}

func (me *JSONController) Index(c xhttp.Context) {
	r, err := me.baseController.Index(c, me.Mgr, me.PreIndex, me.PostIndex)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Show(c xhttp.Context) {
	r, err := me.baseController.Show(c, me.Mgr, me.PreShow, me.PostShow)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Add(c xhttp.Context) {
	r, err := me.baseController.Add(c, me.Mgr, me.PreAdd, me.PostAdd)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Edit(c xhttp.Context) {
	r, err := me.baseController.Edit(c, me.Mgr, me.PreEdit, me.PostEdit)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Save(c xhttp.Context) {
	r, err := me.baseController.Save(c, me.Mgr, me.PreSave, me.PostSave)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Disable(c xhttp.Context) {
	r, err := me.baseController.Disable(c, me.Mgr, me.PreDisable, me.PostDisable)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.Text(xhttp.StatusOK, r.JSON())
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Box(c xhttp.Context) {
	out, err := me.baseController.Box(c, me.Mgr)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.JSON(xhttp.StatusOK, me.Success(c, out))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) Export(c xhttp.Context) {
	r, err := me.baseController.Export(c, me.Mgr, me.PreExport, me.PostExport)
	if err != nil {
		if err.Error() == i18N.Message("exp.limit") {
			err = c.JSON(xhttp.StatusOK, me.FaultMessage(c, err.Error()))
			if err != nil {
				me.Error(c, err)
				return
			}
			return
		}
		me.Error(c, err)
		return
	}
	f, err := me.excel(r)
	if err != nil {
		me.Error(c, err)
		return
	}
	err = c.JSON(xhttp.StatusOK, me.Success(c, f))
	if err != nil {
		me.Error(c, err)
		return
	}
}

func (me *JSONController) excel(r entity.Interfaces) (string, error) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return "", err
	}
	if r == nil {
		return "", errors.New(i18N.Message("exp.data.blank"))
	} else { // header
		e := r.New()
		row = sheet.AddRow()
		cw := 0
		for _, n := range e.ExcelColumns() {
			if t, ok := e.Type(n); ok {
				sheet.SetColWidth(cw, cw, float64(t.Field().Excel().Width()))
				cw++
				cell = row.AddCell()
				cell.Value = t.Field().Excel().Title()
				cell.SetStyle(me.excelHeaderStyle(t.Field().Excel().Align()))
			}
		}
	}
	for i := 0; i < r.Len(); i++ { // body
		e := r.Index(i)
		row = sheet.AddRow()
		for _, n := range e.ExcelColumns() {
			if t, ok := e.Type(n); ok {
				cell = row.AddCell()
				cell.SetStyle(me.excelBodyStyle(t.Field().Excel().Align()))
				val := t.String()
				cell.Value = val
				if strings.IsBlank(val) {
					continue
				}
				dict := t.Column().Dict()
				if strings.IsNotBlank(dict) {
					if schema.ParseDict != nil {
						value := schema.ParseDict(dict, val)
						cell.Value = value
						continue
					}
				}
				format := t.Field().Excel().Format()
				if strings.IsNotBlank(format) {
					value, err := strconv.ParseInt(val, 10, 64)
					if err != err {
						logger.Errorln(err.Error())
						continue
					}
					cell.Value = times.FormatUnix(format, value)
				}
			}
		}
	}
	filename := uuids.New() + ".xlsx"
	if !files.IsExist(xhttp.Conf.Export.Dir) {
		if err = files.MkdirAll(xhttp.Conf.Export.Dir, 0751); err != nil {
			return "", err
		}
	}
	err = file.Save(xhttp.Conf.Export.Dir + "/" + filename)
	return filename, err
}

func (me *JSONController) excelHeaderStyle(align int) *xlsx.Style {
	style := xlsx.NewStyle()
	fill := *xlsx.NewFill("solid", "00808080", "FF000000")
	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	style.Alignment.Horizontal = me.excelAlignment(align)
	style.Fill = fill
	style.Border = border
	style.ApplyAlignment = true
	style.ApplyFill = true
	style.ApplyBorder = true
	return style
}

func (me *JSONController) excelBodyStyle(align int) *xlsx.Style {
	style := xlsx.NewStyle()
	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	style.Alignment.Horizontal = me.excelAlignment(align)
	style.Border = border
	style.ApplyAlignment = true
	style.ApplyBorder = true
	return style
}

func (me *JSONController) excelAlignment(align int) string {
	a := "center"
	switch align {
	case 1:
		a = "left"
	case 3:
		a = "right"
	}
	return a
}
