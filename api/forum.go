package api // import "github.com/JustCup/xenforo-sdk/api"

import (
	"fmt"

	"github.com/JustCup/xenforo-sdk/object"
)

// ForumsGet function.
//
// Gets information about the specified forum.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_forums_id_
func (xf *XF) ForumsGet(id int, params Params) (response object.ForumsGet, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("%s/%d/", "forums", id), params, &response)
	return
}

// ForumsGetThreads function.
//
// Gets a page of threads from the specified forum.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_forums_id_
func (xf *XF) ForumsGetThreads(id int, params Params) (response object.ForumsGetThreads, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("%s/%d/%s/", "forums", id, "threads"), params, &response)
	return
}
