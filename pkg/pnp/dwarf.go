package pnp

import (
	"fmt"
	"math/rand"
)

type Dwarf struct {
	Name string
	PickaxeDurability int
}

func NewDwarf(name string) *Dwarf {
	return &Dwarf{
		Name: name,
		PickaxeDurability: 10, // Dwarf starts with a durable pickaxe
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
				return "You dug a tunnel"
			},
		},
		{
			Description: "Mine for gold",
			OnSelect: func(g *Game) Outcome {
				d.PickaxeDurability -= 1 // Using the pickaxe reduces its durability
				if rand.Intn(10) < 3 { // 30% chance of finding gold
					g.Coins += 5

					return "You found 5 gold coins!"
				}
				return "No gold found this time"
			},
		},
	}
}

func (d *Dwarf) String() string {
	return fmt.Sprintf("Dwarf (%d 󰢷): %s", d.PickaxeDurability, d.Name)
}

func (d *Dwarf) Alive() bool {
	return d.PickaxeDurability > 0
}
