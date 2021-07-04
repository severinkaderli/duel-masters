package fx

import "duel-masters/game/match"

// CardCollection is a slice of cards with a mapping function
type CardCollection []*match.Card

// Map iterates through cards in the collection and executes the function on them
func (c CardCollection) Map(h func(*match.Card)) {

	for _, card := range c {
		h(card)
	}

}

// FindFilter returns a CardCollection matching the filter
func FindFilter(p *match.Player, collection string, h func(card *match.Card) bool) CardCollection {

	result := CardCollection{}

	container, err := p.Container(collection)

	if err != nil {
		return result
	}

	for _, card := range container {
		if h(card) {
			result = append(result, card)
		}
	}

	return result

}

// Find returns a CardCollection for the specified container
func Find(p *match.Player, collection string) CardCollection {
	return FindFilter(p, collection, func(x *match.Card) bool { return true })
}

// When performs the specified function if the test is successful
func When(test func(*match.Card, *match.Context) bool, h func(*match.Card, *match.Context)) func(*match.Card, *match.Context) {

	return func(card *match.Card, ctx *match.Context) {
		if test(card, ctx) {
			h(card, ctx)
		}
	}

}

// Select prompts the user to select n cards from the specified container
func Select(p *match.Player, m *match.Match, containerOwner *match.Player, containerName string, text string, min int, max int, cancellable bool) CardCollection {
	return SelectFilter(p, m, containerOwner, containerName, text, min, max, cancellable, func(x *match.Card) bool { return true })
}

// SelectFilter prompts the user to select n cards from the specified container that matches the given filter
func SelectFilter(p *match.Player, m *match.Match, containerOwner *match.Player, containerName string, text string, min int, max int, cancellable bool, filter func(*match.Card) bool) CardCollection {

	result := make([]*match.Card, 0)

	cards, err := containerOwner.Container(containerName)

	if err != nil || len(cards) < 1 {
		return result
	}

	filtered := make([]*match.Card, 0)

	for _, mCard := range cards {
		if filter(mCard) {
			filtered = append(filtered, mCard)
		}
	}

	if len(filtered) < 1 {
		return result
	}

	m.NewAction(p, filtered, min, max, text, cancellable)

	defer m.CloseAction(p)

	for {

		action := <-p.Action

		if cancellable && action.Cancel {
			break
		}

		if len(action.Cards) < min || len(action.Cards) > max || !match.AssertCardsIn(filtered, action.Cards...) {
			m.ActionWarning(p, "The cards you selected does not meet the requirements")
			continue
		}

		for _, c := range action.Cards {

			selectedCard, err := containerOwner.GetCard(c, containerName)

			if err != nil {
				continue
			}

			result = append(result, selectedCard)

		}

		break

	}

	return result

}

// SelectMultipart prompts the user to select n cards from the specified list of cards
func SelectMultipart(p *match.Player, m *match.Match, cards map[string][]*match.Card, text string, min int, max int, cancellable bool) CardCollection {

	result := make([]*match.Card, 0)

	notEmpty := false

	for _, cardList := range cards {
		if len(cardList) > 0 {
			notEmpty = true
		}
	}

	if notEmpty != true {
		return result
	}

	m.NewMultipartAction(p, cards, min, max, text, cancellable)

	defer m.CloseAction(p)

	for {

		action := <-p.Action

		if cancellable && action.Cancel {
			break
		}

		if len(action.Cards) < min || len(action.Cards) > max {
			m.ActionWarning(p, "The cards you selected does not meet the requirements")
			continue
		}

		for _, vid := range action.Cards {

			for _, cardList := range cards {
				for _, card := range cardList {
					if card.ID == vid {
						result = append(result, card)
					}
				}
			}

		}

		break

	}

	return result

}

// SelectBackside prompts the user to select n cards from the specified container
func SelectBackside(p *match.Player, m *match.Match, containerOwner *match.Player, containerName string, text string, min int, max int, cancellable bool) CardCollection {
	return SelectBacksideFilter(p, m, containerOwner, containerName, text, min, max, cancellable, func(x *match.Card) bool { return true })
}

// SelectBacksideFilter prompts the user to select n cards from the specified container that matches the given filter
func SelectBacksideFilter(p *match.Player, m *match.Match, containerOwner *match.Player, containerName string, text string, min int, max int, cancellable bool, filter func(*match.Card) bool) CardCollection {

	result := make([]*match.Card, 0)

	cards, err := containerOwner.Container(containerName)

	if err != nil || len(cards) < 1 {
		return result
	}

	filtered := make([]*match.Card, 0)

	for _, mCard := range cards {
		if filter(mCard) {
			filtered = append(filtered, mCard)
		}
	}

	if len(filtered) < 1 {
		return result
	}

	m.NewBacksideAction(p, filtered, min, max, text, cancellable)

	defer m.CloseAction(p)

	for {

		action := <-p.Action

		if cancellable && action.Cancel {
			break
		}

		if len(action.Cards) < min || len(action.Cards) > max || !match.AssertCardsIn(filtered, action.Cards...) {
			m.ActionWarning(p, "The cards you selected does not meet the requirements")
			continue
		}

		for _, c := range action.Cards {

			selectedCard, err := containerOwner.GetCard(c, containerName)

			if err != nil {
				continue
			}

			result = append(result, selectedCard)

		}

		break

	}

	return result

}

// Hooks below:
// hooks are shorthands for checking if the context matches a certain condition

// Summoned returns true if the card was summoned
func Summoned(card *match.Card, ctx *match.Context) bool {
	if event, ok := ctx.Event.(*match.CardMoved); ok {

		if event.CardID == card.ID && event.To == match.BATTLEZONE {
			return true
		}

	}

	return false
}

// SpellCast returns true if the spell was cast
func SpellCast(card *match.Card, ctx *match.Context) bool {

	if event, ok := ctx.Event.(*match.SpellCast); ok {

		if event.CardID == card.ID {
			return true
		}

	}

	return false

}

// Attacking returns true if the card is attacking a player or creature
func Attacking(card *match.Card, ctx *match.Context) bool {

	if event, ok := ctx.Event.(*match.AttackCreature); ok {
		if event.CardID == card.ID {
			return true
		}
	}

	if event, ok := ctx.Event.(*match.AttackPlayer); ok {
		if event.CardID == card.ID {
			return true
		}
	}

	return false

}

// AttackingPlayer returns true if the card is attacking a player
func AttackingPlayer(card *match.Card, ctx *match.Context) bool {

	if event, ok := ctx.Event.(*match.AttackPlayer); ok {
		if event.CardID == card.ID {
			return true
		}
	}

	return false

}

// AttackingCreature returns true if the card is attacking a Creature
func AttackingCreature(card *match.Card, ctx *match.Context) bool {

	if event, ok := ctx.Event.(*match.AttackCreature); ok {
		if event.CardID == card.ID {
			return true
		}
	}

	return false

}

// Destroyed returns true if the card was destroyed
func Destroyed(card *match.Card, ctx *match.Context) bool {

	if event, ok := ctx.Event.(*match.CreatureDestroyed); ok {
		if event.Card.ID == card.ID {
			return true
		}
	}

	return false

}
