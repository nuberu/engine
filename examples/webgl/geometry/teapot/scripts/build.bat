@echo off

cd ..
copy /Y %GOROOT%\misc\wasm\wasm_exec.js .\dist\
xcopy /Y .\*.html .\dist\

set GOOS=js
set GOARCH=wasm
go build -o .\dist\main.wasm main.go