package main

import (
	"fmt"
	fmt2 "github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment"
	"os"
)

func GetMiniProgramConfig() *miniProgram.UserConfig {
	return &miniProgram.UserConfig{

		AppID:  os.Getenv("miniprogram_app_id"), // 小程序、公众号或者企业微信的appid
		Secret: os.Getenv("miniprogram_secret"), // 商户号 appID

		ResponseType: os.Getenv("array"),
		Log: miniProgram.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		Cache: kernel.NewRedisClient(&kernel.RedisOptions{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       1,
		}),
		HttpDebug: true,
		//Debug: true,
		//"sandbox": true,

	}

}

func main() {

	fmt.Printf("hello Wechat! \n")

	// init payment app
	configPayment := GetPaymentConfig()
	paymentApp, err := payment.NewPayment(configPayment)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt2.Dump("payment config:", paymentApp.GetConfig().All())

	// init miniProgram app
	configMiniProgram := GetMiniProgramConfig()
	miniProgramApp, err := miniProgram.NewMiniProgram(configMiniProgram)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt2.Dump("miniprogram config:", miniProgramApp.GetConfig().All())

}
