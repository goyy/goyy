@echo off
echo [INFO] Create DDL SQL Files.

cd %~dp0
call xtab -sql

pause