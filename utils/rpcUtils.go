package utils

import (
	"BitcoinWeb/entity"
	"BitcoinWeb/service"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var DataCount entity.Range
func PrepareJSON(method string,params ...interface{}) interface{} {
	rpcReq:=entity.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  method,
		Jsonrpc: RPCVERSION,
		Params:  params,
	}
	reqBytes,err:=json.Marshal(&rpcReq)
	if err!=nil {
		return err.Error()
	}
	//fmt.Println("准备好的json数据为：",string(reqBytes))

	client:=&http.Client{}
	request,err:=http.NewRequest("POST",RPCURL,bytes.NewBuffer(reqBytes))
	if err!=nil {
		return err.Error()
	}
	authStr:=RPCUSER+":"+RPCPASSWORD

	request.Header.Add("Encoding","UTF-8")
	request.Header.Add("Content-Type","application/json")
	request.Header.Add("Authorization","Basic "+service.Base64Str(authStr))
	
	response,err:=client.Do(request)
	if err!=nil {
		return err.Error()
	}
	//var Bytes  entity.RpcBytes
	Bytes,err:=ioutil.ReadAll(response.Body)
	if err!=nil {
		return err.Error()
	}
	ret:=service.Unmarshal(Bytes)
	//fmt.Println("反序列化：",string(ret.Result))
	//fmt.Println("反序列化：",service.UnmarshalBlock(ret.Result))
	dataStr:= strings.Split(string(ret.Result),",")


	//fmt.Println(dataStr)
	//for i:=0;i<=len(dataStr) ;i++  {
	//	// fmt.Println(dataStr[i])
	//	 DataCount.Ranges=append(DataCount.Ranges,dataStr[i])
	//	 //fmt.Println(string(DataCount.Result))
	//}





	defer response.Body.Close()
	//fmt.Println(string(Bytes))

	code:=response.StatusCode
	if code==200 {
		return dataStr
	}else {
		return "请求失败！"
	}






}



















