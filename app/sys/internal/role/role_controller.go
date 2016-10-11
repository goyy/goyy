package role

import (
	"gopkg.in/goyy/goyy.v0/app/sys/api/dict"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(ctl.ApiIndex(), ctl.Index, ctl.PermitView())
	xhttp.POST(ctl.ApiIndex(), ctl.Index, ctl.PermitView())
	xhttp.GET(ctl.ApiShow(), ctl.Show, ctl.PermitView())
	xhttp.POST(ctl.ApiShow(), ctl.Show, ctl.PermitView())
	xhttp.POST(ctl.ApiAdd(), ctl.Add, ctl.PermitAdd())
	xhttp.POST(ctl.ApiEdit(), ctl.Edit, ctl.PermitEdit())
	xhttp.POST(ctl.ApiSave(), ctl.Save, ctl.PermitAdd(), ctl.PermitEdit())
	xhttp.POST(ctl.ApiDisable(), ctl.DisableAndTx, ctl.PermitDisable())
	xhttp.GET(ctl.ApiBox(), ctl.Box)
	xhttp.GET(ctl.ApiBy("data"), ctl.data)

	ctl.PreAdd = preSetPostIds
	ctl.PreEdit = preSetPostIds
	ctl.PreShow = preSetPostIds
	ctl.PostSave = refresh
	ctl.PostDisable = refresh
}

// Gets a list of roles for the current logged in user (data permissions).
func (me *Controller) data(c xhttp.Context) {
	p, err := c.Session().Principal()
	if err != nil {
		c.JSON(xhttp.StatusOK, me.FaultMessage(c, i18N.Message("not.logged.err")))
		return
	}
	box := make([]xtype.Box, 0, 10)
	data := p.Roles.Data
	if strings.IsBlank(data) {
		c.JSON(xhttp.StatusOK, me.Success(c, box))
		return
	}
	classify := c.Param("classify")
	datas := strings.Split(data, ",")
	for _, d := range datas {
		if v, ok := dict.Get("sys_role.id", d); ok {
			if strings.IsNotBlank(classify) && classify != v.Filters {
				continue
			}
			box = append(box, xtype.Box{
				Id:   d,
				Name: v.Mval,
			})
		}
	}
	c.JSON(xhttp.StatusOK, me.Success(c, box))
}

func preSetPostIds(c xhttp.Context) error {
	if postIds, err := posts(c.Param("sIdTR")); err != nil {
		return err
	} else {
		c.Params().Set("postIds", postIds)
		return nil
	}
}

// Refresh the cache data for the role dictionary.
func refresh(c xhttp.Context, r *result.Entity) error {
	refreshCache()
	return nil
}
