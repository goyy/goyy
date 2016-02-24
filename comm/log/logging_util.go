// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"os"
	"time"
)

func (me *Logging) setConsoleLogger() {
	me.console = NewLogger(os.Stderr)
	me.console.SetPriority(me.priority)
	if me.layouts > 0 {
		me.console.SetLayouts(me.layouts)
	}
	me.console.SetPrefix(me.prefix)
}

func (me *Logging) resetConsoleLogger() {
	if me.outputs&Oconsole != 0 {
		if me.console == nil {
			me.setConsoleLogger()
		}
	}
}

func (me *Logging) getDailyFileName() string {
	date := time.Now().Format("2006-01-02")
	format := "./%s/daily.%s.log"
	return fmt.Sprintf(format, logDir, date)
}

func (me *Logging) setDailyFileLogger() {
	me.dailyfilename = me.getDailyFileName()
	if me.isExist(me.dailyfilename) {
		f, err := os.OpenFile(me.dailyfilename, os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			fmt.Println(err.Error())
		}
		me.dailyfile = NewLogger(f)
	} else {
		os.Mkdir(logDir, 0666)
		f, err := os.Create(me.dailyfilename)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			date := time.Now().Format("2006/01/02 15:04:05")
			_, err = f.WriteString("[log] Print " + date + " Create file：time\r\n")
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		me.dailyfile = NewLogger(f)
	}
	me.dailyfile.SetPriority(me.priority)
	if me.layouts > 0 {
		me.dailyfile.SetLayouts(me.layouts)
	}
	me.dailyfile.SetPrefix(me.prefix)
}

// Reports whether the specified file exists.
// Returns true if the file exists, false if it does not exist.
func (me *Logging) isExist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func (me *Logging) resetDailyFileLogger() {
	if me.outputs&Odailyfile != 0 {
		if me.dailyfilename != me.getDailyFileName() {
			me.setDailyFileLogger()
		}
	}
}
