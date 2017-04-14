@echo off
echo [INFO] Export assets to local repository.

cd %~dp0
rd /s /q \app\assets\gyadm
md \app\assets\gyadm
md \app\assets\gyadm\sys
rd /s /q \app\assets\hndev\ui
md \app\assets\hndev\ui
cd ..\..
xcopy sys\static\adm \app\assets\gyadm\sys /E/Y
xcopy adm\static \app\assets\gyadm /E/Y
xcopy ui\static \app\assets\gydev\ui /E/Y
pause