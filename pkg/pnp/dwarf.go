package pnp

import (
	"fmt"
	"math/rand"
)

type Dwarf struct {
	Name string
	GoldProbability int
	PickaxeDurability int
}

func NewDwarf(name string) *Dwarf {
	return &Dwarf{
		Name: name,
		// Dwarf starts with a durable pickaxe
		PickaxeDurability: 2 + rand.Intn(8),
		GoldProbability: 30,
	}
}

func (d *Dwarf) PossibleActions(g *Game) []Action {
	return []Action{
		{
			Description: "Dig a tunnel",
			OnSelect: func(g *Game) Outcome {
				d.PickaxeDurability -= 1 // Using the pickaxe reduces its durability
				if g.Coins < 2 {
					return "Not enough coins to dig a tunnel"
				}
				g.Coins -= 2

				probIncrease := rand.Intn(10)
				d.GoldProbability += probIncrease
				return Outcome(fmt.Sprintf("You dug a tunnel => probability of finding gold increased +%d%%", probIncrease))
			},
		},
		{
			Description: "Mine for gold",
			OnSelect: func(g *Game) Outcome {
				d.PickaxeDurability -= 1 // Using the pickaxe reduces its durability
				if rand.Intn(100) < d.GoldProbability {
					g.Coins += 5

					return "You found 5 gold coins!"
				}
				return "No gold found this time"
			},
		},
	}
}

func (d *Dwarf) String() string {
	return fmt.Sprintf("Dwarf (%d 󰢷, %d 󰴯): %s", d.PickaxeDurability, d.GoldProbability, d.Name)
}

func (d *Dwarf) Alive() bool {
	return d.PickaxeDurability > 0
}

func (d *Dwarf) Heal() {
	d.PickaxeDurability = rand.Intn(10) // Repair the pickaxe to full durability
}
