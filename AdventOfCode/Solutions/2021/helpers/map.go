package helpers

type direction struct {
	X int
	Y int
}

var UpDownLeftRight = []direction{
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
}
