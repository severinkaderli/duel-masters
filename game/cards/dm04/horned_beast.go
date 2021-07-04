package dm04

import (
	"duel-masters/game/civ"
	"duel-masters/game/cnd"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
	"fmt"
)

// NiofaHornedProtector ...
func NiofaHornedProtector(c *match.Card) {

	c.Name = "Niofa, Horned Protector"
	c.Power = 9000
	c.Civ = civ.Nature
	c.Family = family.HornedBeast
	c.ManaCost = 6
	c.ManaRequirement = []string{civ.Nature}


	c.Use(fx.Creature, fx.Evolution, fx.Doublebreaker, fx.When(fx.Summoned, func(card *match.Card, ctx *match.Context) {

		cards := match.Filter(
			card.Player,
			ctx.Match,
			card.Player,
			match.DECK,
			"Select 1 nature creature from your deck that will be shown to your opponent and sent to your hand",
			0,
			1,
			false,
			func(x *match.Card) bool { return x.HasCondition(cnd.Creature) && x.Civ == civ.Nature },
		)

		for _, c := range cards {
			card.Player.MoveCard(c.ID, match.DECK, match.HAND)
			ctx.Match.Chat("Server", fmt.Sprintf("%s was moved from %s's deck to their hand", c.Name, card.Player.Username()))
		}

		card.Player.ShuffleDeck()

	}))

}
