package providers

import "github.com/ArtisanCloud/PowerDouYin/src/kernel"

func RegisterConfigProvider(app kernel.ApplicationInterface) *kernel.Config {

	return kernel.NewConfig(app.GetContainer().GetConfig())
}
