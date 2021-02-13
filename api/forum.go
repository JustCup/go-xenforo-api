package api // import "github.com/JustCup/xenforo-sdk/api"

import (
	"fmt"

	"github.com/JustCup/xenforo-sdk/object"
)

// Forums function.
//
// Gets information about the specified forum.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_forums_id_
func (xf *XF) Forums(id int, params Params) (response object.ForumsResponse, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("%s/%d/", "forums", id), params, &response)
	return
}

// ForumsMarkRead function.
//
// Marks the forum as read up until the specified time. This cannot mark a forum as unread.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_forums_id_mark_read
func (xf *XF) ForumsMarkRead(id int, params Params) (response object.ForumsResponse, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("%s/%d/%s", "forums", id, "mark-read"), params, &response)
	return
}

// ForumsThreads function.
//
// Gets a page of threads from the specified forum.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_forums_id_
func (xf *XF) ForumsThreads(id int, params Params) (response object.ForumsResponse, err error) {
	err = xf.RequestUnmarshal("GET", fmt.Sprintf("%s/%d/%s/", "forums", id, "threads"), params, &response)
	return
}
