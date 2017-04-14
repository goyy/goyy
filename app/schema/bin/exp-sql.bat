@echo off
echo [INFO] Create DDL SQL Files.

cd %~dp0
cd ..
set I18N_LOCALE=en_US
call xtab -sql

pause