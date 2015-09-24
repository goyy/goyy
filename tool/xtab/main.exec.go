// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
)

func main() {
	isEntity := flag.Bool("entity", false, "is generated entity")
	isSQL := flag.Bool("sql", false, "is generated SQL")
	flag.Parse()
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
}
