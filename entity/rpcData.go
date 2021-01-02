package entity

type RpcBytes struct {
	Id     int         `json:"id"`
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
}
