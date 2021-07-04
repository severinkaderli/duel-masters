package dm02

import (
	"duel-masters/game/civ"
	"duel-masters/game/cnd"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
	"fmt"
)

// RumblingTerahorn ...
func RumblingTerahorn(c *match.Card) {

	c.Name = "Rumbling Terahorn"
	c.Power = 3000
	c.Civ = civ.Nature
	c.Family = family.HornedBeast
	c.ManaCost = 5
	c.ManaRequirement = []string{civ.Nature}

	c.Use(fx.Creature, func(card *match.Card, ctx *match.Context) {

		if match.AmISummoned(card, ctx) {

			cards := match.SearchForCnd(card.Player, ctx.Match, card.Player, match.DECK, cnd.Creature, "Select 1 creature from your deck that will be shown to your opponent and sent to your hand", 1, 1, true)

			for _, c := range cards {
				card.Player.MoveCard(c.ID, match.DECK, match.HAND)
				ctx.Match.Chat("Server", fmt.Sprintf("%s was moved from %s's deck to their hand", c.Name, card.Player.Username()))
			}

			card.Player.ShuffleDeck()

		}

	})

}

// LeapingTornadoHorn ...
func LeapingTornadoHorn(c *match.Card) {

	c.Name = "Leaping Tornado Horn"
	c.Power = 2000
	c.Civ = civ.Nature
	c.Family = family.HornedBeast
	c.ManaCost = 3
	c.ManaRequirement = []string{civ.Nature}

	c.PowerModifier = func(m *match.Match, attacking bool) int {

		power := 0

		if (attacking) {
			for _, creature := range fx.Find(c.Player, match.BATTLEZONE) {
				if creature == c {
					continue
				}
				power += 1000
			}	
		}
	
		return power
	}

	c.Use(fx.Creature)

}
