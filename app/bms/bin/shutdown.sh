#!/bin/sh
kill -9 `ps -ef | grep /bms | grep -v 'grep' | awk '{print $2}'`

exit
