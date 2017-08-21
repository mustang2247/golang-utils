package logutil

import "fmt"

/*检查错误信息*/
func CheckErr(err error) bool {
	if err != nil {
		fmt.Println("mysql checkErr: " + err.Error())
		panic(err)
		return true
	}
	return false
}