package api // import "github.com/JustCup/xenforo-sdk/api"
import (
	"github.com/JustCup/xenforo-sdk/object"
)

// Auth function.
//
// Tests a login and password for validity. Only available to super user keys.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_auth_
func (xf *XF) Auth(params Params) (response object.AuthResponse, err error) {
	err = xf.RequestUnmarshal("POST", "auth", params, &response)
	return
}
