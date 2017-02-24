// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package envs implements env utility functions.
package envs

import (
	"os"
	"runtime"

	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// ParseGOPATH returns path support using the %GOPATH% environment variable.
// eg. %GOPATH%/src/gopkg.in/goyy/goyy.v0  =>  ~/gopath/src/gopkg.in/goyy/goyy.v0
func ParseGOPATH(path string) string {
	if strings.IsBlank(path) {
		return ""
	}
	GOPATH := "%GOPATH%"
	if !strings.Contains(path, GOPATH) {
		return path
	}
	gopath := os.Getenv("GOPATH")
	split := ":"
	if strings.ToLower(runtime.GOOS) == "windows" {
		split = ";"
	}
	gopaths := strings.Split(gopath, split)
	for _, v := range gopaths {
		file := strings.Replace(path, GOPATH, v, -1)
		if files.IsExist(file) {
			return file
		}
	}
	return path
}
