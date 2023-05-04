package contract

import (
	response2 "github.com/ArtisanCloud/PowerDouYin/src/kernel/response"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"net/http"
)

type (
	AccessTokenInterface interface {
		GetToken(refresh bool) (resToken *response2.ResponseGetToken, err error)
		Refresh() AccessTokenInterface
		ApplyToRequest(request *http.Request, requestOptions *object.HashMap) (*http.Request, error)
	}
)
