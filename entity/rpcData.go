package entity

import "encoding/json"

type RpcBytes struct {
	Result json.RawMessage `json:"result"`
	Error error`json:"error"`
	Id int `json:"id"`


}
