package customerServiceMessage

import (
	"context"
	"github.com/ArtisanCloud/PowerDouYin/src/kernel"
	"github.com/ArtisanCloud/PowerDouYin/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerDouYin/src/kernel/response"
	"github.com/ArtisanCloud/PowerDouYin/src/miniProgram/customerServiceMessage/request"
	"github.com/ArtisanCloud/PowerDouYin/src/miniProgram/customerServiceMessage/response"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 发送客服消息给用户
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html
func (comp *Client) Send(ctx context.Context, toUser string, msgType string, msg interface{}) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"touser":  toUser,
		"msgtype": msgType,
		msgType:   msg,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/custom/send", data, nil, nil, result)

	return result, err
}
func (comp *Client) SendText(ctx context.Context, toUser string, msg *request.CustomerServiceMsgText) (*response2.ResponseMiniProgram, error) {
	return comp.Send(ctx, toUser, "text", msg)
}
func (comp *Client) SendImage(ctx context.Context, toUser string, msg *request.CustomerServiceMsgImage) (*response2.ResponseMiniProgram, error) {
	return comp.Send(ctx, toUser, "image", msg)
}
func (comp *Client) SendLink(ctx context.Context, toUser string, msg *request.CustomerServiceMsgLink) (*response2.ResponseMiniProgram, error) {
	return comp.Send(ctx, toUser, "link", msg)
}
func (comp *Client) SendMiniProgramPage(ctx context.Context, toUser string, msg *request.CustomerServiceMsgMpPage) (*response2.ResponseMiniProgram, error) {
	return comp.Send(ctx, toUser, "miniprogrampage", msg)
}

// 下发客服当前输入状态给用户
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.setTyping.html
func (comp *Client) SetTyping(ctx context.Context, toUser string, command string) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"touser":  toUser,
		"command": command,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/message/custom/typing", data, nil, nil, result)

	return result, err
}

// 把媒体文件上传到微信服务器
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html
func (comp *Client) UploadTempMedia(ctx context.Context, mediaType string, path string, form *power.HashMap) (*response.ResponseCustomerServiceMessageUploadTempMedia, error) {

	result := &response.ResponseCustomerServiceMessageUploadTempMedia{}

	var files *object.HashMap
	if path != "" {
		files = &object.HashMap{
			"media": path,
		}
	}

	var formData *kernel.UploadForm
	if form != nil {
		formData = &kernel.UploadForm{
			Contents: []*kernel.UploadContent{
				&kernel.UploadContent{
					Name:  (*form)["name"].(string),
					Value: (*form)["value"],
				},
			},
		}
	}

	_, err := comp.BaseClient.HttpUpload(ctx, "cgi-bin/media/upload", files, formData, &object.StringMap{
		"type": mediaType,
	}, nil, result)

	return result, err
}
