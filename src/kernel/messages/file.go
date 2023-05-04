package messages

import (
	"github.com/ArtisanCloud/PowerDouYin/src/kernel/power"
)

type File struct {
	*Media
}

func NewFile(mediaID string, items *power.HashMap) *File {
	m := &File{
		NewMedia(mediaID, "file", items),
	}

	return m
}
