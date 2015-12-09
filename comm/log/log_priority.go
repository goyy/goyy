// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

// Calls Output to print to the logger with the Trace priority.
func Trace(v ...interface{}) {
	console.Trace(v...)
}

// Calls Output to printf to the logger with the Trace priority.
func Tracef(format string, v ...interface{}) {
	console.Tracef(format, v...)
}

// Calls Output to println to the logger with the Trace priority.
func Traceln(v ...interface{}) {
	console.Traceln(v...)
}

// Calls Output to print to the logger with the Debug priority.
func Debug(v ...interface{}) {
	console.Debug(v...)
}

// Calls Output to printf to the logger with the Debug priority.
func Debugf(format string, v ...interface{}) {
	console.Debugf(format, v...)
}

// Calls Output to println to the logger with the Debug priority.
func Debugln(v ...interface{}) {
	console.Debugln(v...)
}

// Calls Output to print to the logger with the Info priority.
func Info(v ...interface{}) {
	console.Info(v...)
}

// Calls Output to printf to the logger with the Info priority.
func Infof(format string, v ...interface{}) {
	console.Infof(format, v...)
}

// Calls Output to println to the logger with the Info priority.
func Infoln(v ...interface{}) {
	console.Infoln(v...)
}

// Calls Output to print to the logger with the Warn priority.
func Warn(v ...interface{}) {
	console.Warn(v...)
}

// Calls Output to printf to the logger with the Warn priority.
func Warnf(format string, v ...interface{}) {
	console.Warnf(format, v...)
}

// Calls Output to println to the logger with the Warn priority.
func Warnln(v ...interface{}) {
	console.Warnln(v...)
}

// Calls Output to print to the logger with the Error priority.
func Error(v ...interface{}) {
	console.Error(v...)
}

// Calls Output to printf to the logger with the Error priority.
func Errorf(format string, v ...interface{}) {
	console.Errorf(format, v...)
}

// Calls Output to println to the logger with the Error priority.
func Errorln(v ...interface{}) {
	console.Errorln(v...)
}

// Calls Output to print to the logger with the Critical priority.
func Critical(v ...interface{}) {
	console.Critical(v...)
}

// Calls Output to printf to the logger with the Critical priority.
func Criticalf(format string, v ...interface{}) {
	console.Criticalf(format, v...)
}

// Calls Output to println to the logger with the Critical priority.
func Criticalln(v ...interface{}) {
	console.Criticalln(v...)
}

// Calls Output to print to the logger with the Print priority.
func Print(v ...interface{}) {
	console.Print(v...)
}

// Calls Output to printf to the logger with the Print priority.
func Printf(format string, v ...interface{}) {
	console.Printf(format, v...)
}

// Calls Output to println to the logger with the Print priority.
func Println(v ...interface{}) {
	console.Println(v...)
}
