package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func add(x, y int) int {
	return x + y
}

func addHandler(w http.ResponseWriter,r *http.Request) {
	// 解析参数
	b, _ := ioutil.ReadAll(r.Body)
	var param addParam
	json.Unmarshal(b, &param)
	// 业务逻辑
	ret := add(param.X, param.Y)
	// 返回相应
	respBytes, _ := json.Marshal(addResult{Code: 0, Data : ret})
	w.Write(respBytes)

}
func main() {
	http.HandleFunc("/add", addHandler)
	log.Fatal(http.ListenAndServe(":9093", nil))
}