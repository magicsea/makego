// 根据go文件生成索引文件
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

var templateFile = flag.String("tmpfile", "example/tmpgo.tpl", "模板文件")
var outputFile = flag.String("output", "example/ouput.go", "生成文件")
var srcDir = flag.String("src", "example/srcgo", "源代码目录")
var srcFilePattern = flag.String("srcfp", "*able.go", "筛选文件样式")
var srcStructPattern = flag.String("srcstruct", `type (?s:(\w+TableTable)) struct {`, "源文件结构样式")

func main() {
	flag.Parse()
	log.Println("templateFile:", *templateFile)
	log.Println("outputFile:", *outputFile)
	log.Println("srcDir:", *srcDir)
	log.Println("srcFilePattern:", *srcFilePattern)
	log.Println("srcStructPattern:", *srcStructPattern)
	log.Println("==============start==============")
	data, err := ioutil.ReadFile(*templateFile)
	if err != nil {
		panic(err)
	}
	tpl, err := template.New("golang").Parse(string(data))
	if err != nil {
		log.Fatalln(err)
	}

	var model = Model{}
	errp := parse(&model)
	if errp != nil {
		panic(errp)
	}

	bf := NewStream(nil, nil)
	err = tpl.Execute(bf.Buffer(), &model)
	if err != nil {
		panic(err)
	}

	errW := bf.WriteFile(*outputFile)
	if errW != nil {
		log.Println("write error:", errW)
		os.Exit(1)
	}
	log.Println("==============generate success=>", *outputFile)
}
