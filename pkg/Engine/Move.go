package Engine

type Move struct {
	pos1 Point		// Move Structure, a Liberty is a place where you 
	pos2 Point		// place a piece on the Go board.
}

func (m *Move) Pos1() Point {
	return m.pos1
}

func (m *Move) Pos2() Point {
	return m.pos2
}