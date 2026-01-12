package domain

func (g *gameSession) GetNextApologiseMove() vec {
	v := vec{0, 0}

	if g.base.field[1][1] == 0 {
		v = vec{1, 1}
	}
	return v
}
