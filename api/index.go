package api

import (
	"github.com/JustCup/xenforo-sdk/object"
)

// Index function.
//
// Gets general information about the site and the API.
//
// https://xenforo.com/community/pages/api-endpoints/#route_get_index_
func (xf *XF) Index() (response object.IndexResponse, err error) {
	err = xf.RequestUnmarshal("GET", "index", Params{}, &response)
	return
}
