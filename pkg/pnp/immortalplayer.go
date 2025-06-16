package pnp

type ImmortalPlayer struct {
	BasePlayer
}

func NewImmortalPlayer(p BasePlayer) Player {
	return ImmortalPlayer{
		BasePlayer: p,
	}
}

func (p ImmortalPlayer) Alive() bool {
	return true
}
