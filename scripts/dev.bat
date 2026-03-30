@echo off
setlocal
set ROOT=%~dp0..

pushd "%ROOT%"
air -c .air.toml
set EXIT=%errorlevel%
popd

exit /b %EXIT%
