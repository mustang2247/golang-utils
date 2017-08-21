package xmlhelper

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

// 解析xml
func Open(path string, value interface{})  {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err.Error())
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	err = xml.Unmarshal(data, value)
	if err != nil {
		fmt.Printf("%s", err.Error())
		panic(err)
	}
}

// 生成xml文件
func MakeXml(path string, value interface{})  {
	output, err := xml.MarshalIndent(value, " ", " ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	os.Stdout.Write([]byte(xml.Header))
	file, err := os.Create(path)
	file.Write([]byte(xml.Header))
	file.Write(output)
}