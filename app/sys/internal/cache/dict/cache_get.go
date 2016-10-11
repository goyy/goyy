package dict

import (
	"bytes"

	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Through dictionary type and key acquisition value.
func Get(genre, mkey string) (*xtype.Dict, bool) {
	if strings.IsBlank(genre) || strings.IsBlank(mkey) {
		return nil, false
	}
	dicts := GetHashCache(cacheField)
	if v, ok := dicts[genre+mkey]; ok {
		return v, true
	} else {
		if HasCache(genre) {
			dicts = GetHashCache(genre)
			if v, ok = dicts[genre+mkey]; ok {
				return v, true
			}
		}
	}
	return nil, false
}

// Through dictionary type and key acquisition value.
func Mval(genre, mkey string) string {
	if v, ok := Get(genre, mkey); ok {
		return v.Mval
	}
	return ""
}

// Get a list of values from the dictionary type and the list of keys
// (comma separated).
func Mvals(genre, mkeys string) string {
	if strings.IsBlank(genre) || strings.IsBlank(mkeys) {
		return ""
	}
	var b bytes.Buffer
	keys := strings.Split(mkeys, ",")
	i := 0
	for _, key := range keys {
		if v := Mval(genre, key); strings.IsNotBlank(v) {
			if i > 0 {
				b.WriteString(",")
			}
			b.WriteString(v)
			i++
		}
	}
	return b.String()
}

// Obtain a list of dictionaries by dictionary type and key list.
func List(dicts []*xtype.Dict) []*xtype.Dict {
	if dicts == nil || len(dicts) == 0 {
		return nil
	}
	results := make([]*xtype.Dict, len(dicts))
	isRedisCount := make(map[string]int, 0)
	cdicts := GetHashCache(cacheField)
	for i, d := range dicts {
		if strings.IsNotBlank(d.Mkey) && strings.IsNotBlank(d.Genre) {
			xd := &xtype.Dict{
				Mkey:  d.Mkey,
				Genre: d.Genre,
			}
			results[i] = xd
			if HasCache(d.Genre) {
				var xdicts map[string]*xtype.Dict
				if _, ok := isRedisCount[d.Genre]; ok {
					xdicts = GetHashNoCache(d.Genre)
				} else {
					isRedisCount[d.Genre] = 0
					xdicts = GetHashCache(d.Genre)
				}
				results[i].Mval = getMval(d.Genre, d.Mkey, d.Descr, xdicts)
			} else {
				results[i].Mval = getMval(d.Genre, d.Mkey, d.Descr, cdicts)
			}
		}
	}
	return results
}

// Get a list of all dictionary types.
func getMval(genre, mkey, descr string, dicts map[string]*xtype.Dict) string {
	if strings.IsNotBlank(genre) && strings.IsNotBlank(mkey) && dicts != nil {
		getValue := func(key string) string {
			if v, ok := dicts[genre+key]; ok {
				if descr == "true" {
					return v.Descr
				} else {
					return v.Mval
				}
			}
			return ""
		}
		if strings.Index(mkey, ",") != -1 {
			var b bytes.Buffer
			keys := strings.Split(mkey, ",")
			i := 0
			for _, key := range keys {
				if v := getValue(key); v != "" {
					if i > 0 {
						b.WriteString(",")
					}
					b.WriteString(v)
					i++
				}
			}
			return b.String()
		} else {
			return getValue(mkey)
		}
	}
	return ""
}

// Get a list of all dictionary types.
func ListByGenre() (dictGenres []string, err error) {
	dictGenres = make([]string, 0, 50)
	dicts := GetCache(cacheField)
	var preGenre string
	for _, v := range dicts {
		genre := v.Genre
		if strings.IsNotBlank(genre) && preGenre != genre {
			preGenre = genre
			dictGenres = append(dictGenres, genre)
		}
	}
	return
}
