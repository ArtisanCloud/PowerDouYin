package messages

import (
	"github.com/ArtisanCloud/PowerDouYin/src/kernel/power"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
)

type ReplyInteractiveTaskCard struct {
	*Message
}

func NewReplyInteractiveTaskCard(replaceName string) *ReplyInteractiveTaskCard {
	m := &ReplyInteractiveTaskCard{
		NewMessage(&power.HashMap{"replace_name": replaceName}),
	}

	m.Type = "update_taskcard"

	m.Properties = []string{
		"replace_name",
	}

	m.OverrideToXmlArray()

	return m
}

// Override ToXmlArray
func (msg *ReplyInteractiveTaskCard) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {
		return &object.HashMap{
			"TaskCard": &object.HashMap{
				"ReplaceName": msg.Get("replace_name", nil),
			},
		}
	}
}
