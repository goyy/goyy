// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"

	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/data/xsql"
)

func main() {
	xsql.SetPriority(log.Perror)
	isEntity := flag.Bool("entity", false, "is generated entity")
	isSQL := flag.Bool("sql", false, "is generated SQL")
	isMenu := flag.Bool("menu", false, "is generated menu")
	isDict := flag.Bool("dict", false, "is generated dict")
	isDB := flag.Bool("db", false, "is generated DB")

	isMerge := flag.Bool("merge", false, "is merged file")
	fileRegexp := flag.String("regexp", `^insert.[\S]+.sql$`, "file regexp")
	newfile := flag.String("newfile", "init.sql", "merged new file name")

	proj := flag.String("proj", "goyy", "project name")
	pkg := flag.String("pkg", "", "path to package for new project")
	params := flag.String("params", "goyy", "params")
	flag.Parse()
	if *isMerge {
		logger.Println("Merging file : start")
		mergeFile(".", *fileRegexp, *newfile)
		logger.Println("Merged file : end")
	} else {
		initgen(*proj, *pkg)
	}
	if *isSQL {
		logger.Println("Exporting sql : start")
		expSQL()
		logger.Println("Exported sql : end")
	}
	if *isEntity {
		logger.Println("Generating entity : start")
		genEntity()
		logger.Println("Generated entity : end")
	}
	if *isMenu {
		logger.Println("Generating menu : start")
		genMenu(*params)
		logger.Println("Generated menu : end")
	}
	if *isDict {
		logger.Println("Generating dict : start")
		genDict(*params)
		logger.Println("Generated dict : end")
	}
	if *isDB {
		logger.Println("Initialize db : start")
		expDB(*params)
		logger.Println("Initialize db : end")
	}
}
