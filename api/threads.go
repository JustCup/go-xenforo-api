package api

import (
	"fmt"

	"github.com/JustCup/xenforo-sdk/object"
)

// ThreadsIDMove function.
//
// Moves the specified thread to a different forum. Only simple title/prefix updates are supported at the same time.
//
// https://xenforo.com/community/pages/api-endpoints/#route_post_threads_id_move
func (xf *XF) ThreadsIDMove(id int, params Params) (response object.ThreadResponse, err error) {
	err = xf.RequestUnmarshal("POST", fmt.Sprintf("%s/%d/%s", "threads", id, "move"), params, &response)
	return
}
