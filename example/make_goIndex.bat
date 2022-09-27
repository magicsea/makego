
@echo off

..\bin\windows\makego.exe --tmpfile=tmpgo.tpl --src=.\srcgo\ --output=.\output.go  --srcstruct="type (?s:(\w+TableTable)) struct {" --srcfp=*.go

