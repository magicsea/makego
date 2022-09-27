// Generated by g.f.com/flora/weiss/tabtoy
// Version: v1.0.23
// DO NOT EDIT!!
package table

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Defined in table: BoxTable
type BoxTable struct {

	// Box
	Box []*BoxDefine
}

// Defined in table: Box
type SelectDrop struct {

	// 类型
	ItemType ItemType

	// ID
	ItemID uint64

	// 数量
	Num uint32
}

// Defined in table: Box
type BoxDefine struct {

	// ID
	ID uint64

	// 类型
	Type BoxType

	// 自选掉落
	Drop1 []*SelectDrop

	// 非自选掉落
	Drop2 []uint32

	// 堆叠上限
	Stack uint32
}

// BoxTable 访问接口
type BoxTableTable struct {

	// 表格原始数据
	BoxTable

	// 索引函数表
	indexFuncByName map[string][]func(*BoxTableTable) error

	// 清空函数表
	clearFuncByName map[string][]func(*BoxTableTable) error

	// 加载前回调
	preFuncList []func(*BoxTableTable) error

	// 加载后回调
	postFuncList []func(*BoxTableTable) error

	// 索引ID
	BoxByID map[uint64]*BoxDefine
}

// 从json文件加载
func (self *BoxTableTable) Load(filename string) error {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	return self.LoadData(data)
}

// 从二进制加载
func (self *BoxTableTable) LoadData(data []byte) error {

	var newTab BoxTable

	// 读取
	err := json.Unmarshal(data, &newTab)
	if err != nil {
		return err
	}

	// 所有加载前的回调
	for _, v := range self.preFuncList {
		if err = v(self); err != nil {
			return err
		}
	}

	// 清除前通知
	for _, list := range self.clearFuncByName {
		for _, v := range list {
			if err = v(self); err != nil {
				return err
			}
		}
	}

	// 复制数据
	self.BoxTable = newTab

	// 生成索引
	for _, list := range self.indexFuncByName {
		for _, v := range list {
			if err = v(self); err != nil {
				return err
			}
		}
	}

	// 所有完成时的回调
	for _, v := range self.postFuncList {
		if err = v(self); err != nil {
			return err
		}
	}

	return nil
}

// 注册外部索引入口, 索引回调, 清空回调
func (self *BoxTableTable) RegisterIndexEntry(name string, indexCallback func(*BoxTableTable) error, clearCallback func(*BoxTableTable) error) {

	indexList, _ := self.indexFuncByName[name]
	clearList, _ := self.clearFuncByName[name]

	if indexCallback != nil {
		indexList = append(indexList, indexCallback)
	}

	if clearCallback != nil {
		clearList = append(clearList, clearCallback)
	}

	self.indexFuncByName[name] = indexList
	self.clearFuncByName[name] = clearList
}

// 注册加载前回调
func (self *BoxTableTable) RegisterPreEntry(callback func(*BoxTableTable) error) {

	self.preFuncList = append(self.preFuncList, callback)
}

// 注册所有完成时回调
func (self *BoxTableTable) RegisterPostEntry(callback func(*BoxTableTable) error) {

	self.postFuncList = append(self.postFuncList, callback)
}

// 创建一个BoxTable表读取实例
func NewBoxTableTable() *BoxTableTable {
	return &BoxTableTable{

		indexFuncByName: map[string][]func(*BoxTableTable) error{

			"Box": {func(tab *BoxTableTable) error {

				// Box
				for _, def := range tab.Box {

					if _, ok := tab.BoxByID[def.ID]; ok {
						panic(fmt.Sprintf("duplicate index in BoxByID: %v", def.ID))
					}

					tab.BoxByID[def.ID] = def

				}

				return nil
			}},
		},

		clearFuncByName: map[string][]func(*BoxTableTable) error{

			"Box": {func(tab *BoxTableTable) error {

				// Box

				tab.BoxByID = make(map[uint64]*BoxDefine)

				return nil
			}},
		},

		BoxByID: make(map[uint64]*BoxDefine),
	}
}
