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
				return "You did nothing"
			},
		},
	}
}

func (p *NamelessPlayer) Name() string {
	return "No Name"
}

func (p *NamelessPlayer) Alive() bool {
	if !p.IsDead {
		p.IsDead = rand.Intn(10) < 3 // Randomly set dead or alive
	}

	return !p.IsDead
}

