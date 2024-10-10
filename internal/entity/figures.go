package entity

type Circle struct {
	X int `json:"x"`
	Y int `json:"y"`
	R int `json:"radius"`
}

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Rectangle struct {
	LeftBottom Coordinate `json:"left_bottom"`
	RightTop   Coordinate `json:"right_top"`
}

type Tasks struct {
	Result [][]Circle `json:"result"`
}
