// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func mergeFile(dir, fileRegexp, newfile string) error {
	if strings.IsBlank(fileRegexp) {
		fileRegexp = `^insert.[\S]+.sql$`
	}
	if strings.IsBlank(newfile) {
		newfile = "merge-file.sql"
	}
	var b bytes.Buffer
	filepath.Walk(dir, func(subpath string, f os.FileInfo, err error) error {
		if err != nil {
			logger.Error(err)
			return err
		}
		match, _ := regexp.MatchString(fileRegexp, f.Name())
		if f.IsDir() == false && match {
			c, err := files.Read(subpath)
			if err != nil {
				logger.Error(err)
				return err
			}
			b.WriteString(c)
		}
		return nil
	})
	initfile := dir + "/" + newfile
	if files.IsExist(initfile) {
		err := files.Remove(initfile)
		if err != nil {
			logger.Error(err)
			return err
		}
	}
	return files.Write(initfile, b.String(), 0755)
}
