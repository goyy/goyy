package assets

import (
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/controller"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

type Controller struct {
	controller.Controller
}

func (me *Controller) upload(c xhttp.Context) {
	d := c.Param("filedir")
	n := c.Param("fileinput")
	if strings.IsNotBlank(d) && strings.IsNotBlank(n) {
		if f, err := files.Upload(c.ResponseWriter(), c.Request(), n, xhttp.Conf.Upload.Dir, d); err == nil {
			c.Params().Set("img", f)
			c.Text(xhttp.StatusOK, `{"success":true,"filename":"`+f+`"}`)
			return
		}
	}
	c.Text(xhttp.StatusOK, `{"success":false}`)
}

func (me *Controller) ckupload(c xhttp.Context) {
	d := c.Param("filedir")
	n := c.Param("fileinput")
	CKEditorFuncNum := c.Param("CKEditorFuncNum")
	if strings.IsBlank(d) {
		d = "comm/assets"
	}
	if strings.IsBlank(n) {
		n = "upload"
	}
	if f, err := files.Upload(c.ResponseWriter(), c.Request(), n, xhttp.Conf.Upload.Dir, d); err == nil {
		p := xhttp.Conf.Upload.URL
		v := map[string]string{"CKEditorFuncNum": CKEditorFuncNum, "filepath": p + f}
		c.HTML(xhttp.StatusOK, "core/comm/ckeditor", v)
		return
	} else {
		v := map[string]string{"CKEditorFuncNum": CKEditorFuncNum, "filepath": i18N.Message("upload.img.err") + err.Error()}
		c.HTML(xhttp.StatusOK, "core/comm/ckeditor", v)
		return
	}
}

var ctl = &Controller{
	Controller: controller.Controller{
		Settings: controller.Settings{
			Project: "comm",
			Module:  "assets",
			Title:   "ASSETS",
		},
	},
}

func init() {
	xhttp.POST(ctl.PathBy("upload"), ctl.upload)
	xhttp.POST(ctl.PathBy("ckeditor/upload"), ctl.ckupload)
}
