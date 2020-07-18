/**********************************************
** @Des:
** @Author: 1@lg1024.com
** @Last Modified time: 2020/3/3 上午11:52
***********************************************/

package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

type ResponseModel struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data VerifyResponseModel
}
type VerifyResponseModel struct {
	Rmb    string `json:"rmb"`
	Userid string `json:"userid"`
	Num    string `json:"num"`
	Ext2   string `json:"ext2"`
}

func main() {
	extra.RegisterFuzzyDecoders()

	//str := `{"a": 10, "b": 10}`
	str := `{
    "ret":"200",
    "msg":"success",
    "data":{
        "rmb":"100",
        "num":1,
        "userid":"783582589",
        "ext2":"u8_order_id=2073548259904716800&productName=关羽大刀&gameid=1233344891086442496&coopid=1233236139087761408"
   `
	var a ResponseModel
	err := jsoniter.UnmarshalFromString(str, &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", a)
}
