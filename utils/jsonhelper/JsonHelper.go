package jsonhelper

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

// 解析json fils
func Open(path string, value interface{})  {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err.Error())
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)

	Unmarshal(data, value)
}

// 解析
func Unmarshal(data []byte, value interface{}){
	json.Unmarshal(data, value)
}

// 写入
func Marshal(path string, value interface{}) []byte {
	js, _ := json.Marshal(value)
	fmt.Printf("Json: %s", js)

	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	file.Write(js)
	return js
}