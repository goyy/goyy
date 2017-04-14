// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjWebConfLog = `<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE env PUBLIC "-//GOYY//DTD ENV 1.0//EN" "http://goyy.org/dtd/env-1.0.dtd">
<configuration>
	<environments default="development">
		<environment id="development">
			<log name="<%.NewProjName%>">
				<!-- 0:Pall 1:Ptrace 2:Pdebug 3:Pinfo 4:Pwarn 5:Perror 6:Pcritical 7:Pprint 8:Poff -->
				<priority>2</priority> <!-- log.Pdebug -->
				<!-- 1:Ldate 2:Ltime 4:Lmicroseconds 8:Llongfile 16:Lshortfile 32:LUTC 64:Lpriority 3:LstdFlags(Ldate | Ltime) -->
				<layout>75</layout> <!-- log.LstdFlags | log.Lpriority | log.Llongfile -->
				<!-- 1:Oconsole 2:Odailyfile 4:Orollingfile 3:Ostd(Oconsole | Odailyfile) -->
				<output>1</output> <!-- log.Oconsole -->
				<dir>logs</dir>
			</log>
		</environment>
		<environment id="test">
			<log name="<%.NewProjName%>">
				<!-- 0:Pall 1:Ptrace 2:Pdebug 3:Pinfo 4:Pwarn 5:Perror 6:Pcritical 7:Pprint 8:Poff -->
				<priority>2</priority> <!-- log.Pdebug -->
				<!-- 1:Ldate 2:Ltime 4:Lmicroseconds 8:Llongfile 16:Lshortfile 32:LUTC 64:Lpriority 3:LstdFlags(Ldate | Ltime) -->
				<layout>75</layout> <!-- log.LstdFlags | log.Lpriority | log.Llongfile -->
				<!-- 1:Oconsole 2:Odailyfile 4:Orollingfile 3:Ostd(Oconsole | Odailyfile) -->
				<output>2</output> <!-- log.Odailyfile -->
				<dir>logs</dir>
			</log>
		</environment>
		<environment id="production">
			<log name="<%.NewProjName%>">
				<!-- 0:Pall 1:Ptrace 2:Pdebug 3:Pinfo 4:Pwarn 5:Perror 6:Pcritical 7:Pprint 8:Poff -->
				<priority>5</priority> <!-- log.Perror -->
				<!-- 1:Ldate 2:Ltime 4:Lmicroseconds 8:Llongfile 16:Lshortfile 32:LUTC 64:Lpriority 3:LstdFlags(Ldate | Ltime) -->
				<layout>75</layout> <!-- log.LstdFlags | log.Lpriority | log.Llongfile -->
				<!-- 1:Oconsole 2:Odailyfile 4:Orollingfile 3:Ostd(Oconsole | Odailyfile) -->
				<output>2</output> <!-- log.Odailyfile -->
				<dir>/app/logs/<%.NewProjName%>web</dir>
			</log>
		</environment>
	</environments>
</configuration>
`
