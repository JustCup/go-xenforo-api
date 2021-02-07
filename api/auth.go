package api // import "github.com/JustCup/go-xenforo-api/api"
import (
	"github.com/JustCup/go-xenforo-api/object"
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
