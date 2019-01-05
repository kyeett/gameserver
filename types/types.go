package types

type Coord struct{ X, Y int }

func (c Coord) Add(d Coord) Coord {
	return Coord{
		X: c.X + d.X,
		Y: c.Y + d.Y,
	}
}

type Position struct {
	Coord
	Theta int
}

type Tile int

const (
	Invalid Tile = iota + 1
	Water
	Grass
	GrassUp = 12 + Water
)

type World struct {
	tiles         []Tile
	Width, Height int
}

var FirstWorld = World{
	[]Tile{
		Water, Water, Water, Water, Water, Water, Water, Water, Water, Water,
		Water, Grass, Grass, Grass, Water, Water, Grass, Grass, Water, Water,
		Water, Grass, Grass, Grass, Grass, Water, Water, Grass, Grass, Water,
		Water, Grass, Grass, Grass, Grass, Grass, Water, Grass, Grass, Water,
		Water, Water, Water, Water, Grass, Grass, Water, Grass, Grass, Water,
		Water, Water, Grass, Grass, Grass, Water, Water, Grass, Grass, Water,
		Water, Grass, Grass, Grass, Water, Water, Grass, Grass, Grass, Water,
		Water, Grass, Grass, Grass, Grass, Grass, Grass, Grass, Grass, Water,
		Water, Water, Grass, Grass, Grass, Grass, Grass, Grass, Grass, Water,
		Water, Water, Water, Water, Water, Water, Water, Water, Water, Water,
	},
	10, 10,
}

func (w *World) ValidTarget(t Position) bool {
	if t.X < 0 || t.X >= w.Width || t.Y < 0 || t.Y >= w.Height {
		return false
	}

	return w.tiles[t.Y*w.Width+t.X] == Grass
}

func (w World) At(p Coord) Tile {
	if p.X < 0 || p.X >= w.Width || p.Y < 0 || p.Y >= w.Height {
		return Invalid
	}
	return w.tiles[p.Y*w.Width+p.X]
}

func (w World) Set(p Coord, t Tile) {
	if p.X < 0 || p.X >= w.Width || p.Y < 0 || p.Y >= w.Height {
		return
	}

	w.tiles[p.Y*w.Width+p.X] = t
}
