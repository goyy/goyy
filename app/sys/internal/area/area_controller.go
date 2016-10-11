package area

import (
	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/data/result"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.GET(ctl.ApiIndex(), ctl.Index, ctl.PermitView(profile.BMS))
	xhttp.POST(ctl.ApiIndex(), ctl.Index, ctl.PermitView(profile.BMS))
	xhttp.GET(ctl.ApiShow(), ctl.Show, ctl.PermitView(profile.BMS))
	xhttp.POST(ctl.ApiShow(), ctl.Show, ctl.PermitView(profile.BMS))
	xhttp.POST(ctl.ApiSave(), ctl.SaveAndTx, ctl.PermitAdd(profile.BMS), ctl.PermitEdit(profile.BMS))
	xhttp.GET(ctl.ApiTree(), ctl.Tree)

	xhttp.POST(ctl.ApiDisable(), ctl.DisableAndTx, ctl.PermitDisable())
	xhttp.POST(ctl.ApiAdd(), ctl.Add, ctl.PermitAdd())
	xhttp.POST(ctl.ApiEdit(), ctl.Edit, ctl.PermitEdit())

	xhttp.GET(ctl.ApiBy("name"), ctl.name)
	xhttp.GET(ctl.ApiBy("fullname"), ctl.fullname)
	xhttp.GET(ctl.ApiBy("parentid"), ctl.parentid)
	xhttp.GET(ctl.ApiBy("parentpid"), ctl.parentpid)
	xhttp.GET(ctl.ApiBy("parentpname"), ctl.parentpname)
	ctl.PostSave = refresh
	ctl.PostDisable = refresh
}

// Get a name based on the identity.
func (me *Controller) name(c xhttp.Context) {
	id := c.Param("id")
	if strings.IsNotBlank(id) {
		a := NewEntity()
		a.SetId(id)
		err := Mgr.Get(a)
		if err != nil {
			logger.Error(err.Error())
			me.Error(c, err)
			return
		}
		c.JSON(xhttp.StatusOK, me.Success(c, a.Name()))
		return
	}
	c.JSON(xhttp.StatusOK, me.Fault(c, i18N.Message("get.name.err")))
	return
}

// Get full name based on identity.
func (me *Controller) fullname(c xhttp.Context) {
	id := c.Param("id")
	if strings.IsNotBlank(id) {
		a := NewEntity()
		a.SetId(id)
		err := Mgr.Get(a)
		if err != nil {
			logger.Error(err.Error())
			me.Error(c, err)
			return
		}
		c.JSON(xhttp.StatusOK, me.Success(c, a.Fullname()))
		return
	}
	c.JSON(xhttp.StatusOK, me.Fault(c, i18N.Message("get.fullname.err")))
	return
}

// Gets the superior ID based on the identity.
func (me *Controller) parentid(c xhttp.Context) {
	id := c.Param("id")
	if strings.IsNotBlank(id) {
		a := NewEntity()
		a.SetId(id)
		err := Mgr.Get(a)
		if err != nil {
			logger.Error(err.Error())
			me.Error(c, err)
			return
		}
		c.JSON(xhttp.StatusOK, me.Success(c, a.ParentId()))
		return
	}
	c.JSON(xhttp.StatusOK, me.Fault(c, i18N.Message("get.parentid.err")))
	return
}

// Gets the superior ID of the superior based on the identity.
func (me *Controller) parentpid(c xhttp.Context) {
	id := c.Param("id")
	if strings.IsNotBlank(id) {
		// Query city based on identity
		a := NewEntity()
		a.SetId(id)
		err := Mgr.Get(a)
		if err != nil {
			logger.Error(err.Error())
			me.Error(c, err)
			return
		}
		// Query city according to the parent ID
		p := NewEntity()
		p.SetId(a.ParentId())
		err = Mgr.Get(p)
		if err != nil {
			logger.Error(err.Error())
			me.Error(c, err)
			return
		}
		c.JSON(xhttp.StatusOK, me.Success(c, p.ParentId()))
		return
	}
	c.JSON(xhttp.StatusOK, me.Fault(c, i18N.Message("get.parentpid.err")))
	return
}

// Gets the superior name of the superior based on the identity.
func (me *Controller) parentpname(c xhttp.Context) {
	id := c.Param("id")
	if strings.IsNotBlank(id) {
		// Query city based on identity
		a := NewEntity()
		a.SetId(id)
		err := Mgr.Get(a)
		if err != nil {
			logger.Error(err.Error())
			me.Error(c, err)
			return
		}
		// Query city according to the parent ID
		fullname := strings.Split(a.Fullname(), " - ")
		if len(fullname) >= 1 {
			c.JSON(xhttp.StatusOK, me.Success(c, fullname[0]))
			return
		} else {
			c.JSON(xhttp.StatusOK, me.Fault(c, i18N.Message("get.parentpname.err")))
			return
		}
	}
	c.JSON(xhttp.StatusOK, me.Fault(c, i18N.Message("get.parentpname.err")))
	return
}

// Refresh the cache data for the area dictionary.
func refresh(c xhttp.Context, r *result.Entity) error {
	refreshCache()
	return nil
}
