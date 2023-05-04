package messages

import (
	"github.com/ArtisanCloud/PowerDouYin/src/kernel/power"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
)

type Text struct {
	*Message
}

func NewText(content string) *Text {
	m := &Text{
		NewMessage(&power.HashMap{"content": content}),
	}

	m.Type = "text"
	m.Properties = []string{"content"}
	m.OverrideToXmlArray()

	return m
}

// Override ToXmlArray
func (msg *Text) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {
		return &object.HashMap{
			"Content": msg.Get("content", nil),
		}
	}
}
