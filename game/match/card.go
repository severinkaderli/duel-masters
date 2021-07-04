package match

import (
	"github.com/sirupsen/logrus"
	"github.com/ventu-io/go-shortid"
)

// Condition is used to store turn-specific state to the card such as power amplifiers
type Condition struct {
	ID  string
	Val interface{}
	Src interface{}
}

// Card holds information about a specific card
type Card struct {
	ID      string
	ImageID string
	Player  *Player
	Tapped  bool
	Zone    string

	Name            string
	Power           int
	Civ             string
	Family          string
	ManaCost        int
	ManaRequirement []string
	PowerModifier   func(m *Match, attacking bool) int

	attachedCards []*Card
	conditions    []Condition
	handlers      []HandlerFunc
}

// NewCard returns a new, initialized card
func NewCard(p *Player, image string) (*Card, error) {

	id, err := shortid.Generate()

	if err != nil {
		logrus.Debug("Failed to generate id for card")
		return nil, err
	}

	c := &Card{
		ID:              id,
		ImageID:         image,
		Player:          p,
		Tapped:          false,
		Zone:            DECK,
		Name:            "undefined_card",
		Power:           0,
		Civ:             "undefind_civ",
		Family:          "undefined_family",
		ManaCost:        1,
		ManaRequirement: make([]string, 0),
		PowerModifier:   func(m *Match, attacking bool) int { return 0 },
	}

	cardctor, err := CardCtor(image)

	if err != nil {
		logrus.Warn(err)
		return nil, err
	}

	cardctor(c)

	return c, nil

}

// Use allows different cards to hook into match events
// Can be compared to a typical middleware function
func (c *Card) Use(handlers ...HandlerFunc) {
	c.handlers = append(c.handlers, handlers...)
}

// Conditions returns a slice with the cards conditions
func (c *Card) Conditions() []Condition {
	return c.conditions
}

// AddCondition stores a string to the state of the card that will stay there until removed
func (c *Card) AddCondition(cnd string, val interface{}, src interface{}) {
	c.conditions = append(c.conditions, Condition{cnd, val, src})
}

// AddUniqueSourceCondition adds a condition only if the specified cnd and source is not already added
func (c *Card) AddUniqueSourceCondition(cnd string, val interface{}, src interface{}) {

	for _, condition := range c.conditions {
		if condition.ID == cnd && condition.Src == src {
			return
		}
	}

	c.conditions = append(c.conditions, Condition{cnd, val, src})
}

// HasCondition returns true or false based on if a given string is added to the cards list of conditions
func (c *Card) HasCondition(cnd string) bool {

	for _, condition := range c.conditions {
		if condition.ID == cnd {
			return true
		}
	}

	return false

}

// RemoveCondition removes all instances of the given string from the cards conditions
func (c *Card) RemoveCondition(cnd string) {

	tmp := make([]Condition, 0)

	for _, condition := range c.conditions {

		if condition.ID != cnd {
			tmp = append(tmp, condition)
		}

	}

	c.conditions = tmp

}

// RemoveConditionBySource removes all instances of conditions with given source
func (c *Card) RemoveConditionBySource(src string) {

	tmp := make([]Condition, 0)

	for _, condition := range c.conditions {

		if condition.Src != src {
			tmp = append(tmp, condition)
		}

	}

	c.conditions = tmp

}

// ClearConditions removes all conditions from the card
func (c *Card) ClearConditions() {

	c.conditions = make([]Condition, 0)

}

// Attach adds a *Card to the card's list of attached cards
func (c *Card) Attach(toAttach ...*Card) {
	c.attachedCards = append(c.attachedCards, toAttach...)
}

// Detach removes a *Card from the card's list of attached cards
func (c *Card) Detach(toDetach *Card) {

	tmp := make([]*Card, 0)

	for _, card := range c.attachedCards {

		if card.ID != toDetach.ID {
			tmp = append(tmp, card)
		}

	}

	c.attachedCards = tmp

}

// Attachments returns a copy of the card's attached cards
func (c *Card) Attachments() []*Card {

	return c.attachedCards

}

// ClearAttachments removes all attached cards
func (c *Card) ClearAttachments() {
	c.attachedCards = make([]*Card, 0)
}
