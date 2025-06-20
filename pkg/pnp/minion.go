package pnp

import (
	_ "embed"
	"fmt"
	"math/rand"

	"github.com/ronna-s/gc-eu-25/pkg/maybe"
)

func NewMinion(name string) *Minion {
	return &Minion{
		Name: name,
		Skill: 150,
	}
}

type Minion struct {
	ImmortalPlayer
	Name string
	Skill int
}

//go:embed resources/minion.txt
var minionArt string

func (m *Minion) AsciiArt() string {
	return minionArt
}

func (m *Minion) PossibleActions(g *Game) []Action {
	var actions []Action
	if g.Coins > 0 {
		actions = append(actions, Action{
			Description: "Buy a banana and eat it (costs 1 gold coin)",
			OnSelect: func(g *Game) Outcome {
				g.Coins--
				skill := rand.Intn(50)
				m.Skill += skill
				return Outcome(fmt.Sprintf("You ate a banana => +%d skill", skill))
			},
		})
	}
	actions = append(actions,
		Action{
			Description: "Add a feature to the code",
			OnSelect: func(g *Game) Outcome {
				addScore := -m.Skill/3 + rand.Intn(m.Skill)
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

				scoreStr := maybe.If[string](addScore > -1).Then(fmt.Sprintf("+%d", addScore)).Else(fmt.Sprintf("%d", addScore))
				prodReaction := maybe.This(g.Prod.Upset).If(addScore < 0).Or(maybe.This(g.Prod.NoImpact).If(addScore == 0)).Else(g.Prod.CalmDown)

				return Outcome(fmt.Sprintf("Feature added! Score %s %s=> %s", scoreStr, addedValue, prodReaction()))
			},
		},
	)
	return actions
}

func (m *Minion) String() string {
	return fmt.Sprintf("Minion (%d skill): %s", m.Skill, m.Name)
}

func (m *Minion) IsMinion() bool {
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
