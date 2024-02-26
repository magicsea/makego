package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
)

func parse(model *Model) error {
	err := filepath.Walk(*srcDir, func(path string, info fs.FileInfo, err error) error {
		c := filepath.ToSlash(path)
		b, err := filepath.Match(*srcFilePattern, filepath.Base(c))
		if err != nil {
			return err
		}
		if b {
			if err := parseFile(path, model); err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func parseFile(path string, model *Model) error {
	filename := filepath.Base(path)
	log.Println("parseFile:", path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	//type BoxTableTable struct {
	regNs := regexp.MustCompile(*srcStructPattern)
	ns := regNs.FindAllStringSubmatch(string(data), -1)
	if len(ns) > 0 {
		for _, n := range ns {
			//name := strings.TrimSpace(n[1])
			log.Println("struct:", n, filename)
			model.Structs = append(model.Structs, &SrcStruct{
				data: n,
				file: filename,
			})
		}
	}
	return nil
}
