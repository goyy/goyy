// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

// Trace calls output to print to the logger with the Trace priority.
func (me *Logger) Trace(v ...interface{}) {
	me.print(Ptrace, v...)
}

// Tracef calls output to printf to the logger with the Trace priority.
func (me *Logger) Tracef(format string, v ...interface{}) {
	me.printf(Ptrace, format, v...)
}

// Traceln calls output to println to the logger with the Trace priority.
func (me *Logger) Traceln(v ...interface{}) {
	me.println(Ptrace, v...)
}

// Debug calls output to print to the logger with the Debug priority.
func (me *Logger) Debug(v ...interface{}) {
	me.print(Pdebug, v...)
}

// Debugf calls output to printf to the logger with the Debug priority.
func (me *Logger) Debugf(format string, v ...interface{}) {
	me.printf(Pdebug, format, v...)
}

// Debugln calls output to println to the logger with the Debug priority.
func (me *Logger) Debugln(v ...interface{}) {
	me.println(Pdebug, v...)
}

// Info calls output to print to the logger with the Info priority.
func (me *Logger) Info(v ...interface{}) {
	me.print(Pinfo, v...)
}

// Infof calls output to printf to the logger with the Info priority.
func (me *Logger) Infof(format string, v ...interface{}) {
	me.printf(Pinfo, format, v...)
}

// Infoln calls output to println to the logger with the Info priority.
func (me *Logger) Infoln(v ...interface{}) {
	me.println(Pinfo, v...)
}

// Warn calls output to print to the logger with the Warn priority.
func (me *Logger) Warn(v ...interface{}) {
	me.print(Pwarn, v...)
}

// Warnf calls output to printf to the logger with the Warn priority.
func (me *Logger) Warnf(format string, v ...interface{}) {
	me.printf(Pwarn, format, v...)
}

// Warnln calls output to println to the logger with the Warn priority.
func (me *Logger) Warnln(v ...interface{}) {
	me.println(Pwarn, v...)
}

// Error calls output to print to the logger with the Error priority.
func (me *Logger) Error(v ...interface{}) {
	me.print(Perror, v...)
}

// Errorf calls output to printf to the logger with the Error priority.
func (me *Logger) Errorf(format string, v ...interface{}) {
	me.printf(Perror, format, v...)
}

// Errorln calls output to println to the logger with the Error priority.
func (me *Logger) Errorln(v ...interface{}) {
	me.println(Perror, v...)
}

// Criticall calls output to print to the logger with the Critical priority.
func (me *Logger) Critical(v ...interface{}) {
	me.print(Pcritical, v...)
}

// Criticalf calls output to printf to the logger with the Critical priority.
func (me *Logger) Criticalf(format string, v ...interface{}) {
	me.printf(Pcritical, format, v...)
}

// Criticalln calls output to println to the logger with the Critical priority.
func (me *Logger) Criticalln(v ...interface{}) {
	me.println(Pcritical, v...)
}

// Print calls output to print to the logger with the Print priority.
func (me *Logger) Print(v ...interface{}) {
	me.print(Pprint, v...)
}

// Printf calls output to printf to the logger with the Print priority.
func (me *Logger) Printf(format string, v ...interface{}) {
	me.printf(Pprint, format, v...)
}

// Println calls output to println to the logger with the Print priority.
func (me *Logger) Println(v ...interface{}) {
	me.println(Pprint, v...)
}
