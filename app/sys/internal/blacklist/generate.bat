@echo off
echo [INFO] run go generate.

cd %~dp0
set I18N_LOCALE=en_US
call go generate
pause