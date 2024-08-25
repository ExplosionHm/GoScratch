@echo off
setlocal
set GOROOT=C:\Program Files\Go
set GOPATH=%USERPROFILE%\go
set PATH=%PATH%;%GOROOT%\bin;%GOPATH%\bin

pushd %~dp0
if "%~f1"=="" goto usage
goto run

:usage
echo Usage: run.bat [path_to_main.go]
pause
exit /b 1

:run
go run %*
popd
pause
exit /b 0
