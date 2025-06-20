package pnp

import "fmt"

type ProductManager struct {
	Fired bool
	pizzaRequested bool
}

func (p *ProductManager) PossibleActions(g *Game) []Action {
	return []Action{
		{
			Description: fmt.Sprintf("Pay wages (cost %d)", g.NumberOfPlayersAlive()),
			OnSelect: func(g *Game) Outcome {
				if g.Coins < g.NumberOfPlayersAlive() {
					p.Fired = true
					return "Not enough coins to pay wages. Band is bankrupt. PM is fired!"
				}
				g.Coins -= g.NumberOfPlayersAlive()
				return "Wages paid"
			},
		},
		{
			Description: fmt.Sprintf("Order Pizza (cost %d)", g.NumberOfPlayersAlive() * 2),
			OnSelect: func(g *Game) Outcome {
				if g.Coins < g.NumberOfPlayersAlive() {
					p.Fired = true
					return "Not enough coins to pay wages. Band is bankrupt. PM is fired!"
				}
				if g.Coins < g.NumberOfPlayersAlive() * 2 {
					g.Coins -= g.NumberOfPlayersAlive()
					return "Not enough coins to buy pizza , can only afford the wages"
				}
				g.Coins -= g.NumberOfPlayersAlive() * 2

				p.pizzaRequested = true

				return "Pizza's ordered, please wait for delivery"
			},
		},
	}
}

func (p *ProductManager) PizzaRequested() bool {
	return p.pizzaRequested
}

func (p *ProductManager) PizzaDelivered() {
	p.pizzaRequested = false
}

func (p *ProductManager) String() string {
	return "Sir Tan Lee Knot"
}

func (p *ProductManager) AsciiArt() string {
	return `
 O
/|\
/ \`
}

func (p *ProductManager) Alive() bool {
	return !p.Fired
}

func (p *ProductManager) Heal() {
	p.Fired = false
}
