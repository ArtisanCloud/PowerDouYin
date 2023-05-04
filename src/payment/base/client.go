package base

import (
	"context"
	response2 "github.com/ArtisanCloud/PowerDouYin/src/kernel/response"
	"github.com/ArtisanCloud/PowerDouYin/src/payment/kernel"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"net/http"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 付款码支付
// https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1

func (comp *Client) Pay(ctx context.Context, params *object.StringMap) *response2.ResponseWork {

	result := &response2.ResponseWork{}

	endpoint := comp.BaseClient.Wrap("/v3/pay/micropay")
	comp.BaseClient.Request(ctx, endpoint, params, http.MethodPost, nil, false, nil, result)

	return result
}

// 付款码查询openid
// https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_13&index=9
func (comp *Client) AuthCodeToOpenID(ctx context.Context, authCode string) *response2.ResponseWork {

	config := (*comp.BaseClient.App).GetConfig()
	appID := config.GetString("app_id", "")

	result := &response2.ResponseWork{}

	comp.BaseClient.Request(ctx, "tools/authcodetoopenid", &object.StringMap{
		"appid":     appID,
		"auth_code": authCode,
	}, http.MethodPost, nil, false, nil, result)

	return result
}
