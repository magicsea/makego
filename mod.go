package main

import (
	"path/filepath"
	"strings"
)

type Model struct {
	Structs []*SrcStruct
}

type SrcStruct struct {
	data []string
	file string
}

func (s SrcStruct) Data(i int) string {
	return s.data[i]
}

func (s SrcStruct) FileName() string {
	ext := filepath.Ext(s.file)
	sn := strings.TrimSuffix(s.file, ext)
	return sn
}

func (s SrcStruct) FilePath() string {
	return s.file
}
