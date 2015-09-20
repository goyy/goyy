// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package profile

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	actives  []string
	defaults = []string{DEFAULT}
)

// Return the set of profiles explicitly made active for this environment.
// Profiles are used for creating logical groupings of bean definitions to be registered conditionally,
// for example based on deployment environment.
// Profiles can be activated by calling profile.SetActives(...string).
// If no profiles have explicitly been specified as active, then any default profiles will automatically be activated.
func Actives() []string {
	return actives
}

func SetActives(profiles ...string) {
	i := 0
	for _, v := range profiles {
		if strings.IsNotBlank(v) {
			if i == 0 {
				actives = []string{}
			}
			actives = append(actives, v)
			i++
		}
	}
}

// Return the set of profiles explicitly made active for this environment.
// Profiles are used for creating logical groupings of bean definitions to be registered conditionally,
// for example based on deployment environment.
// Profiles can be activated by calling profile.SetDefaults(...string).
// If no profiles have explicitly been specified as active, then any default profiles will automatically be activated.
func Defaults() []string {
	return defaults
}

func SetDefaults(profiles ...string) {
	i := 0
	for _, v := range profiles {
		if strings.IsNotBlank(v) {
			if i == 0 {
				defaults = []string{}
			}
			defaults = append(defaults, v)
			i++
		}
	}
}

// Return whether one or more of the given profiles is active or,
// in the case of no explicit active profiles,
// whether one or more of the given profiles is included in the set of default profiles.
// If a profile begins with '!' the logic is inverted,
// i.e. the method will return true if the given profile is not active.
// For example, profile.Accepts("p1", "!p2")
// will return true if profile 'p1' is active or 'p2' is not active.
func Accepts(profiles ...string) bool {
	if profiles == nil || len(profiles) == 0 {
		return false
	}
	if actives == nil || len(actives) == 0 {
		return isAccepts(_defaults_, profiles...)
	} else {
		return isAccepts(_actives_, profiles...)
	}
}

func isExistActive(name string) bool {
	return isExist(_actives_, name)
}

func isNotExistActive(name string) bool {
	return isNotExist(_actives_, name)
}

func isExistDefault(name string) bool {
	return isExist(_defaults_, name)
}

func isNotExistDefault(name string) bool {
	return isNotExist(_defaults_, name)
}

func isExist(typ, name string) bool {
	if strings.IsBlank(name) {
		return false
	}
	var data []string
	if _actives_ == typ {
		data = actives
	} else {
		data = defaults
	}
	for _, v := range data {
		if name == v {
			return true
		}
	}
	return false
}

func isNotExist(typ, name string) bool {
	if strings.IsBlank(name) {
		return true
	}
	var data []string
	if _actives_ == typ {
		data = actives
	} else {
		data = defaults
	}
	for _, v := range data {
		if name == v {
			return false
		}
	}
	return true
}

func isAccepts(typ string, profiles ...string) bool {
	for _, v := range profiles {
		if strings.IsBlank(v) {
			continue
		}
		val := strings.TrimSpace(v)
		if _reverse_ == strings.Left(val, 1) {
			c := strings.After(val, _reverse_)
			if _actives_ == typ {
				if isNotExistActive(c) {
					return true
				}
			} else {
				if isNotExistDefault(c) {
					return true
				}
			}
		} else {
			if _actives_ == typ {
				if isExistActive(val) {
					return true
				}
			} else {
				if isExistDefault(val) {
					return true
				}
			}
		}
	}
	return false
}
