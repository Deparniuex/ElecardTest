package entity

type Body struct {
	Key    string      `json:"key"`
	Method string      `json:"method"`
	Params []Rectangle `json:"params"`
}
