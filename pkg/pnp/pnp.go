package pnp

import (
	_ "embed"
	"math/rand"
	"time"

	"github.com/ronna-s/gc-eu-25/pkg/repo"
)

type (
	// Game represents a Platforms and Programmers Game
	// where a band of player will attempt to take on production
	Game struct {
		Score          int
		Players        []Player
		Prod           ProductionState
		BandName       string
		Coins          int
		CurrentPlayer  int
		ProductManager string
		Pizzas         int
	}

	Outcome string

	// Player represents a P&P player
	BasePlayer interface {
		PossibleActions(g *Game) []Action
	}

	MortalPlayer interface {
		Alive() bool
	}

	Player interface {
		BasePlayer
		MortalPlayer
	}

	// Engine represents the game's user interface rendering engine
	Engine interface {
		Welcome(leaderboard []repo.ScoreEntry, fn func(bandName string))
		GameOver()
		GameWon()
		RenderGame(g *Game)
		SelectAction(g *Game, player Player, cb func(selected Action))
		RenderOutcome(outcome Outcome, cb func())
		PizzaDelivery(cb func())
		WithOnExit(cb func()) Engine
	}
)

// New returns a new P&P game
func New(players ...Player) *Game {
	g := Game{Players: append(players, NewDwarf("Gimli"), NewNamelessPlayer(), NewMinion("Jurgen"), NewVibeCoder("R. Stallman")), Prod: NewProduction(), Coins: 15}
	return &g
}

// Run starts a new game
func (g *Game) Run(e Engine) {
	leaderboard, _ := repo.GetTop(10)
	e.Welcome(leaderboard, func(bandName string) {
		g.BandName = bandName
		e = e.WithOnExit(func() {
			err := repo.Persist(repo.ScoreEntry{Score: g.Score, BandName: g.BandName})
			if err != nil {
				panic(err)
			}
		})
		g.MainLoop(e)
	})
}

// MainLoop kicks off the next players round
func (g *Game) MainLoop(e Engine) {
	e.RenderGame(g)
	if allPlayersDead(g.Players) {
		e.GameOver()
	}

	if g.Coins > 100 {
		e.GameWon()
	}

	e.SelectAction(g, g.Players[g.CurrentPlayer], func(selected Action) {
		outcome := selected.Selected(g)
		e.RenderOutcome(outcome, func() {
			g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
			for !g.Players[g.CurrentPlayer].Alive() && !allPlayersDead(g.Players) {
				g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
			}

			if pr, ok := g.Players[g.CurrentPlayer].(interface {
				PizzaRequested() bool
				PizzaDelivered()
			}); ok {
				if pr.PizzaRequested() {
					// TODO: Look at the concurrency package to implement this
					// concurrency.Run(...)
					go func() {
						time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
						e.PizzaDelivery(func() {
							g.Pizzas += g.NumberOfPlayersAlive()
							for _, player := range g.Players {
								if v, ok := player.(interface{ Heal() }); ok {
									v.Heal() // Heal players that can heal
								}
							}
						})
						pr.PizzaDelivered()
					}()
				}
			}

			g.MainLoop(e)
		})
	})

}

func (g *Game) NumberOfPlayersAlive() int {
	result := 0
	for _, player := range g.Players {
		if player.Alive() {
			result++
		}
	}

	return result
}

func allPlayersDead(players []Player) bool {
	for _, player := range players {
		if !isMinion(player) && player.Alive() {
			return false
		}
	}
	return true
}

type Action struct {
	OnSelect    func(g *Game) Outcome
	Description string
}

func (o Action) String() string {
	return o.Description
}

func (o Action) Selected(g *Game) Outcome {
	return o.OnSelect(g)
}
