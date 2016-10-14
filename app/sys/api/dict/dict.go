package dict

import (
	"gopkg.in/goyy/goyy.v0/app/sys/internal/cache/dict"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/schema"
)

// Through dictionary type and key acquisition value.
func Get(genre, mkey string) (*xtype.Dict, bool) {
	return dict.Get(genre, mkey)
}

// Through dictionary type and key acquisition value.
func Mval(genre, mkey string) string {
	return dict.Mval(genre, mkey)
}

// Get a list of values from the dictionary type and the list of keys
// (comma separated).
func Mvals(genre, mkeys string) string {
	return dict.Mvals(genre, mkeys)
}

// Register cache.
func RegisterCache(field string, init func() []*xtype.Dict) {
	dict.RegisterCache(field, init)
}

// Refresh cache data.
func RefreshCache(field string) {
	dict.RefreshCache(field)
}

func init() {
	schema.ParseDict = Mval
}
