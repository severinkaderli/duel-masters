package dm02

import (
	"duel-masters/game/civ"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
)

// XenoMantis ...
func XenoMantis(c *match.Card) {

	c.Name = "Xeno Mantis"
	c.Power = 6000
	c.Civ = civ.Nature
	c.Family = family.GiantInsect
	c.ManaCost = 7
	c.ManaRequirement = []string{civ.Nature}

	c.Use(func(card *match.Card, ctx *match.Context) {

		if event, ok := ctx.Event.(*match.AttackPlayer); ok {

			if event.CardID != card.ID {
				return
			}

			ctx.ScheduleAfter(func() {

				blockers := make([]*match.Card, 0)

				for _, blocker := range event.Blockers {
					if ctx.Match.GetPower(blocker, false) > 5000 {
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
					if ctx.Match.GetPower(blocker, false) > 5000 {
						blockers = append(blockers, blocker)
					}
				}

				event.Blockers = blockers

			})

		}

	}, fx.Creature, fx.Doublebreaker)

}
