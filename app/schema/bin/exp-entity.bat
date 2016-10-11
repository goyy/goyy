@echo off
echo [INFO] Create Entity Files.

cd %~dp0
cd ..
set I18N_LOCALE=en_US
call xtab -entity

pause