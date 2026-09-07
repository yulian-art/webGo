package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 调用函数所需参数
type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// 返回结果
type addResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}


func main(){
	url := "http://127.0.0.1:9093/add"
	param := addParam{
		X: 10,
		Y: 20,
	}
	//解析参数
	paramBytes, _ := json.Marshal(param)
	// 发送请求
	resp, _ := http.Post(url, "appliaction/json", bytes.NewReader(paramBytes))
	defer resp.Body.Close()
	// 处理参数
	respBytes, _ := ioutil.ReadAll(resp.Body)
	var respData addResult
	json.Unmarshal(respBytes, &respData)
	fmt.Println(respData.Data)
}