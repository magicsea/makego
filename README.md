# makego 
通用代码生成工具，根据原有的代码文件生成新的索引文件。   
设计初衷，由于go的反射能力非常弱，如果要批量注册类型只能手动。 所以做了此工具批量生成注册类型代码。   
适合协议生成，表生成的辅助。


### windows测试步骤
- 执行build.bat编译
- 双击makego.exe，默认会调用example里的go生成

### 示例
example里有两个示例，生成go和cs。
- example/make_csIndex.bat为csharp生成的示例。
- example/make_goIndex.bat为go生成的示例。

### 输入参数
```
var templateFile = flag.String("tmpfile", "example/tmp.tpl", "模板文件")
var outputFile = flag.String("output", "example/ouput.go", "生成文件")
var srcDir = flag.String("src", "example/srcgo", "源代码目录")

//筛选文件样式，可使用“/*.“等符号模糊匹配目录
var srcFilePattern = flag.String("srcfp", "*able.go", "筛选文件样式")

//源文件结构样式。
//使用正则表达式,匹配参数使用?s。匹配内容放再括号内部，匹配出结果也包含关键字。例如：(?s:(\w+TableTable))。
//多个参数使用多个(?s:(\w+))匹配。例如：type (?s:(\w+TableTable)) (?s:(\w+)k) struct {
//多行内容可使用(?s:(.*?))匹配。
var srcStructPattern = flag.String("srcstruct", `type (?s:(\w+TableTable)) struct {`, "源文件结构样式")
```

### 模板参数
- {{range .Structs}}：struct数组
  - {{.FileName}}:源文件名，无后缀名
  - {{.FilePath}}:源文件名，有后缀名
  - {{.Data 1}}:数据区,i对应srcstruct参数里的匹配字段，参数从1开始