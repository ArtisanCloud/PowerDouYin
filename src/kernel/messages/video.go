package messages

import (
	"github.com/ArtisanCloud/PowerDouYin/src/kernel/power"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
)

type Video struct {
	*Media
}

func NewVideo(mediaID string, item *power.HashMap) *Video {

	m := &Video{
		NewMedia(mediaID, "video", item),
	}

	m.Properties = []string{
		"title",
		"description",
		"media_id",
		"thumb_media_id",
	}

	m.OverrideToXmlArray()

	return m
}

// Override ToXmlArray
func (msg *Video) OverrideToXmlArray() {
	msg.ToXmlArray = func() *object.HashMap {

		return &object.HashMap{
			"Video": &object.HashMap{
				"MediaId":     msg.GetString("media_id", ""),
				"Title":       msg.GetString("title", ""),
				"Description": msg.GetString("description", ""),
			},
		}

	}
}
