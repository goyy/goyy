// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"strings"

	"gopkg.in/goyy/goyy.v0/util/files"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	// entity
	entipath := flag.String("entity", "", "entity file path")
	admPath := flag.String("admpath", "", "import path for admin project")
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
	newProjPkg := flag.String("pkg", "", "path to package for new project")
	newProjTitle := flag.String("title", "", "title of new project")
	newProjHost := flag.String("host", "", "host of new project")
	flag.Parse()
	f := factory{
		AdmPath:          *admPath,
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
		NewProjTitle:     *newProjTitle,
		NewProjHost:      *newProjHost,
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
	if f.HasGenProj {
		if strings.TrimSpace(*newProjPkg) == "" {
			dir, err := files.AbsDir(".")
			if err != nil {
				log.Println(err.Error())
			}
			dir = strings.Replace(dir, "\\", "/", -1)
			gopath := os.Getenv("GOPATH")
			split := ":"
			if strings.ToLower(runtime.GOOS) == "windows" {
				split = ";"
				gopath = strings.ToLower(gopath)
				dir = strings.ToLower(dir)
			}
			gopath = strings.Replace(gopath, "\\", "/", -1)
			gopaths := strings.Split(gopath, split)
			for _, v := range gopaths {
				if strings.Contains(dir, v) {
					f.NewProjPkg = strings.Replace(dir, v+"/src/", "", 1)
					break
				}
			}
		}
		if strings.TrimSpace(*newProjTitle) == "" {
			f.NewProjTitle = f.NewProjName
		}
		if strings.TrimSpace(*newProjHost) == "" {
			f.NewProjHost = f.NewProjName + ".com"
		} else {
			host := strings.ToLower(*newProjHost)
			if strings.HasPrefix(host, "http://") {
				host = strings.TrimLeft(host, "http://")
			}
			if strings.HasPrefix(host, "https://") {
				host = strings.TrimLeft(host, "https://")
			}
			if strings.HasPrefix(host, "www.") {
				host = strings.TrimLeft(host, "www.")
			}
			f.NewProjHost = host
		}
	}
	if err := f.Write(); err != nil {
		log.Printf("%v", err)
		return
	}
}
