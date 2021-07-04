package dm04

import (
	"duel-masters/game/civ"
	"duel-masters/game/cnd"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
)

// ShadowMoonCursedShade ...
func ShadowMoonCursedShade(c *match.Card) {

	c.Name = "Shadow Moon, Cursed Shade"
	c.Power = 3000
	c.Civ = civ.Darkness
	c.Family = family.Ghost
	c.ManaCost = 4
	c.ManaRequirement = []string{civ.Darkness}

	c.Use(fx.Creature, func(card *match.Card, ctx *match.Context) {

		if card.Zone != match.BATTLEZONE {
			return
		}

		getDarknessCreatures(card, ctx).Map(func(x *match.Card) {
			x.AddUniqueSourceCondition(cnd.PowerAmplifier, 2000, card.ID)
		})
	})

}

func getDarknessCreatures(card *match.Card, ctx *match.Context) fx.CardCollection {

	darknessCreatures := fx.FindFilter(
		card.Player,
		match.BATTLEZONE,
		func(x *match.Card) bool { return x.Civ == civ.Darkness && x.ID != card.ID },
	)

	darknessCreatures = append(darknessCreatures,

		fx.FindFilter(
			ctx.Match.Opponent(card.Player),
			match.BATTLEZONE,
			func(x *match.Card) bool { return x.Civ == civ.Darkness && x.ID != card.ID },
		)...,
	)

	return darknessCreatures
}

// VolcanoSmogDeceptiveShade ...
func VolcanoSmogDeceptiveShade(c *match.Card) {

	c.Name = "Volcano Smog, Deceptive Shade"
	c.Power = 5000
	c.Civ = civ.Darkness
	c.Family = family.Ghost
	c.ManaCost = 6
	c.ManaRequirement = []string{civ.Darkness}

	c.Use(fx.Creature, func(card *match.Card, ctx *match.Context) {

		if card.Zone != match.BATTLEZONE {
			return
		}

		if event, ok := ctx.Event.(*match.PlayCardEvent); ok {

			p := ctx.Match.CurrentPlayer()

			playedCard, err := p.Player.GetCard(event.CardID, match.HAND)

			if err != nil {
				return
			}

			if playedCard.Civ != civ.Light {
				return
			}

			playedCard.AddUniqueSourceCondition(cnd.IncreasedCost, 2, card.ID)
		}
	})

}
