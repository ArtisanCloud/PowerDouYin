package customerServiceMessage

import (
	"github.com/ArtisanCloud/PowerDouYin/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil

}
