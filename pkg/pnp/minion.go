package pnp

import (
	_ "embed"
	"math/rand"
)

func NewMinion(name string) Minion {
	return Minion{
		Name: name,
	}
}

type Minion struct {
	ImmortalPlayer
	Name string
}

//go:embed resources/minion.txt
var minionArt string

func (m Minion) AsciiArt() string {
	return minionArt
}

func (m Minion) PossibleActions(g *Game) []Action {
	var actions []Action
	if g.Coins > 0 {
		actions = append(actions, Action{
			Description: "Buy a banana and eat it (costs 1 gold coin)",
			OnSelect: func(g *Game) Outcome {
				g.Coins--
				return "You ate a banana"
			},
		})
	}
	actions = append(actions,
		Action{
			Description: "Add a bug to the code",
			OnSelect: func(g *Game) Outcome {
				return Outcome(g.Prod.Upset())
			},
		},
		Action{
			Description: "Add a feature to the code",
			OnSelect: func(g *Game) Outcome {
				g.Score += rand.Intn(100)
				return Outcome(g.Prod.CalmDown())
			},
		},
	)
	return actions
}

func (m Minion) String() string {
	return "Minion: " + m.Name
}

func (m Minion) IsMinion() bool {
	return true
}

type minionPlayer interface {
	IsMinion() bool
}

func isMinion(p Player) bool {
	if mp, ok := p.(minionPlayer); ok {
		return mp.IsMinion()
	}

	return false
}
