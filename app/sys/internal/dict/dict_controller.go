package dict

import (
	"encoding/json"

	"gopkg.in/goyy/goyy.v0/app/sys/internal/cache/dict"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/domain"
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
	xhttp.POST(ctl.ApiSave(), ctl.SaveAndTx, ctl.PermitAdd(), ctl.PermitEdit())
	xhttp.POST(ctl.ApiDisable(), ctl.DisableAndTx, ctl.PermitDisable())
	xhttp.GET(ctl.ApiBox(), ctl.Box)
	xhttp.GET(ctl.ApiBy("mval"), ctl.mval)
	xhttp.GET(ctl.ApiBy("mvals"), ctl.mvals)
	xhttp.GET(ctl.ApiBy("genres"), ctl.genres)
	xhttp.POST(ctl.ApiBy("list"), ctl.list)
	ctl.PostSave = refresh
	ctl.PostDisable = refresh
}

// Update dictionary cache.
func refresh(c xhttp.Context, r *result.Entity) error {
	refreshCache()
	return nil
}

// Through dictionary type and key acquisition value.
func (me *Controller) mval(c xhttp.Context) {
	val := dict.Mval(c.Param("genre"), c.Param("mkey"))
	c.JSON(xhttp.StatusOK, me.Success(c, val))
}

// Get a list of values from the dictionary type and the list of keys
// (comma separated).
func (me *Controller) mvals(c xhttp.Context) {
	val := dict.Mvals(c.Param("genre"), c.Param("mkeys"))
	c.JSON(xhttp.StatusOK, me.Success(c, val))
}

// Get a list of all dictionary types.
func (me *Controller) genres(c xhttp.Context) {
	dicts, err := dict.ListByGenre()
	if err != nil {
		me.Error(c, err)
		return
	}
	boxes := make([]xtype.Box, 0, len(dicts))
	for _, dict := range dicts {
		box := xtype.Box{}
		box.Id = dict
		box.Name = dict
		boxes = append(boxes, box)
	}
	err = c.JSON(xhttp.StatusOK, me.Success(c, boxes))
	if err != nil {
		me.Error(c, err)
		return
	}
}

// Obtain a list of dictionaries by dictionary type and key list.
func (me *Controller) list(c xhttp.Context) {
	p := c.Param("params")
	if strings.IsNotBlank(p) {
		var dicts []*xtype.Dict
		if err := json.Unmarshal([]byte(p), &dicts); err == nil {
			c.JSON(xhttp.StatusOK, me.Success(c, dict.List(dicts)))
			return
		} else {
			logger.Error(err)
		}
	}
	c.JSON(xhttp.StatusOK, me.Success(c, ""))
}

func (me *Controller) Box(c xhttp.Context) {
	sifts, err := domain.NewSifts(c.Params())
	var genre string
	var filters string
	for _, v := range sifts {
		if v.Key() == "Genre" {
			genre = v.Value()
		}
		if v.Key() == "Filters" {
			filters = v.Value()
		}
	}
	out := make([]xtype.Box, 0)
	if strings.IsNotBlank(genre) {
		var boxId, boxName, boxTime string
		var isSysDict bool
		dicts := dict.GetCache(cacheField)
		for _, v := range dicts { // genre from the sys_dict table
			if v.Genre == genre {
				isSysDict = true
				if strings.IsNotBlank(filters) && v.Filters != filters {
					continue
				}
				boxId = v.Mkey
				boxName = v.Mval
				boxTime = v.Created
				out = append(out, xtype.Box{
					Id:   boxId,
					Name: boxName,
					Time: boxTime,
				})
			}
		}
		if !isSysDict { // genre from non sys_dict table：For example, from the sys_role table
			if dict.HasCache(genre) {
				dicts = dict.GetCache(genre)
				for _, v := range dicts {
					if v.Genre == genre {
						if strings.IsNotBlank(filters) && v.Filters != filters {
							continue
						}
						boxId = v.Mkey
						boxName = v.Mval
						boxTime = v.Created
						out = append(out, xtype.Box{
							Id:   boxId,
							Name: boxName,
							Time: boxTime,
						})
					}
				}
			}
		}
	}
	err = c.JSON(xhttp.StatusOK, me.Success(c, out))
	if err != nil {
		me.Error(c, err)
		return
	}
}
