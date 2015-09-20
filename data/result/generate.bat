@echo off
echo [INFO] run go generate.

cd %~dp0
call go generate
pause