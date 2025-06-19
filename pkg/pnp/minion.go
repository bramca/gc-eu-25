package pnp

import (
	_ "embed"
	"fmt"
	"math/rand"

	"github.com/ronna-s/gc-eu-25/pkg/maybe"
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
			Description: "Add a feature to the code",
			OnSelect: func(g *Game) Outcome {
				addScore := -50 + rand.Intn(200)
				g.Score += addScore

				addedValue := ""
				if addScore > 50 {
					addCoins := rand.Intn(20)
					g.Coins += addCoins
					addedValue = fmt.Sprintf("=> feature added a lot of value => +%d coins ", addCoins)
				}

				if addScore < 0 {
					costCoins := 1 + rand.Intn(5)
					g.Coins -= costCoins
					addedValue = fmt.Sprintf("=> feature was actually a bug => -%d coins ", costCoins)
				}

				return Outcome(fmt.Sprintf("Feature added! Score %s %s=> %s", maybe.If[string](addScore > -1).Then(fmt.Sprintf("+%d", addScore)).Else(fmt.Sprintf("%d", addScore)), addedValue, maybe.This(g.Prod.Upset()).If(addScore < 0).Else(g.Prod.CalmDown())))
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
