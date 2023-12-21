package types

type Point struct {
	X, Y int64
}

type Pipe struct {
	Tile
	Point
	Symbol string
}

func (t Pipe) String() string {
	return t.Symbol
}

func (t Pipe) Ok() bool {
	return t.Tile != Ground
}

func (t Pipe) Allowed(dir Direction) bool {
	switch t.Tile {
	case NorthSouth:
		switch dir {
		case North, South:
			return true
		default:
			return false
		}
	case EastWest:
		switch dir {
		case East, West:
			return true
		default:
			return false
		}
	case NorthEast:
		switch dir {
		case North, East:
			return true
		default:
			return false
		}
	case NorthWest:
		switch dir {
		case North, West:
			return true
		default:
			return false
		}
	case SouthWest:
		switch dir {
		case South, West:
			return true
		default:
			return false
		}
	case SouthEast:
		switch dir {
		case South, East:
			return true
		default:
			return false
		}
	default:
		return false
	}
}

func (t Pipe) Next(dir Direction) Direction {
	if t.Tile == EastWest && dir == East {
		return East
	}
	if t.Tile == EastWest && dir == West {
		return West
	}
	if t.Tile == NorthSouth && dir == North {
		return North
	}
	if t.Tile == NorthSouth && dir == South {
		return South
	}
	if t.Tile == NorthEast && dir == South {
		return East
	}
	if t.Tile == NorthEast && dir == West {
		return North
	}
	if t.Tile == NorthWest && dir == South {
		return West
	}
	if t.Tile == NorthWest && dir == East {
		return North
	}
	if t.Tile == SouthWest && dir == North {
		return West
	}
	if t.Tile == SouthWest && dir == East {
		return South
	}
	if t.Tile == SouthEast && dir == North {
		return East
	}
	if t.Tile == SouthEast && dir == West {
		return South
	}
	panic("unexpected tile")
}

func FromStr(ch string) Tile {
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

func NewPipe(ch string, p Point) Pipe {
	tile := FromStr(ch)
	return Pipe{
		Point:  p,
		Tile:   tile,
		Symbol: ch,
	}
}
