// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"strings"

	"gopkg.in/goyy/goyy.v0/util/files"
)

func main() {
	// entity
	entipath := flag.String("entity", "", "entity file path")
	clipath := flag.String("clipath", "", "import path for client project")
	apipath := flag.String("apipath", "", "import path for api project")
	tstpath := flag.String("tstpath", "", "import path for test project")
	hasScaffold := flag.Bool("scaffold", false, "is generated service and controller")
	hasService := flag.Bool("service", false, "is generated service")
	hasController := flag.Bool("controller", false, "is generated controller")
	hasDto := flag.Bool("dto", false, "is generated dto")
	hasAPI := flag.Bool("api", false, "is generated api")
	hasSQL := flag.Bool("sql", false, "is generated sql")
	hasLog := flag.Bool("log", false, "is generated log")
	hasUtil := flag.Bool("util", false, "is generated util")
	hasConst := flag.Bool("const", false, "is generated const")
	hasHTML := flag.Bool("html", false, "is generated html")
	hasJs := flag.Bool("js", false, "is generated js")
	// new project
	newProjName := flag.String("new", "", "name of new project")
	flag.Parse()
	f := factory{
		CliPath:          *clipath,
		APIPath:          *apipath,
		TstPath:          *tstpath,
		HasGenService:    *hasService,
		HasGenController: *hasController,
		HasGenDto:        *hasDto,
		HasGenAPI:        *hasAPI,
		HasGenSQL:        *hasSQL,
		HasGenLog:        *hasLog,
		HasGenUtil:       *hasUtil,
		HasGenConst:      *hasConst,
		HasGenHTML:       *hasHTML,
		HasGenJs:         *hasJs,
		NewProjName:      *newProjName,
	}
	if *hasScaffold {
		f.HasGenService = true
		f.HasGenController = true
		f.HasGenDto = false
		f.HasGenAPI = true
		f.HasGenSQL = true
		f.HasGenLog = true
		f.HasGenUtil = true
		f.HasGenConst = true
		f.HasGenHTML = true
		f.HasGenJs = true
	}
	if strings.TrimSpace(*entipath) != "" {
		f.HasGenEntity = true
		if err := f.Init(*entipath); err != nil {
			log.Printf("%v", err)
			return
		}
	}
	if strings.TrimSpace(*newProjName) != "" {
		dir, err := files.AbsDir(".")
		if err != nil {
			log.Println(err.Error())
		}
		f.HasGenProj = true
		f.NewProjPath = dir
	}
	if err := f.Write(); err != nil {
		log.Printf("%v", err)
		return
	}
}
