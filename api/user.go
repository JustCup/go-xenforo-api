package api // import "github.com/JustCup/go-xenforo-api/api"
import (
	"fmt"

	"github.com/JustCup/go-xenforo-api/object"
)

// UsersGetID funciton.
//
// Gets information about the specified user.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_users_id_
func (xf *XF) UsersGetID(id int, params Params) (response object.UsersGet, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("%s/%d/", "users", id), params, &response)
	return
}
