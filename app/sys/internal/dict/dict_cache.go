package dict

import (
	"gopkg.in/goyy/goyy.v0/app/sys/api/dict"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/util/times"
)

const cacheField = "dict"

// Refresh the cache data for the dictionary.
func refreshCache() {
	dict.RefreshCache(cacheField)
}

// Initialize the cache data for the dictionary.
func initCache() []*xtype.Dict {
	sDeletionEQ, _ := domain.NewSift("sDeletionEQ", "0")
	sGenreOA, _ := domain.NewSift("sGenreOA", "10")
	sOrdinalOA, _ := domain.NewSift("sOrdinalOA", "20")
	es := NewEntities(300)
	if err := Mgr.SelectListBySift(es, sDeletionEQ, sGenreOA, sOrdinalOA); err == nil {
		dicts := make([]*xtype.Dict, es.Len())
		for i, d := range es.Values() {
			e := &xtype.Dict{
				Mkey:    d.Mkey(),
				Mval:    d.Mval(),
				Genre:   d.Genre(),
				Filters: d.Filters(),
				Descr:   d.Descr(),
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

// Refresh the cache data for the dictionary.
func init() {
	dict.RegisterCache(cacheField, initCache)
}
