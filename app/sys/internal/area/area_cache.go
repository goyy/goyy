package area

import (
	"gopkg.in/goyy/goyy.v0/app/sys/api/dict"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/util/times"
)

const cacheField = "sys_area.id"

// Refresh the cache data for the area dictionary.
func refreshCache() {
	dict.RefreshCache(cacheField)
}

// Initialize the cache data for the area dictionary.
func initCache() []*xtype.Dict {
	sDeletionEQ, _ := domain.NewSift("sDeletionEQ", "0")
	sOrdinalOA, _ := domain.NewSift("sOrdinalOA", "20")
	es := NewEntities(5000)
	if err := Mgr.SelectListBySift(es, sDeletionEQ, sOrdinalOA); err == nil {
		dicts := make([]*xtype.Dict, es.Len())
		for i, d := range es.Values() {
			e := &xtype.Dict{
				Mkey:    d.Id(),
				Mval:    d.Name(),
				Genre:   cacheField,
				Descr:   d.Fullname(),
				Ordinal: d.Ordinal(),
				Created: times.FormatUnixYYMDHMS(d.Created()),
			}
			dicts[i] = e
		}
		return dicts
	} else {
		logger.Errorln(err)
		return nil
	}
}

// Refresh the cache data for the area dictionary.
func init() {
	dict.RegisterCache(cacheField, initCache)
}
