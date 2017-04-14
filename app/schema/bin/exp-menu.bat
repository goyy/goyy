@echo off
echo [INFO] Create Menu Files.

cd %~dp0
cd ..
set I18N_LOCALE=en_US
call xtab -menu

pause