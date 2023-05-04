package messages

import (
	"github.com/ArtisanCloud/PowerDouYin/src/kernel/power"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
)

type NewsArticle struct {
	*Message
}

func NewNewsArticle(items *power.HashMap) *NewsArticle {
	m := &NewsArticle{
		NewMessage(items),
	}

	m.Type = "mpnewsarticle"
	m.Properties = []string{
		"article_id",
	}

	m.JsonAliases = &object.HashMap{
		//"content_source_url": "source_url",
		//"show_cover_pic":     "show_cover",
	}

	m.SetAttribute("required", []string{
		"article_id",
	})

	return m
}
