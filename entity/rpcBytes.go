package entity

type RpcCommand struct {
	Command          string `form:"command"`
	First_parameter  string `form:"first_parameter"`
	Second_parameter string `form:"second_parameter"`
}
