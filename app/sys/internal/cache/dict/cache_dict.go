package dict

import (
	"strconv"
	"sync"

	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/cache"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

const (
	cacheKey   = "sys_dict_cache_key"
	cacheField = "dict"
)

var (
	cacheDatas = make(map[string]*dictCache, 0)
	mutex      sync.Mutex
)

type dictCache struct {
	ver  int // If the version number is not equal to the value of the cache, you want to get a new value from the database
	data []*xtype.Dict
	hash map[string]*xtype.Dict // In order to quickly find a dictionary based on Genre and Mkey：map.key=Genre+Mkey
	init func() []*xtype.Dict
}

// Register cache for the dictionary.
func RegisterCache(field string, init func() []*xtype.Dict) {
	if _, ok := cacheDatas[field]; !ok {
		xhttp.RegisterPreRun(func() {
			if !cache.HExists(cacheKey, field) { // Register cache when the default version of redis is 1
				if err := cache.HSet(cacheKey, field, "1"); err == nil {
				} else {
					logger.Errorln(err)
				}
			}
			ifSetCache(field, true)
		})
	}
	c := &dictCache{init: init}
	cacheDatas[field] = c
}

// To determine whether there is a corresponding field cache.
func HasCache(field string) bool {
	_, ok := cacheDatas[field]
	return ok
}

// Get cached data.
func GetCache(field string) []*xtype.Dict {
	ifSetCache(field, false)
	if val, ok := cacheDatas[field]; ok {
		return val.data
	} else {
		return nil
	}
}

// Gets the cached data for the map type.
func GetHashCache(field string) map[string]*xtype.Dict {
	ifSetCache(field, false)
	if val, ok := cacheDatas[field]; ok {
		return val.hash
	} else {
		return nil
	}
}

// Does not determine whether the cache for the new version,
// directly access to the map type of cache data.
func GetHashNoCache(field string) map[string]*xtype.Dict {
	if val, ok := cacheDatas[field]; ok {
		return val.hash
	} else {
		return nil
	}
}

// Refresh cache data: redis cache version +1
func RefreshCache(field string) {
	if v, err := cache.HGet(cacheKey, field); err == nil {
		if ver, err := strconv.Atoi(v); err == nil {
			if err = cache.HSet(cacheKey, field, strconv.Itoa(ver+1)); err == nil {
				ifSetCache(field, true)
			} else {
				logger.Errorln(err)
			}
		} else {
			logger.Errorln(err)
		}
	} else {
		logger.Errorln(err)
	}
}

// Set cache data.
func ifSetCache(field string, isPROD bool) {
	// When the production environment starts and the cache is refreshed:
	// the dictionary is initialized, which improves the running performance
	// and reduces the service startup speed.
	//
	// Development and testing of the environment for the first time
	// to obtain a dictionary to initialize the dictionary:
	// to improve service startup speed, reduce the performance of the operation.
	if profile.Accepts(profile.PROD) == isPROD {
		setCache(field)
	}
}

// Set cache data.
func setCache(field string) {
	if val, ok := cacheDatas[field]; ok {
		if v, err := cache.HGet(cacheKey, field); err == nil {
			if ver, err := strconv.Atoi(v); err == nil {
				if val.ver != ver {
					mutex.Lock()
					if val.ver != ver {
						val.data = val.init()
						val.ver = ver
						if val.hash == nil {
							val.hash = make(map[string]*xtype.Dict, 0)
						}
						for _, data := range val.data {
							key := data.Genre + data.Mkey
							val.hash[key] = data
						}
					}
					mutex.Unlock()
				}
			} else {
				logger.Errorln(err)
			}
		} else {
			logger.Errorln(err)
		}
	}
}
