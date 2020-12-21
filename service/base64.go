package service

import (
	"BitcoinWeb/entity"
	"BitcoinWeb/structs"
	"CryptoHashCodeClass3006/CryptoHashCodeClass3/base"
	"encoding/json"
	"fmt"
)

func Base64Str(data string) string {
	return string(base.Base64Encode([]byte(data)))
}
func Unmarshal(body []byte) *entity.RpcBytes {
	RPCResult := new(entity.RpcBytes)
	err := json.Unmarshal(body,RPCResult)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return RPCResult
}
func UnmarshalBlock(body []byte) *structs.GetBlock {
	RPCResult := new(structs.GetBlock)
	err := json.Unmarshal(body,RPCResult)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return RPCResult
}
