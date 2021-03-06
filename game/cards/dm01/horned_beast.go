package dm01

import (
	"duel-masters/game/civ"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
)

// RoaringGreatHorn ...
func RoaringGreatHorn(c *match.Card) {

	c.Name = "Roaring Great-Horn"
	c.Power = 6000
	c.Civ = civ.Nature
	c.Family = family.HornedBeast
	c.ManaCost = 7
	c.ManaRequirement = []string{civ.Nature}

	c.Use(fx.Creature, fx.Doublebreaker, fx.PowerAttacker2000)

}

// StampedingLonghorn ...
func StampedingLonghorn(c *match.Card) {

	c.Name = "Stampeding Longhorn"
	c.Power = 4000
	c.Civ = civ.Nature
	c.Family = family.HornedBeast
	c.ManaCost = 5
	c.ManaRequirement = []string{civ.Nature}

	c.Use(func(card *match.Card, ctx *match.Context) {

		if event, ok := ctx.Event.(*match.AttackPlayer); ok {

			if event.CardID != card.ID {
				return
			}

			ctx.ScheduleAfter(func() {

				blockers := make([]*match.Card, 0)

				for _, blocker := range event.Blockers {
					if blocker.Power >= 3000 {
						blockers = append(blockers, blocker)
					}
				}

				event.Blockers = blockers

			})

		}

		if event, ok := ctx.Event.(*match.AttackCreature); ok {

			if event.CardID != card.ID {
				return
			}

			ctx.ScheduleAfter(func() {

				blockers := make([]*match.Card, 0)

				for _, blocker := range event.Blockers {
					if blocker.Power >= 3000 {
						blockers = append(blockers, blocker)
					}
				}

				event.Blockers = blockers

			})

		}

	}, fx.Creature)

}

// TriHornShepherd ...
func TriHornShepherd(c *match.Card) {

	c.Name = "Tri-Horn Shepherd"
	c.Power = 5000
	c.Civ = civ.Nature
	c.Family = family.HornedBeast
	c.ManaCost = 5
	c.ManaRequirement = []string{civ.Nature}

	c.Use(fx.Creature)

}
