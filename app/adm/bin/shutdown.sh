#!/bin/sh
kill -9 `ps -ef | grep /adm | grep -v 'grep' | awk '{print $2}'`
exit
