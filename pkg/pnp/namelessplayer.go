package pnp

import "math/rand"

type NamelessPlayer struct {
	IsDead bool
}

func NewNamelessPlayer() *NamelessPlayer {
	return &NamelessPlayer{}
}

func (p *NamelessPlayer) PossibleActions(g *Game) []Action {
	return []Action{
		{
			Description: "Do nothing",
			OnSelect: func(g *Game) Outcome {
				if !p.IsDead {
					p.IsDead = rand.Intn(10) < 3 // Randomly set dead or alive
				}

				return "You did nothing"
			},
		},
	}
}

func (p *NamelessPlayer) Alive() bool {

	return !p.IsDead
}

func (p *NamelessPlayer) Heal() {
	// TODO: Implement
	p.IsDead = false
}

