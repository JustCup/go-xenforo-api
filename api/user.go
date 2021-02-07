package api // import "github.com/JustCup/go-xenforo-api/api"
import (
	"fmt"

	"github.com/JustCup/go-xenforo-api/object"
)

// UsersGet function.
//
// Gets a list of users (alphabetically).
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_users_
func (xf *XF) UsersGet(params Params) (response object.UsersGet, err error) {
	err = xf.RequestUnmarshal("GET", "users", params, &response)
	return
}

// UsersCreate function.
//
// Creates a user.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_users_
func (xf *XF) UsersCreate(params Params) (response object.UserCreate, err error) {
	err = xf.RequestUnmarshal("POST", "users", params, &response)
	return
}

// UsersFindName function.
//
// Finds users by a prefix of their user name.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_users_find_name
func (xf *XF) UsersFindName(params Params) (response object.UserGetFindName, err error) {
	err = xf.RequestUnmarshal("GET", "users/find-name", params, &response)
	return
}

// UsersGetID function.
//
// Gets information about the specified user.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_users_id_
func (xf *XF) UsersGetID(id int, params Params) (response object.UsersGetID, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("%s/%d/", "users", id), params, &response)
	return
}

// UsersEdit function.
//
// Updates an existing user.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_users_id_
func (xf *XF) UsersEdit(id int, params Params) (response object.UserCreate, err error) {
	err = xf.RequestUnmarshal("POST", fmt.Sprintf("%s/%d/", "users", id), params, &response)
	return
}
