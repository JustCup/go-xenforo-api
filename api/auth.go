package api // import "github.com/JustCup/xenforo-sdk/api"
import (
	"github.com/JustCup/xenforo-sdk/object"
)

// Auth function.
//
// Tests a login and password for validity.
//
// Only available to super user keys.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_auth_
func (xf *XF) Auth(params Params) (response object.AuthResponse, err error) {
	err = xf.RequestUnmarshal("POST", "auth", params, &response)
	return
}

// AuthFromSession function.
//
// Looks up the active XenForo user based on session ID or remember cookie value.
// This can be used to help with seamless SSO with XF, assuming the session or remember cookies are available to your page.
// At least one of session_id and remember_cookie must be provided.
//
// Only available to super user keys.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_auth_from_session
func (xf *XF) AuthFromSession(params Params) (response object.AuthResponse, err error) {
	err = xf.RequestUnmarshal("POST", "auth/from-session", params, &response)
	return
}

// AuthLoginToken function.
//
// Generates a token that can automatically log into a specific XenForo user when the login URL is visited.
// If the visitor is already logged into a XenForo account, they will not be logged into the specified account.
//
// Only available to super user keys.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_auth_login_token
func (xf *XF) AuthLoginToken(params Params) (response object.AuthResponse, err error) {
	err = xf.RequestUnmarshal("POST", "auth/login-token", params, &response)
	return
}
