#!/bin/sh

echo [INFO] run go generate.

export I18N_LOCALE=en_US
go generate
