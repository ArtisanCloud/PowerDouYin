package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/PowerDouYin/src/kernel"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
)

type AccessToken struct {
	*kernel.AccessToken
}

func NewAccessToken(app *kernel.ApplicationInterface) (*AccessToken, error) {
	kernelToken, err := kernel.NewAccessToken(app)
	token := &AccessToken{
		kernelToken,
	}

	// Override fields and functions
	token.EndpointToGetToken = "https://api.weixin.qq.com/cgi-bin/token"
	token.OverrideGetCredentials()

	return token, err
}

// Override GetCredentials
func (accessToken *AccessToken) OverrideGetCredentials() {
	config := (*accessToken.App).GetConfig()

	accessToken.GetCredentials = func() *object.StringMap {
		return &object.StringMap{
			"grant_type": "client_credential",
			"appid":      config.GetString("app_id", ""),
			"secret":     config.GetString("secret", ""),
			"neededText": "",
		}
	}
}
