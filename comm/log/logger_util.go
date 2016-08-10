// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
)

// Sets the output prefix for the logger.
func (me *Logger) setPrefix(priority int) {
	prefix := ""
	if me.prefix != "" {
		prefix = me.prefix + " "
	}
	if me.logger.Flags()&Lpriority != 0 {
		switch priority {
		case Ptrace:
			me.logger.SetPrefix(prefix + "Trace ")
		case Pdebug:
			me.logger.SetPrefix(prefix + "Debug ")
		case Pinfo:
			me.logger.SetPrefix(prefix + "Info ")
		case Pwarn:
			me.logger.SetPrefix(prefix + "Warn ")
		case Perror:
			me.logger.SetPrefix(prefix + "Error ")
		case Pcritical:
			me.logger.SetPrefix(prefix + "Critical ")
		case Pprint:
			me.logger.SetPrefix(prefix + "Print ")
		default:
			me.logger.SetPrefix(prefix)
		}
	} else {
		me.logger.SetPrefix(prefix)
	}
}

// print calls me.logger.Output to print to the logger.
func (me *Logger) print(priority int, v ...interface{}) {
	if priority >= me.priority {
		me.setPrefix(priority)
		me.logger.Output(me.calldepth(), fmt.Sprint(v...))
	}
}

// printf calls me.logger.Output to print to the logger.
func (me *Logger) printf(priority int, format string, v ...interface{}) {
	if priority >= me.priority {
		me.setPrefix(priority)
		me.logger.Output(me.calldepth(), fmt.Sprintf(format, v...))
	}
}

// println calls me.logger.Output to print to the logger.
func (me *Logger) println(priority int, v ...interface{}) {
	if priority >= me.priority {
		me.setPrefix(priority)
		me.logger.Output(me.calldepth(), fmt.Sprintln(v...))
	}
}

func (me *Logger) calldepth() int {
	calldepth := 4
	if me.isConsole {
		calldepth++
	}
	return calldepth
}
