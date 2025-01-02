package utils

type Vector2i struct {
	X int
	Y int
}

func (this *Vector2i) Substract(vec *Vector2i) *Vector2i {
	return &Vector2i{
		X: this.X - vec.X,
		Y: this.Y - vec.Y,
	}
}
