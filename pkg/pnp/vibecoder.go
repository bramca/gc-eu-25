package pnp

import (
	"fmt"
	"math/rand"

	"github.com/ronna-s/gc-eu-25/pkg/maybe"
)

type VibeCoder struct {
	name string
	Fired bool
	Contribution int
}

func NewVibeCoder(name string) *VibeCoder {
	return &VibeCoder{
		name: name,
	}
}

func (p *VibeCoder) PossibleActions(g *Game) []Action {
	return []Action{
		{
			Description: "Do nothing",
			OnSelect: func(g *Game) Outcome {
				p.Contribution -= 1
				g.Score += 5
				if p.Contribution <= -10 {
					p.Fired = true

					return "Your contribution is too low, you are fired!"
				}

				return "You did nothing => +5 Score"
			},
		},
		{
			Description: "Ask ChatGPT",
			OnSelect: func(g *Game) Outcome {
				contribution := -10 + rand.Intn(25)
				cost := -1 * rand.Intn(5)
				p.Contribution += contribution
				if g.Coins + cost < 0 {
					p.Fired = true

					return "You're costing to much, you are fired!"
				}
				if p.Contribution <= -20 {
					p.Fired = true

					return "Your contribution is too low, you are fired!"
				}
				g.Coins += cost
				g.Score += contribution

				prodReaction := maybe.If[func() string](contribution > 9).Then(g.Prod.CalmDown).Or(maybe.This(g.Prod.NoImpact).If(contribution > -1)).Else(g.Prod.Upset)
				coinsStr := maybe.If[string](contribution > -1).Then(fmt.Sprintf("+%d", contribution)).Else(fmt.Sprintf("%d", contribution))

				return Outcome(fmt.Sprintf("AI did something => cost %d coins => %s score => %s", cost, coinsStr, prodReaction()))
			},
		},
	}
}

func (p *VibeCoder) Name() string {
	return p.name
}

func (p *VibeCoder) Alive() bool {
	return !p.Fired
}

func (p *VibeCoder) Heal() {
	// TODO: implement
	p.Fired = false
}
