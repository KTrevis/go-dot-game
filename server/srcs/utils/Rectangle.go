package utils

type Rectangle struct {
	a	Vector2i
	b	Vector2i
}

// 0,0|1,0|2,0|3,0|4,0|5,0|
// 0,1|1,1|2,1|3,1|4,1|5,1|
// 0,2|1,2|2,2|3,2|4,2|5,2|
// 0,3|1,3|2,3|3,3|4,3|5,3|
// 0,4|1,4|2,4|3,4|4,4|5,4|
// 0,5|1,5|2,5|3,5|4,5|5,5|

func (this *Rectangle) IsInRect(pos *Vector2i) bool {
	if pos.X < this.a.X || pos.Y < this.a.Y {
		return false
	}

	if pos.X > this.b.X || pos.Y > this.b.Y {
		return false
	}

	return true
}
