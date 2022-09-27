package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type _HashKey [sha256.Size]byte

type ExtendStream struct {
	buf     bytes.Buffer
	hashTab map[_HashKey]int32
	strTab  map[string]int32
}

func (self *ExtendStream) Len() int {
	if self == nil {
		return 0
	}
	return self.buf.Len()
}

func (self *ExtendStream) WriteBytes(bs []byte) int32 {
	if self.hashTab == nil {
		self.hashTab = map[_HashKey]int32{}
	}

	sum := sha256.Sum256(bs)

	offset, ok := self.hashTab[sum]
	if ok {
		return offset
	}

	offset = int32(self.buf.Len())
	self.hashTab[sum] = offset

	self.WriteVarInt32(int32(len(bs)))
	binary.Write(&self.buf, binary.LittleEndian, bs)

	return offset
}

func (self *ExtendStream) WriteString(str string) int32 {
	if self.strTab == nil {
		self.strTab = map[string]int32{}
	}

	offset, ok := self.strTab[str]
	if ok {
		return offset
	}

	offset = int32(self.buf.Len())
	self.strTab[str] = offset

	rawStr := []byte(str)

	self.WriteVarInt32(int32(len(rawStr)))
	binary.Write(&self.buf, binary.LittleEndian, rawStr)

	return offset
}

func (self *ExtendStream) WriteVarInt32(v int32) {
	buf := make([]byte, binary.MaxVarintLen64)
	l := binary.PutVarint(buf, int64(v))
	binary.Write(&self.buf, binary.LittleEndian, buf[:l])
}

func (self *ExtendStream) Buffer() *bytes.Buffer {
	return &self.buf
}

func NewExtendStream() *ExtendStream {
	return &ExtendStream{}
}

type Stream struct {
	buf                  bytes.Buffer
	varInt               bool
	strStream, stuStream *ExtendStream
}

func (self *Stream) GetStrStream() *ExtendStream {
	return self.strStream
}

func (self *Stream) GetStuStream() *ExtendStream {
	return self.stuStream
}

func (self *Stream) Len() int {
	return self.buf.Len()
}

func (self *Stream) Buffer() *bytes.Buffer {
	return &self.buf
}

func (self *Stream) WriteRawBytes(b []byte) {
	self.buf.Write(b)
}

func (self *Stream) Printf(format string, args ...interface{}) {
	self.buf.WriteString(fmt.Sprintf(format, args...))
}

func (self *Stream) WriteFile(outfile string) error {
	// 自动创建目录
	os.MkdirAll(filepath.Dir(outfile), 0755)

	err := ioutil.WriteFile(outfile, self.buf.Bytes(), 0666)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (self *Stream) WriteFixedInt16(v int16) {
	binary.Write(&self.buf, binary.LittleEndian, v)
}

func (self *Stream) WriteFixedInt32(v int32) {
	binary.Write(&self.buf, binary.LittleEndian, v)
}

func (self *Stream) WriteFixedInt64(v int64) {
	binary.Write(&self.buf, binary.LittleEndian, v)
}

func (self *Stream) WriteVarInt16(v int16) {
	buf := make([]byte, binary.MaxVarintLen64)
	l := binary.PutVarint(buf, int64(v))
	binary.Write(&self.buf, binary.LittleEndian, buf[:l])
}

func (self *Stream) WriteVarInt32(v int32) {
	buf := make([]byte, binary.MaxVarintLen64)
	l := binary.PutVarint(buf, int64(v))
	binary.Write(&self.buf, binary.LittleEndian, buf[:l])
}

func (self *Stream) WriteVarInt64(v int64) {
	buf := make([]byte, binary.MaxVarintLen64)
	l := binary.PutVarint(buf, v)
	binary.Write(&self.buf, binary.LittleEndian, buf[:l])
}

func (self *Stream) WriteFixedUInt16(v uint16) {
	binary.Write(&self.buf, binary.LittleEndian, v)
}

func (self *Stream) WriteFixedUInt32(v uint32) {
	binary.Write(&self.buf, binary.LittleEndian, v)
}

func (self *Stream) WriteFixedUInt64(v uint64) {
	binary.Write(&self.buf, binary.LittleEndian, v)
}

func (self *Stream) WriteVarUInt16(v uint16) {
	buf := make([]byte, binary.MaxVarintLen64)
	l := binary.PutUvarint(buf, uint64(v))
	binary.Write(&self.buf, binary.LittleEndian, buf[:l])
}

func (self *Stream) WriteVarUInt32(v uint32) {
	buf := make([]byte, binary.MaxVarintLen64)
	l := binary.PutUvarint(buf, uint64(v))
	binary.Write(&self.buf, binary.LittleEndian, buf[:l])
}

func (self *Stream) WriteVarUInt64(v uint64) {
	buf := make([]byte, binary.MaxVarintLen64)
	l := binary.PutUvarint(buf, v)
	binary.Write(&self.buf, binary.LittleEndian, buf[:l])
}

func (self *Stream) WriteRefString(v string) {
	offset := self.strStream.WriteString(v)
	self.WriteVarInt32(offset)
}

func (self *Stream) WriteNoRefString(v string) {
	rawStr := []byte(v)
	self.WriteVarInt32(int32(len(rawStr)))
	binary.Write(&self.buf, binary.LittleEndian, rawStr)
}

func (self *Stream) WriteRefBytes(v []byte) {
	offset := self.stuStream.WriteBytes(v)
	self.WriteVarInt32(offset)
}

func (self *Stream) WriteNoRefBytes(v []byte) {
	self.WriteVarInt32(int32(len(v)))
	binary.Write(&self.buf, binary.LittleEndian, v)
}

func NewStream(strStream, stuStream *ExtendStream) *Stream {
	return &Stream{strStream: strStream, stuStream: stuStream, varInt: true}
}
