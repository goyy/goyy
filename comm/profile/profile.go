// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package profile

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
)

var (
	actives  []string
	defaults = DEFAULT
)

// Return the set of profiles explicitly made active for this environment.
// Profiles are used for creating logical groupings of bean definitions to be registered conditionally,
// for example based on deployment environment.
// Profiles can be activated by calling profile.SetActives(...string).
// If no profiles have explicitly been specified as active, then default profile will automatically be activated.
func Actives() []string {
	return actives
}

func SetActives(profiles ...string) {
	i := 0
	for _, v := range profiles {
		if strings.IsNotBlank(v) {
			if i == 0 {
				actives = make([]string, 0, len(profiles))
			}
			actives = append(actives, v)
			i++
		}
	}
}

// Return the set of profiles explicitly made active for this environment.
// Profiles are used for creating logical groupings of bean definitions to be registered conditionally,
// for example based on deployment environment.
// Profiles can be activated by calling profile.SetDefault(string).
// If no profiles have explicitly been specified as active, then default profile will automatically be activated.
func Default() string {
	return defaults
}

func SetDefault(profile string) {
	if strings.IsNotBlank(profile) {
		defaults = profile
	}
}

// Return whether one or more of the given profiles is active or,
// in the case of no explicit active profiles,
// whether one or more of the given profiles is included in the set of default profile.
// If a profile begins with '!' the logic is inverted,
// i.e. the method will return true if the given profile is not active.
// For example, profile.Accepts("p1", "!p2")
// will return true if profile 'p1' is active or 'p2' is not active.
func Accepts(profiles ...string) bool {
	if profiles == nil || len(profiles) == 0 {
		return false
	}
	for _, v := range profiles {
		if strings.IsBlank(v) {
			continue
		}
		val := strings.TrimSpace(v)
		if reverse == strings.Left(val, 1) {
			c := strings.After(val, reverse)
			if isNotExist(c) {
				return true
			}
		} else {
			if isExist(val) {
				return true
			}
		}
	}
	return false
}

func isExist(name string) bool {
	if strings.IsBlank(name) {
		return false
	}
	for _, v := range actives {
		if name == v {
			return true
		}
	}
	return name == defaults
}

func isNotExist(name string) bool {
	if strings.IsBlank(name) {
		return true
	}
	for _, v := range actives {
		if name == v {
			return false
		}
	}
	return name != defaults
}
