package jssdk

import "github.com/ArtisanCloud/PowerDouYin/src/kernel"

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	return NewClient(&app)

}
