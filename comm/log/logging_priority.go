// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

// Trace calls Output to print to the logging with the Trace priority.
func (me *Logging) Trace(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Trace(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Trace(v...)
		}
	}
}

// Tracef calls Output to printf to the logging with the Trace priority.
func (me *Logging) Tracef(format string, v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Tracef(format, v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Tracef(format, v...)
		}
	}
}

// Traceln calls Output to println to the logging with the Trace priority.
func (me *Logging) Traceln(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Traceln(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Traceln(v...)
		}
	}
}

// Debug calls Output to print to the logging with the Debug priority.
func (me *Logging) Debug(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Debug(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Debug(v...)
		}
	}
}

// Debugf calls Output to printf to the logging with the Debug priority.
func (me *Logging) Debugf(format string, v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Debugf(format, v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Debugf(format, v...)
		}
	}
}

// Debugln calls Output to println to the logging with the Debug priority.
func (me *Logging) Debugln(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Debugln(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Debugln(v...)
		}
	}
}

// Info calls Output to print to the logging with the Info priority.
func (me *Logging) Info(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Info(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Info(v...)
		}
	}
}

// Infof calls Output to printf to the logging with the Info priority.
func (me *Logging) Infof(format string, v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Infof(format, v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Infof(format, v...)
		}
	}
}

// Infoln calls Output to println to the logging with the Info priority.
func (me *Logging) Infoln(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Infoln(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Infoln(v...)
		}
	}
}

// Warn calls Output to print to the logging with the Warn priority.
func (me *Logging) Warn(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Warn(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Warn(v...)
		}
	}
}

// Warnf calls Output to printf to the logging with the Warn priority.
func (me *Logging) Warnf(format string, v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Warnf(format, v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Warnf(format, v...)
		}
	}
}

// Warnln calls Output to println to the logging with the Warn priority.
func (me *Logging) Warnln(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Warnln(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Warnln(v...)
		}
	}
}

// Error calls Output to print to the logging with the Error priority.
func (me *Logging) Error(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Error(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Error(v...)
		}
	}
}

// Errorf calls Output to printf to the logging with the Error priority.
func (me *Logging) Errorf(format string, v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Errorf(format, v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Errorf(format, v...)
		}
	}
}

// Errorln calls Output to println to the logging with the Error priority.
func (me *Logging) Errorln(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Errorln(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Errorln(v...)
		}
	}
}

// Critical calls Output to print to the logging with the Critical priority.
func (me *Logging) Critical(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Critical(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Critical(v...)
		}
	}
}

// Criticalf calls Output to printf to the logging with the Critical priority.
func (me *Logging) Criticalf(format string, v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Criticalf(format, v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Criticalf(format, v...)
		}
	}
}

// Criticalln calls Output to println to the logging with the Critical priority.
func (me *Logging) Criticalln(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Criticalln(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Criticalln(v...)
		}
	}
}

// Print calls Output to print to the logging with the Print priority.
func (me *Logging) Print(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Print(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Print(v...)
		}
	}
}

// Printf calls Output to printf to the logging with the Print priority.
func (me *Logging) Printf(format string, v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Printf(format, v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Printf(format, v...)
		}
	}
}

// Println calls Output to println to the logging with the Print priority.
func (me *Logging) Println(v ...interface{}) {
	me.pre()
	if me.outputs&Oconsole != 0 {
		me.resetConsoleLogger()
		if me.console != nil {
			me.console.Println(v...)
		}
	}
	if me.outputs&Odailyfile != 0 {
		me.resetDailyFileLogger()
		if me.dailyfile != nil {
			me.dailyfile.Println(v...)
		}
	}
}
