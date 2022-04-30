package vector

const (
	North Direction = "north"
	South           = "south"
	West            = "west"
	East            = "east"
	Right           = "right"
	Left            = "left"
	Up              = "up"
	Down            = "down"
)

type Direction string

type Vector struct {
	magnitude float64
	direction Direction
	angle     float64
}

func NewVector() Vector {
	return Vector{}
}

func Resolve(vecs []Vector) Vector {
	return Vector{}
}
