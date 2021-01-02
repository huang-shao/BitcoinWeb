package utils

import (
	"BitcoinWeb/entity"
	"BitcoinWeb/service"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func PrepareJSON(method string, params ...interface{}) interface{} {
	rpcReq := entity.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  method,
		Jsonrpc: RPCVERSION,
		Params:  params,
	}
	reqBytes, err := json.Marshal(&rpcReq)
	fmt.Println("序列化后的内容：", string(reqBytes))
	if err != nil {
		fmt.Println("出现错误1:", err.Error())
		return err.Error()
	}
	//fmt.Println("准备好的json数据为：",string(reqBytes))

	client := &http.Client{}
	request, err := http.NewRequest("POST", RPCURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println("出现错误2:", err.Error())
		return err.Error()
	}
	authStr := RPCUSER + ":" + RPCPASSWORD

	request.Header.Add("Encoding", "UTF-8")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Basic "+service.Base64Str(authStr))

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("出现错误3:", err.Error())
		return err.Error()
	}
	//var Bytes  entity.RpcBytes
	Bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("出现错误4:", err.Error())
		return err.Error()
	}
	ret := service.Unmarshal(Bytes)

	//fmt.Println("反序列化：",string(ret.Result))
	//fmt.Println("反序列化：",service.UnmarshalBlock(ret.Result))
	//dataStr:= strings.Split(string(ret.Result),",")

	//fmt.Println(dataStr)
	//for i:=0;i<=len(dataStr) ;i++  {
	//	// fmt.Println(dataStr[i])
	//	 DataCount.Ranges=append(DataCount.Ranges,dataStr[i])
	//	 //fmt.Println(string(DataCount.Result))
	//}

	defer response.Body.Close()
	//fmt.Println(string(Bytes))

	code := response.StatusCode
	fmt.Println("状态码:", code)
	if code == 200 {
		//fmt.Println("请求有问题：", ret)
		return ret.Result
	} else {
		return "请求失败！"
	}
}
