package api

import "github.com/JustCup/xenforo-sdk/object"

// Me function.
//
// Gets information about the current API user.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_me_
func (xf *XF) Me() (response object.MeResponse, err error) {
	err = xf.RequestUnmarshal("GET", "me", Params{}, &response)
	return
}

// MeUpdate function.
//
// Updates information about the current user.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_me_
func (xf *XF) MeUpdate(params Params) (response object.MeResponse, err error) {
	err = xf.RequestUnmarshal("POST", "me", params, &response)
	return
}

// TODO: MeUpdateAvatar function.

// MeDeleteAvatar function.
//
// Deletes the current user's avatar.
//
// https://xenforo.com/community/pages/api-endpoints/#route_delete_me_avatar
func (xf *XF) MeDeleteAvatar() (response object.MeResponse, err error) {
	err = xf.RequestUnmarshal("DELETE", "me/avatar", Params{}, &response)
	return
}

// MeEmail function.
//
// Updates the current user's email address.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_me_email
func (xf *XF) MeEmail(params Params) (response object.MeResponse, err error) {
	err = xf.RequestUnmarshal("POST", "me/email", params, &response)
	return
}

// MePassword function.
//
// Updates the current user's password.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_me_password
func (xf *XF) MePassword(params Params) (response object.MeResponse, err error) {
	err = xf.RequestUnmarshal("POST", "me/password", params, &response)
	return
}
