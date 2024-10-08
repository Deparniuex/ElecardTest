package entity

type Circle struct {
	X int `json:"x"`
	Y int `json:"y"`
	R int `json:"radius"`
}

type Rectangle struct {
	LeftBottom  [2]int `json:"left_bottom"`
	RightBottom [2]int `json:"right_bottom"`
}
