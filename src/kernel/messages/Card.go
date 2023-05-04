package messages

import (
	"github.com/ArtisanCloud/PowerDouYin/src/kernel/power"
)

type Card struct {
	*Message
}

func NewCard(cardID string) *Card {
	m := &Card{
		NewMessage(&power.HashMap{"card_id": cardID}),
	}

	m.Type = "wxcard"

	m.Properties = []string{"card_id"}

	return m
}
