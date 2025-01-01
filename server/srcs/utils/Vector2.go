package utils

type Vector2 struct {
	X int
	Y int
}

func (this *Vector2) Substract(vec *Vector2) *Vector2 {
	return &Vector2{
		X: this.X - vec.X,
		Y: this.Y - vec.Y,
	}
}
