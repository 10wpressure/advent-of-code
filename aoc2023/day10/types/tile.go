package types

type Tile int

const (
	Ground Tile = iota
	NorthSouth
	EastWest
	NorthEast
	NorthWest
	SouthWest
	SouthEast
	Start
)

func ParseTile(ch string) Tile {
	switch ch {
	case "|":
		return NorthSouth
	case "-":
		return EastWest
	case "J":
		return NorthWest
	case "L":
		return NorthEast
	case "7":
		return SouthWest
	case "F":
		return SouthEast
	case "S":
		return Start
	default:
		return Ground
	}
}

func (t Tile) String() string {
	switch t {
	case NorthSouth:
		return "|"
	case EastWest:
		return "-"
	case NorthEast:
		return "J"
	case NorthWest:
		return "L"
	case SouthWest:
		return "7"
	case SouthEast:
		return "F"
	case Start:
		return "S"
	default:
		return "."
	}
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "N"
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	default:
		return ""
	}
}

func (d Direction) Delta() Point {
	switch d {
	case North:
		return Point{0, -1}
	case East:
		return Point{1, 0}
	case South:
		return Point{0, 1}
	case West:
		return Point{-1, 0}
	default:
		return Point{0, 0}
	}
}
