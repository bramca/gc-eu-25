package pnp

type ProductManager struct {
	Fired bool
}

func (p *ProductManager) PossibleActions(g *Game) []Action {
	return []Action{
		{
			Description: "Pay wages",
			OnSelect: func(g *Game) Outcome {
				if g.Coins < g.NumberOfPlayersAlive() {
					p.Fired = true
					return "Not enough coins to pay wages. Band is bankrupt. PM is fired!"
				}
				g.Coins -= g.NumberOfPlayersAlive()
				return "Wages paid"
			},
		},
		// TODO: Add action that starts a goroutine that
		// sleeps a random amount of time and then delivers a pizza
		/* wrap in goroutine
		time.Sleep(rand.Intn(20)*time.Second)
		e.PizzaDelivery(func() {
			for _, player := range g.Players {
				if v, ok := player.(interface{ Heal() }); ok {
					v.Heal() // Heal players that can heal
				}
			}
		})
		*/
	}
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
