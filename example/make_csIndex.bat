
@echo off

..\bin\windows\makego.exe --tmpfile=tmpcs.tpl --src=.\srccs\ --output=.\output.cs  --srcstruct="public partial class (?s:(\w+)) : ITable" --srcfp=*.cs

