package grid

type XYZ struct {
	X, Y, Z int
}

func (xyz XYZ) Add(other XYZ) XYZ {
	xyz.X += other.X
	xyz.Y += other.Y
	xyz.Z += other.Z
	return xyz
}

var ZYXSides = []XYZ{
	{X: -1, Y: 0, Z: 0},
	{X: 1, Y: 0, Z: 0},
	{X: 0, Y: -1, Z: 0},
	{X: 0, Y: 1, Z: 0},
	{X: 0, Y: 0, Z: -1},
	{X: 0, Y: 0, Z: 1},
}
