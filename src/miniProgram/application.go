package miniProgram

import (
	"github.com/ArtisanCloud/PowerDouYin/src/kernel"
	"github.com/ArtisanCloud/PowerDouYin/src/kernel/providers"
	"github.com/ArtisanCloud/PowerDouYin/src/miniProgram/auth"
	"github.com/ArtisanCloud/PowerDouYin/src/miniProgram/base"
	"github.com/ArtisanCloud/PowerDouYin/src/miniProgram/customerServiceMessage"
	"github.com/ArtisanCloud/PowerLibs/v3/logger"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
)

type MiniProgram struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken
	Auth        *auth.Client

	Encryptor              *Encryptor
	CustomerServiceMessage *customerServiceMessage.Client

	Config *kernel.Config
	Logger *logger.Logger
}

type UserConfig struct {
	AppID  string
	Secret string

	RefreshToken      string
	ComponentAppID    string
	ComponentAppToken string
	Token             string
	AESKey            string

	ResponseType string
	Log          Log
	OAuth        OAuth
	Cache        kernel.CacheInterface

	HttpDebug bool
	Debug     bool
	NotifyURL string
	Sandbox   bool
}
type Log struct {
	Level string
	File  string
	ENV   string
}

type OAuth struct {
	Callback string
	Scopes   []string
}

func NewMiniProgram(config *UserConfig, extraInfos ...*kernel.ExtraInfo) (*MiniProgram, error) {
	var err error

	userConfig, err := MapUserConfig(config)
	if err != nil {
		return nil, err
	}

	var extraInfo, _ = kernel.NewExtraInfo()
	if len(extraInfos) > 0 {
		extraInfo = extraInfos[0]
	}

	// init an app container
	container, err := kernel.NewServiceContainer(userConfig, extraInfo)
	if err != nil {
		return nil, err
	}
	container.GetConfig()

	// init app
	app := &MiniProgram{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	app.Logger, err = logger.NewLogger("", &object.HashMap{
		"env":        app.Config.GetString("log.env", "develop"),
		"outputPath": app.Config.GetString("log.file", "./wechat.log"),
		"errorPath":  app.Config.GetString("log.file", "./wechat.log"),
	})
	if err != nil {
		return nil, err
	}

	//-------------- register auth,AccessToken --------------
	app.AccessToken, err = auth.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	app.Auth, err = auth.RegisterAuthProvider(app)
	if err != nil {
		return nil, err
	}

	// -------------- register Encryptor --------------
	app.Encryptor, err = NewEncryptor(
		(app.Config).GetString("app_id", ""),
		(app.Config).GetString("token", ""),
		(app.Config).GetString("aes_key", ""),
	)
	if err != nil {
		return nil, err
	}

	//-------------- register Base --------------
	app.Base, err = base.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register CustomerServiceMessage --------------
	app.CustomerServiceMessage, err = customerServiceMessage.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	return app, err
}

func (app *MiniProgram) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *MiniProgram) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}

func (app *MiniProgram) GetConfig() *kernel.Config {
	return app.Config
}

func (app *MiniProgram) GetComponent(name string) interface{} {

	switch name {
	case "Base":
		return app.Base
	case "AccessToken":
		return app.AccessToken
	case "auth":
		return app.Auth
	case "Config":
		return app.Config

	case "Encryptor":
		return app.Encryptor.BaseEncryptor

	case "CustomerServiceMessage":
		return app.CustomerServiceMessage

	case "Logger":
		return app.Logger

	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{

		"app_id": userConfig.AppID,
		"secret": userConfig.Secret,

		"token":               userConfig.Token,
		"aes_key":             userConfig.AESKey,
		"refresh_token":       userConfig.RefreshToken,
		"component_app_id":    userConfig.ComponentAppID,
		"component_app_token": userConfig.ComponentAppToken,

		"response_type": userConfig.ResponseType,
		"log": &object.HashMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
			"env":   userConfig.Log.ENV,
		},
		"cache": userConfig.Cache,

		"http_debug": userConfig.HttpDebug,
		"debug":      userConfig.Debug,
	}

	return config, nil

}
